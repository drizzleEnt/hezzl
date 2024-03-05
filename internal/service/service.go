package service

import (
	"context"

	"github.com/drizzleent/hezzl/internal/model"
)

type Service interface {
	Create(context.Context) (*model.Good, error)
	GetList(context.Context) error
	Remove(context.Context) error
	Update(context.Context) error
	Repriotiize(context.Context) error
}
