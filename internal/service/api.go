package service

import (
	"context"

	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/repository"
)

type IApiService interface {
	GetAll(ctx context.Context) ([]model.Api, error)
	Get(ctx context.Context, apiID uint) (model.Api, []model.Call, error)
	Create(ctx context.Context, input model.Api) (model.Api, error)
	Start(ctx context.Context, apiID uint) (model.Api, error)
	Stop(ctx context.Context, apiID uint) (model.Api, error)
	Delete(ctx context.Context, apiID uint) error
}

type apiService struct {
	apiRepository  repository.ApiRepository
	callRepository repository.CallRepository
}

func (s *apiService) GetAll(ctx context.Context) ([]model.Api, error) {
	apis, err := s.apiRepository.GetAll(ctx)
	if err != nil {
		return []model.Api{}, err
	}

	return apis, nil
}

func (s *apiService) Get(ctx context.Context, apiID uint) (model.Api, []model.Call, error) {
	api, err := s.apiRepository.Get(ctx, apiID)
	if err != nil {
		return api, []model.Call{}, err
	}

	calls, err := s.callRepository.Get(ctx, apiID)
	if err != nil {
		return api, []model.Call{}, err
	}

	return api, calls, nil
}

func (s *apiService) Create(ctx context.Context, input model.Api) (model.Api, error) {
	api, err := s.apiRepository.Create(ctx, input.Url, input.Method, input.CallInterval, input.Headers, input.Body)
	if err != nil {
		return model.Api{}, err
	}

	return api, nil
}

func (s *apiService) Start(ctx context.Context, apiID uint) (model.Api, error) {
	api, err := s.apiRepository.Start(ctx, apiID)
	if err != nil {
		return model.Api{}, err
	}

	return api, nil
}

func (s *apiService) Stop(ctx context.Context, apiID uint) (model.Api, error) {
	api, err := s.apiRepository.Stop(ctx, apiID)
	if err != nil {
		return model.Api{}, err
	}

	return api, nil
}

func (s *apiService) Delete(ctx context.Context, apiID uint) error {
	err := s.apiRepository.Delete(ctx, apiID)
	if err != nil {
		return err
	}

	return nil
}

func NewApi(apiRepository repository.ApiRepository, callRepository repository.CallRepository) IApiService {
	return &apiService{
		apiRepository:  apiRepository,
		callRepository: callRepository,
	}
}
