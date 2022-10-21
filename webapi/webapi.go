package sample

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/is-hoku/goa-sample/webapi/usecase"
)

func newMySQLConfig() *mysql.Config {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.Addr = net.JoinHostPort(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	config.DBName = os.Getenv("DB_NAME")
	config.Timeout = 30 * time.Second
	config.RejectReadOnly = true
	config.ParseTime = true
	return config
}

type StudentApp struct {
	getter      usecase.StudentByNumberGetter
	multigetter usecase.StudentsGetter
	creator     usecase.StudentCreator
}

func newStudentApp(ctx context.Context, getter usecase.StudentByNumberGetter, multigetter usecase.StudentsGetter, creator usecase.StudentCreator) (*StudentApp, error) {
	return &StudentApp{
		getter:      getter,
		multigetter: multigetter,
		creator:     creator,
	}, nil
}
