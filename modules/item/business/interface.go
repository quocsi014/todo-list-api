package business

import (
	"github.com/gin-gonic/gin"
	"todolist/common"
	"todolist/modules/item/model"
)

type IItemRepository interface {
	InsertItem(ctx *gin.Context, data *model.TodoItemCreation) error
	GetItemById(ctx *gin.Context, ID int) (*model.TodoItem, error)
	GetItems(ctx *gin.Context, paging *common.Paging, filter *model.ItemFilter) ([]model.TodoItem, error)
	UpdateItem(ctx *gin.Context, id int, update model.TodoItemUpdate) error
	DeleteItem(ctx *gin.Context, ID int) error
}

type IItemService interface {
	CreateItem(ctx *gin.Context, data *model.TodoItemCreation) error
	FindItemById(ctx *gin.Context, id int) (*model.TodoItem, error)
	FindItems(ctx *gin.Context, paging *common.Paging, filter *model.ItemFilter) ([]model.TodoItem, error)
	UpdateItem(ctx *gin.Context, id int, item model.TodoItemUpdate) error
	DeleteItem(ctx *gin.Context, Id int) error
}
