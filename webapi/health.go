package sample

import (
	"context"
	"log"

	health "github.com/is-hoku/goa-sample/webapi/gen/health"
)

type healthsrvc struct {
	logger *log.Logger
}

func NewHealth(logger *log.Logger) health.Service {
	return &healthsrvc{logger}
}

// ヘルスチェック
func (s *healthsrvc) Check(ctx context.Context) (*health.HealthResult, error) {
	res := &health.HealthResult{
		Message: "I'm healthy.",
	}
	s.logger.Print("health.check")
	return res, nil
}
