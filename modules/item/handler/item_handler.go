package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"todolist/common"
	"todolist/common/app_error"
	"todolist/modules/item/business"
	"todolist/modules/item/model"
)

type itemHandler struct {
	service business.IItemService
}

func NewItemHandler(service business.IItemService) *itemHandler {
	return &itemHandler{service: service}
}

func (handler *itemHandler) CreateItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))

			return
		}

		if err := handler.service.CreateItem(ctx, &data); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		uri := fmt.Sprintf("%s/items/%d", os.Getenv("DOMAIN"), data.Id)
		ctx.Header("Location", uri)
		ctx.Status(http.StatusCreated)
	}
}

func (handler *itemHandler) GetItemById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}

		data, serviceErr := handler.service.FindItemById(ctx, id)

		if serviceErr != nil {
			var appErr *app_error.AppError
			errors.As(serviceErr, &appErr)
			ctx.JSON(appErr.StatusCode, appErr)
			return
		}

		ctx.JSON(http.StatusOK, *data)
	}
}

func (handler *itemHandler) UpdateItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}

		if err := handler.service.UpdateItem(ctx, id, data); err != nil {
			var appErr *app_error.AppError
			errors.As(err, &appErr)
			ctx.JSON(appErr.StatusCode, appErr)
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}

func (handler *itemHandler) DeleteItem() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := handler.service.DeleteItem(ctx, id); err != nil {
			var appErr *app_error.AppError
			errors.As(err, &appErr)
			ctx.JSON(appErr.StatusCode, appErr)
			return

		}

		ctx.Status(http.StatusNoContent)
	}
}

func (handler *itemHandler) GetItems() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var p common.Paging
		if err := ctx.ShouldBind(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}
		p.Process()

		var f model.ItemFilter
		if err := ctx.ShouldBind(&f); err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
		}

		data, err := handler.service.FindItems(ctx, &p, &f)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(data, p, f))
	}
}

func (handler *itemHandler) SetupRoute(group *gin.RouterGroup) {
	group.POST("", handler.CreateItem())
	group.GET("/:id", handler.GetItemById())
	group.PATCH("/:id", handler.UpdateItem())
	group.DELETE("/:id", handler.DeleteItem())
	group.GET("", handler.GetItems())
}
