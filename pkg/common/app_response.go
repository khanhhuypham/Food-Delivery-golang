package common

type AppReponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func Response(data interface{}) *AppReponse {
	return &AppReponse{Data: data}
}

func ResponseWithPaging(data interface{}, paging interface{}) *AppReponse {
	return &AppReponse{Data: data}
}
