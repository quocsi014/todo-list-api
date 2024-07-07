package model

type ItemFilter struct {
	Status *string `json:"status,omitempty" form:"status"`
	Title  *string `json:"title,omitempty" form:"title"`
}
