package dto

import (
	"time"

	"github.com/logicalangel/tashil_test/internal/model"
)

type Call struct {
	CreatedAt time.Time `json:"created_at"`
	Status    uint      `json:"status"`
	Time      uint      `json:"time"`
}

type ApiDetailResponse struct {
	Api   Api    `json:"api"`
	Calls []Call `json:"calls"`
}

func NewCall(input model.Call) Call {
	return Call{
		CreatedAt: input.CreatedAt,
		Status:    input.Status,
		Time:      input.Time,
	}
}

type Calls []model.Call

func NewCalls(input Calls) []Call {
	result := []Call{}
	for _, call := range input {
		result = append(result, NewCall(call))
	}

	return result
}
