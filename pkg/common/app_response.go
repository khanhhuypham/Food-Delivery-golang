package common

import "net/http"

type AppReponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationResult struct {
	Data interface{} `json:"data"`
	Paging
}

func Response(data interface{}) *AppReponse {
	return &AppReponse{Data: data}
}

func ResponseWithPaging(data interface{}, paging Paging) *AppReponse {
	return &AppReponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: PaginationResult{
			Data:   data,
			Paging: paging,
		},
	}

}
