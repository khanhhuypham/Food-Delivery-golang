package rating_dto

type RatingDTO struct {
	Like  bool `json:"like"`
	Score int  `json:"score"`
}
