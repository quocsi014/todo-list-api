package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"todolist/common"
	"todolist/common/app_error"
	"todolist/modules/item/model"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (repo ItemRepository) InsertItem(ctx *gin.Context, item *model.TodoItemCreation) error {
	if err := repo.db.Create(item).Error; err != nil {
		return app_error.ErrDB(err)
	}
	return nil
}

func (repo ItemRepository) GetItemById(ctx *gin.Context, id int) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := repo.db.Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_error.ErrEntityNotFound(model.TodoItem{}.EntityName(), err)
		}
		return nil, app_error.ErrDB(err)
	}

	return &data, nil
}

func (repo ItemRepository) GetItems(ctx *gin.Context, paging *common.Paging, filter *model.ItemFilter) ([]model.TodoItem, error) {
	var totalItem int64
	var result []model.TodoItem
	var db *gorm.DB = repo.db

	//filter handle
	if filter != nil {
		if filter.Status != nil {
			db = db.Where("status = ?", *(filter.Status))
		}
		if filter.Title != nil {
			db = db.Where("title like ?", "%"+*(filter.Title)+"%")
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&totalItem).Error; err != nil {
		return nil, err
	}

	paging.TotalPage = int64(math.Ceil(float64(totalItem) / float64(paging.Limit)))

	if err := db.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&result).Error; err != nil {
		return nil, app_error.ErrDB(err)
	}

	return result, nil
}

func (repo ItemRepository) UpdateItem(ctx *gin.Context, id int, item model.TodoItemUpdate) error {
	updateResult := repo.db.Where("id = ?", id).Updates(item)
	if err := updateResult.Error; err != nil {
		return app_error.ErrDB(err)
	}
	if updateResult.RowsAffected == 0 {
		return app_error.ErrEntityNotFound(model.TodoItem{}.EntityName(), app_error.ErrRecordNotFound)
	}
	return nil
}

func (repo ItemRepository) DeleteItem(ctx *gin.Context, id int) error {
	delResult := repo.db.Where("id = ?", id).Delete(&model.TodoItem{})

	if err := delResult.Error; err != nil {
		return app_error.ErrDB(err)
	}

	if delResult.RowsAffected == 0 {
		return app_error.ErrEntityNotFound(model.TodoItem{}.EntityName(), app_error.ErrRecordNotFound)
	}

	return nil
}
