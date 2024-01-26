package controller

import (
	"errors"

	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/logicalangel/tashil_test/internal/consts"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/service"
	"github.com/logicalangel/tashil_test/internal/transport/server/fiber/dto"
)

type Api struct {
	apiService service.IApiService
}

// GetAll godoc
// @Summary This function return all apis
// @Description see apis
// @Description Errors: ErrBadBodyRequest, internal_error
// @Tags Apis
// @Accept		json
// @Produce		json
// @Success 	200 {object}	dto.Response
// @Failure		400	{object}	dto.ResponseError
// @Failure		500	{object}	dto.ResponseError
// @Router /api [get].
func (a *Api) GetAll(ctx *fiber.Ctx) error {
	apis, err := a.apiService.GetAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Message: consts.ErrInternalError.Error(),
		})
	}

	return ctx.JSON(dto.Response{
		Status: true,
		Data:   dto.NewApis(apis),
	})
}

// Get godoc
// @Summary This function return apis calls
// @Description see apis calls
// @Description Errors: ErrBadBodyRequest, internal_error, api_notfound
// @Tags Apis
// @Accept		json
// @Produce		json
// @Param			api_id	path		int		true	"API ID"
// @Success 	200 {object}	dto.Response
// @Failure		400	{object}	dto.ResponseError
// @Failure		500	{object}	dto.ResponseError
// @Router /api/{api_id} [get].
func (a *Api) Get(ctx *fiber.Ctx) error {
	apiID, err := ctx.ParamsInt("api_id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Message: consts.ErrBadBodyRequest.Error(),
		})
	}

	api, calls, err := a.apiService.Get(ctx.Context(), uint(apiID))
	if err != nil {
		if errors.Is(err, consts.ErrApiNotfound) {
			return ctx.Status(fiber.StatusNotFound).JSON(dto.ResponseError{
				Message: consts.ErrApiNotfound.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Message: consts.ErrInternalError.Error(),
		})
	}

	return ctx.JSON(dto.Response{
		Status: true,
		Data: dto.ApiDetailResponse{
			Api:   dto.NewApi(api),
			Calls: dto.NewCalls(calls),
		},
	})
}

// Create godoc
// @Summary This function create a new api
// @Description create api
// @Description Errors: ErrBadBodyRequest, internal_error
// @Tags Apis
// @Accept		json
// @Produce		json
// @Success 	200 {object}	dto.Response
// @Param			api	body		dto.Api	true	"Api infos"
// @Failure		400	{object}	dto.ResponseError
// @Failure		500	{object}	dto.ResponseError
// @Router /api [post].
func (a *Api) Create(ctx *fiber.Ctx) error {
	body := dto.Api{}
	err := ctx.BodyParser(&body)
	if err != nil {
		log.WithFields(log.Fields{"message": err.Error()}).Warn("Create:Body")
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Message: consts.ErrBadBodyRequest.Error(),
		})
	}

	err = ValidateBody(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	api, err := a.apiService.Create(ctx.Context(), model.Api{
		Status:       body.Status,
		Url:          body.Url,
		Method:       body.Method,
		CallInterval: body.CallInterval,
		Body:         body.Body,
		Headers:      body.Headers,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Message: consts.ErrInternalError.Error(),
		})
	}

	return ctx.JSON(dto.Response{
		Status: true,
		Data:   api,
	})
}

// Start godoc
// @Summary Start tracking on an api
// @Description start api tracking
// @Description Errors: ErrBadBodyRequest, internal_error, api_notfound
// @Tags Apis
// @Accept		json
// @Produce		json
// @Param			api_id	path		int		true	"API ID"
// @Success 	200 {object}	dto.Response
// @Failure		400	{object}	dto.ResponseError
// @Failure		500	{object}	dto.ResponseError
// @Router /api/{api_id}/start [post].
func (a *Api) Start(ctx *fiber.Ctx) error {
	apiID, err := ctx.ParamsInt("api_id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Message: consts.ErrBadBodyRequest.Error(),
		})
	}

	api, err := a.apiService.Start(ctx.Context(), uint(apiID))
	if err != nil {
		if errors.Is(err, consts.ErrApiNotfound) {
			return ctx.Status(fiber.StatusNotFound).JSON(dto.ResponseError{
				Message: consts.ErrApiNotfound.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Message: consts.ErrInternalError.Error(),
		})
	}

	return ctx.JSON(dto.Response{
		Status: true,
		Data:   api,
	})
}

// Stop godoc
// @Summary Stop tracking on an api
// @Description stop api tracking
// @Description Errors: ErrBadBodyRequest, internal_error, api_notfound
// @Tags Apis
// @Accept		json
// @Produce		json
// @Param			api_id	path		int		true	"API ID"
// @Success 	200 {object}	dto.Response
// @Failure		400	{object}	dto.ResponseError
// @Failure		500	{object}	dto.ResponseError
// @Router /api/{api_id}/stop [post].
func (a *Api) Stop(ctx *fiber.Ctx) error {
	apiID, err := ctx.ParamsInt("api_id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Message: consts.ErrBadBodyRequest.Error(),
		})
	}

	api, err := a.apiService.Stop(ctx.Context(), uint(apiID))
	if err != nil {
		if errors.Is(err, consts.ErrApiNotfound) {
			return ctx.Status(fiber.StatusNotFound).JSON(dto.ResponseError{
				Message: consts.ErrApiNotfound.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Message: consts.ErrInternalError.Error(),
		})
	}

	return ctx.JSON(dto.Response{
		Status: true,
		Data:   api,
	})
}

// Delete godoc
// @Summary Delete an api
// @Description delete api
// @Description Errors: ErrBadBodyRequest, internal_error, api_notfound
// @Tags Apis
// @Accept		json
// @Produce		json
// @Param			api_id	path		int		true	"API ID"
// @Success 	200 {object}	dto.Response
// @Failure		400	{object}	dto.ResponseError
// @Failure		500	{object}	dto.ResponseError
// @Router /api/{api_id} [delete].
func (a *Api) Delete(ctx *fiber.Ctx) error {
	apiID, err := ctx.ParamsInt("api_id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{
			Message: consts.ErrBadBodyRequest.Error(),
		})
	}

	err = a.apiService.Delete(ctx.Context(), uint(apiID))
	if err != nil {
		if errors.Is(err, consts.ErrApiNotfound) {
			return ctx.Status(fiber.StatusNotFound).JSON(dto.ResponseError{
				Message: consts.ErrApiNotfound.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{
			Message: consts.ErrInternalError.Error(),
		})
	}

	return ctx.JSON(dto.Response{
		Status: true,
	})
}

func NewApi(apiService service.IApiService) Api {
	return Api{
		apiService: apiService,
	}
}
