package business

import (
	"github.com/gin-gonic/gin"
	"todolist/common"
	"todolist/modules/item/model"
)

type itemBus struct {
	repo IItemRepository
}

func NewItemBusiness(repository IItemRepository) *itemBus {
	return &itemBus{repo: repository}
}

func (bus itemBus) CreateItem(ctx *gin.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}
	return bus.repo.InsertItem(ctx, data)
}

func (bus itemBus) FindItemById(ctx *gin.Context, id int) (*model.TodoItem, error) {
	return bus.repo.GetItemById(ctx, id)
}

func (bus itemBus) FindItems(ctx *gin.Context, paging *common.Paging, filter *model.ItemFilter) ([]model.TodoItem, error) {
	return bus.repo.GetItems(ctx, paging, filter)
}

func (bus itemBus) UpdateItem(ctx *gin.Context, id int, data model.TodoItemUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	return bus.repo.UpdateItem(ctx, id, data)
}

func (bus itemBus) DeleteItem(ctx *gin.Context, id int) error {
	return bus.repo.DeleteItem(ctx, id)
}
