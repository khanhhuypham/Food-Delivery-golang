package restaurant_dto

type Statistic struct {
	Total         int64 `json:"total"`
	TotalActive   int64 `json:"total_active"`
	TotalInActive int64 `json:"total_inactive"`
}
