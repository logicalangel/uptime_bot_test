package dto

import "github.com/logicalangel/tashil_test/internal/model"

type Api struct {
	ID           uint                   `json:"id"`
	Status       model.ApiStatus        `json:"status"`
	Url          string                 `json:"url"`
	Method       model.ApiMethod        `json:"method"`
	CallInterval uint                   `json:"call_interval"`
	Body         string                 `json:"body"`
	Headers      map[string]interface{} `json:"headers"`
	Webhook      string                 `json:"webhook"`
}

type Apis []model.Api

func NewApis(input Apis) []Api {
	result := []Api{}
	for _, api := range input {
		result = append(result, NewApi(api))
	}

	return result
}

func NewApi(input model.Api) Api {
	return Api{
		ID:           input.ID,
		Status:       input.Status,
		Url:          input.Url,
		Method:       input.Method,
		CallInterval: input.CallInterval,
		Body:         input.Body,
		Headers:      input.Headers,
		Webhook:      input.Webhook,
	}
}
