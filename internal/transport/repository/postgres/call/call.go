package call

import (
	"context"
	"github.com/apex/log"
	"github.com/logicalangel/tashil_test/internal/consts"
	"github.com/logicalangel/tashil_test/internal/model"
	"github.com/logicalangel/tashil_test/internal/transport/database"
	"github.com/logicalangel/tashil_test/internal/transport/repository"
	"gorm.io/gorm"
)

type callRepository struct {
	db *gorm.DB
}

func (r *callRepository) Create(ctx context.Context, apiID uint, status uint, time uint) (model.Call, error) {
	call := model.Call{
		ApiId:  apiID,
		Status: status,
		Time:   time,
	}

	tx := r.db.WithContext(ctx).Create(&call)
	if tx.Error != nil {
		log.WithFields(log.Fields{"message": tx.Error}).Error("CreateCall")
		return call, consts.ErrInternalError
	}

	return call, nil
}

func (r *callRepository) Get(ctx context.Context, apiID uint) ([]model.Call, error) {
	calls := []model.Call{}
	tx := r.db.WithContext(ctx).
		Where("api_id = ?", apiID).
		Order("updated_at DESC").
		Find(&calls)
	if tx.Error != nil {
		log.WithFields(log.Fields{"message": tx.Error}).Error("Get")
		return calls, consts.ErrInternalError
	}

	return calls, nil
}
func New(db database.IDatabase) repository.CallRepository {
	return &callRepository{
		db: db.GetConnection(),
	}
}
