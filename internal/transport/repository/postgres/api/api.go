package api

import (
	"context"
	"errors"
	"github.com/apex/log"
	"github.com/logicalangel/tashil_test/internal/consts"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/database"
	"github.com/logicalangel/tashil_test/internal/transport/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type apiRepository struct {
	db *gorm.DB
}

func (r *apiRepository) GetApisScheduled(ctx context.Context) ([]model.Api, error) {
	apis := []model.Api{}
	tx := r.db.WithContext(ctx).
		Model(&apis).
		Clauses(clause.Returning{}).
		Where(`EXTRACT(EPOCH FROM (current_timestamp - updated_at))/60 > call_interval`).
		Where("status = ?", model.ApiRunningStatus).
		Update("updated_at", time.Now())

	if tx.Error != nil {
		log.WithFields(log.Fields{"message": tx.Error}).Error("GetApisScheduled")
		return apis, consts.ErrInternalError
	}

	return apis, nil
}

func (r *apiRepository) GetAll(ctx context.Context) ([]model.Api, error) {
	apis := []model.Api{}
	tx := r.db.WithContext(ctx).Find(&apis)
	if tx.Error != nil {
		log.WithFields(log.Fields{"message": tx.Error}).Error("GetAll")
		return apis, consts.ErrInternalError
	}

	return apis, nil
}

func (r *apiRepository) Get(ctx context.Context, apiID uint) (model.Api, error) {
	api := model.Api{}
	tx := r.db.WithContext(ctx).Where("id = ?", apiID).First(&api)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return api, consts.ErrApiNotfound
		}
		log.WithFields(log.Fields{"message": tx.Error}).Error("Get")
		return api, consts.ErrInternalError
	}

	return api, nil
}

func (r *apiRepository) Create(
	ctx context.Context, url string, method model.ApiMethod, interval uint, headers map[string]interface{}, body string,
) (model.Api, error) {
	api := model.Api{
		Url:          url,
		Method:       method,
		Status:       model.ApiRunningStatus,
		CallInterval: interval,
		Headers:      headers,
		Body:         body,
	}
	tx := r.db.WithContext(ctx).Create(&api)
	if tx.Error != nil {
		log.WithFields(log.Fields{"message": tx.Error}).Error("Create")
		return api, consts.ErrInternalError
	}

	return api, nil
}

func (r *apiRepository) Start(ctx context.Context, apiID uint) (model.Api, error) {
	api := model.Api{}
	tx := r.db.WithContext(ctx).
		Model(&api).
		Clauses(clause.Returning{}).
		Where("id = ?", apiID).
		Update("status", model.ApiRunningStatus)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return api, consts.ErrApiNotfound
		}
		log.WithFields(log.Fields{"message": tx.Error}).Error("Create")
		return api, consts.ErrInternalError
	}

	if tx.RowsAffected < 1 {
		return api, consts.ErrApiNotfound
	}

	return api, nil
}

func (r *apiRepository) Stop(ctx context.Context, apiID uint) (model.Api, error) {
	api := model.Api{}
	tx := r.db.WithContext(ctx).
		Model(&api).
		Clauses(clause.Returning{}).
		Where("id = ?", apiID).
		Update("status", model.ApiStopedStatus)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return api, consts.ErrApiNotfound
		}
		log.WithFields(log.Fields{"message": tx.Error}).Error("Create")
		return api, consts.ErrInternalError
	}
	if tx.RowsAffected < 1 {
		return api, consts.ErrApiNotfound
	}

	return api, nil
}

func (r *apiRepository) Delete(ctx context.Context, apiID uint) error {
	tx := r.db.WithContext(ctx).
		Where("id = ?", apiID).
		Delete(&model.Api{})
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return consts.ErrApiNotfound
		}
		log.WithFields(log.Fields{"message": tx.Error}).Error("Create")
		return consts.ErrInternalError
	}

	if tx.RowsAffected < 1 {
		return consts.ErrApiNotfound
	}

	return nil
}

func New(db database.IDatabase) repository.ApiRepository {
	return &apiRepository{
		db: db.GetConnection(),
	}
}
