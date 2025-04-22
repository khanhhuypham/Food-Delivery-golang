package constant

type Cuisine_Type string

const (
	CUISINE_THAI        = "Thai"
	CUISINE_VIETNAM     = "VietNamese"
	CUISINE_SOUTH_KOREA = "South Korean"
	CUISINE_JAPANESE    = "Japanese"
	CUISINE_CHINESE     = "Chinese"
)

type Recommendation_Type string

const (
	RECOMMEND_TYPE_BEST     = "best food"
	RECOMMEND_TYPE_POPULAR  = "most popular"
	RECOMMEND_TYPE_SUGGESTE = "most recommended"
)

type Interaction_type int

const (
	INTERACT_TYPE_VIEW  = 1
	INTERACT_TYPE_LIKE  = 2
	INTERACT_TYPE_SHARE = 3
)
