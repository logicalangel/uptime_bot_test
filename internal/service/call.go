package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/client"
	"github.com/logicalangel/tashil_test/internal/transport/repository"
	"net/http"
)

type ICallService interface {
	CallScheduledApis(ctx context.Context) error
}

type callService struct {
	client         client.Client
	apiRepository  repository.ApiRepository
	callRepository repository.CallRepository
}

func (s *callService) CallScheduledApis(ctx context.Context) error {
	apis, err := s.apiRepository.GetApisScheduled(ctx)
	if err != nil {
		return err
	}

	for _, api := range apis {
		apiS, _ := json.Marshal(api)
		fmt.Println(string(apiS))
		call := model.Call{}
		switch api.Method {
		case model.GetMethod:
			call, err = s.client.Get(ctx, api.Url, api.Headers)
			if err != nil {
				return err
			}
		case model.PostMethod:
			call, err = s.client.Post(ctx, api.Url, api.Body, api.Headers)
			if err != nil {
				return err
			}
		case model.PutMethod:
			call, err = s.client.Put(ctx, api.Url, api.Body, api.Headers)
			if err != nil {
				return err
			}
		case model.DeleteMethod:
			call, err = s.client.Delete(ctx, api.Url, api.Headers)
			if err != nil {
				return err
			}
		}

		if call.Status != http.StatusOK && api.Webhook != "" {
			callS, err := json.Marshal(call)
			if err != nil {
				return err
			}

			call, err = s.client.Post(ctx, api.Webhook, string(callS), nil)
			if err != nil {
				return err
			}
		}

		_, err = s.callRepository.Create(ctx, api.ID, call.Status, call.Time)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewCall(
	client client.Client,
	apiRepository repository.ApiRepository,
	callRepository repository.CallRepository,
) ICallService {
	return &callService{
		client:         client,
		apiRepository:  apiRepository,
		callRepository: callRepository,
	}
}
