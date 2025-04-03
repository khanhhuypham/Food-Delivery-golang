package common

type GeneralReponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func Response(data interface{}) *GeneralReponse {
	return &GeneralReponse{Data: data}
}

func ResponseWithPaging(data interface{}, paging interface{}) *GeneralReponse {
	return &GeneralReponse{Data: data}
}
