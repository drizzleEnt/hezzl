package goods

import "github.com/drizzleent/hezzl/internal/service"

type srv struct {
}

func NewService() service.Service {
	return &srv{}
}
