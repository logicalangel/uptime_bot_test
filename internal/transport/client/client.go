package client

import (
	"context"

	"github.com/logicalangel/tashil_test/internal/model"
)

type Client interface {
	Get(ctx context.Context, url string, headers map[string]interface{}) (model.Call, error)
	Post(ctx context.Context, url string, body string, headers map[string]interface{}) (model.Call, error)
	Put(ctx context.Context, url string, body string, headers map[string]interface{}) (model.Call, error)
	Delete(ctx context.Context, url string, headers map[string]interface{}) (model.Call, error)
}
