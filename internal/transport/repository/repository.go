package repository

import (
	"context"
	"github.com/logicalangel/tashil_test/internal/model"
)

type ApiRepository interface {
	GetAll(ctx context.Context) ([]model.Api, error)
	GetApisScheduled(ctx context.Context) ([]model.Api, error)
	Get(ctx context.Context, apiID uint) (model.Api, error)
	Create(ctx context.Context, url string, method model.ApiMethod, interval uint, headers map[string]interface{}, body string) (model.Api, error)
	Start(ctx context.Context, apiID uint) (model.Api, error)
	Stop(ctx context.Context, apiID uint) (model.Api, error)
	Delete(ctx context.Context, apiID uint) error
}

type CallRepository interface {
	Get(ctx context.Context, apiID uint) ([]model.Call, error)
	Create(ctx context.Context, apiID uint, status uint, time uint) (model.Call, error)
}
