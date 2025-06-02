package restaurant_dto

type Statistic struct {
	Total                    int64 `json:"total"`
	TotalOpen                int64 `json:"total_open"`
	TotalClosed              int64 `json:"total_closed"`
	TotalTemporarilyClosed   int64 `json:"total_temporarily_closed"`
	TotalLimitedAvailability int64 `json:"total_limited_availability"`
	TotalSuspended           int64 `json:"total_suspended"`
}
