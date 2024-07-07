package model

import (
	"errors"
	"strings"
	"time"
	"todolist/common"
	"todolist/common/app_error"
)

type TodoItem struct {
	common.SQLModel
	Title       string     `json:"title" gorm:"column:title"`
	ImageUrl    string     `json:"image_url,omitempty" gorm:"column:image_url"`
	Description string     `json:"description,omitempty" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}
func (TodoItem) EntityName() string { return "TodoItem" }

var (
	ErrTitleNull  = app_error.NewErrorResponse(errors.New("title is null"), "Title can not be null", "", "TODO_ITEM_TITLE_NULL")
	ErrTitleBlank = app_error.NewErrorResponse(errors.New("title is blank"), "Title can not be blank", "", "TODO_ITEM_TITLE_BLANK")
)

type TodoItemCreation struct {
	Id          int        `json:"-" gorm:"column:id"`
	Title       *string    `json:"title" gorm:"column:title"`
	Description string     `json:"description,omitempty" gorm:"column:description"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (item TodoItemCreation) Validate() error {
	if item.Title == nil {
		return ErrTitleNull
	}

	title := strings.TrimSpace(*(item.Title))
	if title == "" {
		return ErrTitleBlank
	}

	return nil
}

func (item TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       *string    `json:"title" gorm:"column:title"`
	ImageUrl    *string    `json:"image_url,omitempty" gorm:"column:image_url"`
	Description *string    `json:"description,omitempty" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
}

func (item TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
func (item TodoItemUpdate) Validate() error {
	if item.Title == nil {
		return nil
	}

	title := strings.TrimSpace(*(item.Title))
	if title == "" {
		return ErrTitleBlank
	}

	return nil
}
