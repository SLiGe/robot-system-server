package v1

type SignInForQqRequest struct {
	QQ     string `json:"qq"`
	Points int64  `json:"points"`
}

type QuerySignInDataRequest struct {
	Id int64  `json:"id"`
	QQ string `json:"qq"`
}

type AddSignPointsRequest struct {
	QQ     string `json:"qq"`
	Points int64  `json:"points"`
}

type SignInDataResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    SignInDetailResponse `json:"dataResponse"`
}
type SignInDetailResponse struct {
	Qq           string `json:"qq"`
	Points       int64  `json:"points"`
	MonthDay     int64  `json:"monthDay"`
	TotalDay     int64  `json:"totalDay"`
	CurrentLevel string `json:"currentLevel"`
	TodayMsg     string `json:"todayMsg"`
	Ranking      int64  `json:"ranking"`
}
