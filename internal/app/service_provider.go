package app

import (
	"context"
	"log"

	"github.com/drizzleent/hezzl/internal/api"
	"github.com/drizzleent/hezzl/internal/api/http/handler"
	"github.com/drizzleent/hezzl/internal/config"
	"github.com/drizzleent/hezzl/internal/config/env"
	"github.com/drizzleent/hezzl/internal/service"
	"github.com/drizzleent/hezzl/internal/service/goods"
)

type serviceProvicer struct {
	httpConfig config.HTTPConfig

	service service.Service

	handler api.Handler
}

func newServiceProvider() *serviceProvicer {
	return &serviceProvicer{}
}

func (s *serviceProvicer) HttpConfig() config.HTTPConfig {
	if nil == s.httpConfig {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvicer) Service(ctx context.Context) service.Service {
	if nil == s.service {
		s.service = goods.NewService()
	}

	return s.service
}

func (s *serviceProvicer) Handler(ctx context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler(s.Service(ctx))
	}

	return s.handler
}
