package mock

import (
	"context"
	"math/rand"

	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/client"
)

type mockClient struct {
}

var statusCodes = []uint{200, 400, 401, 500, 511}

func (c *mockClient) Get(ctx context.Context, url string, headers map[string]interface{}) (model.Call, error) {
	ping := rand.Intn(1000) + 100
	statusCodeInx := rand.Intn(len(statusCodes))
	return model.Call{
		Time:   uint(ping),
		Status: statusCodes[statusCodeInx],
	}, nil
}

func (c *mockClient) Post(ctx context.Context, url string, body string, headers map[string]interface{}) (model.Call, error) {
	ping := rand.Intn(1000) + 100
	statusCodeInx := rand.Intn(len(statusCodes))
	return model.Call{
		Time:   uint(ping),
		Status: statusCodes[statusCodeInx],
	}, nil
}

func (c *mockClient) Put(ctx context.Context, url string, body string, headers map[string]interface{}) (model.Call, error) {
	ping := rand.Intn(1000) + 100
	statusCodeInx := rand.Intn(len(statusCodes))
	return model.Call{
		Time:   uint(ping),
		Status: statusCodes[statusCodeInx],
	}, nil
}

func (c *mockClient) Delete(ctx context.Context, url string, headers map[string]interface{}) (model.Call, error) {
	ping := rand.Intn(1000) + 100
	statusCodeInx := rand.Intn(len(statusCodes))
	return model.Call{
		Time:   uint(ping),
		Status: statusCodes[statusCodeInx],
	}, nil
}

func New() client.Client {
	return &mockClient{}
}
