package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	sample "github.com/is-hoku/goa-sample/webapi"
	health "github.com/is-hoku/goa-sample/webapi/gen/health"
	student "github.com/is-hoku/goa-sample/webapi/gen/student"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "0.0.0.0", "Server host (valid values: 0.0.0.0)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[sample] ", log.Ltime)
	}

	// New App
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	studentApp, err := sample.NewStudentApp(ctx)
	if err != nil {
		logger.Fatalf("unepected error:%s", err)
	}

	// Initialize the services.
	var (
		healthSvc  health.Service
		studentSvc student.Service
	)
	{
		healthSvc = sample.NewHealth(logger)
		studentSvc = sample.NewStudent(logger, studentApp)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		healthEndpoints  *health.Endpoints
		studentEndpoints *student.Endpoints
	)
	{
		healthEndpoints = health.NewEndpoints(healthSvc)
		studentEndpoints = student.NewEndpoints(studentSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "0.0.0.0":
		{
			addr := "http://0.0.0.0:8080"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, healthEndpoints, studentEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: 0.0.0.0)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
