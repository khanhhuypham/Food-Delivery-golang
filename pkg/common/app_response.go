package common

type AppReponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(data interface{}) *AppReponse {
	return &AppReponse{Data: data}
}

func ResponseWithPaging(data interface{}, paging interface{}) *AppReponse {
	return &AppReponse{Data: data}
}
