package native

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/logicalangel/tashil_test/internal/consts"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/client"
)

type nativeClient struct {
}

func (c *nativeClient) Get(ctx context.Context, url string, headers map[string]interface{}) (model.Call, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s", url), nil)
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("NewRequest")
		return model.Call{}, consts.ErrInternalError
	}

	for key, val := range headers {
		req.Header.Add(key, val.(string))
	}

	start := time.Now()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("DoRequest")
		return model.Call{}, consts.ErrInternalError
	}

	elapsed := time.Since(start)

	return model.Call{
		Status: uint(res.StatusCode),
		Time:   uint(elapsed.Milliseconds()),
	}, nil
}

func (c *nativeClient) Post(ctx context.Context, url string, body string, headers map[string]interface{}) (model.Call, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://%s", url), bodyReader)
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("NewRequest")
		return model.Call{}, consts.ErrInternalError
	}

	for key, val := range headers {
		req.Header.Add(key, val.(string))
	}

	start := time.Now()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("DoRequest")
		return model.Call{}, consts.ErrInternalError
	}

	elapsed := time.Since(start)

	return model.Call{
		Status: uint(res.StatusCode),
		Time:   uint(elapsed.Milliseconds()),
	}, nil
}

func (c *nativeClient) Put(ctx context.Context, url string, body string, headers map[string]interface{}) (model.Call, error) {
	bodyReader := bytes.NewReader([]byte(body))

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("https://%s", url), bodyReader)
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("NewRequest")
		return model.Call{}, consts.ErrInternalError
	}

	for key, val := range headers {
		req.Header.Add(key, val.(string))
	}

	start := time.Now()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("DoRequest")
		return model.Call{}, consts.ErrInternalError
	}

	elapsed := time.Since(start)

	return model.Call{
		Status: uint(res.StatusCode),
		Time:   uint(elapsed.Milliseconds()),
	}, nil
}

func (c *nativeClient) Delete(ctx context.Context, url string, headers map[string]interface{}) (model.Call, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("https://%s", url), nil)
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("NewRequest")
		return model.Call{}, consts.ErrInternalError
	}

	for key, val := range headers {
		req.Header.Add(key, val.(string))
	}

	start := time.Now()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.WithFields(log.Fields{"message": err}).Error("DoRequest")
		return model.Call{}, consts.ErrInternalError
	}

	elapsed := time.Since(start)

	return model.Call{
		Status: uint(res.StatusCode),
		Time:   uint(elapsed.Milliseconds()),
	}, nil
}

func New() client.Client {
	return &nativeClient{}
}
