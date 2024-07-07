package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func SuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{Data: data, Paging: nil, Filter: nil}
}
