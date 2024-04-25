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

func (r SignInDataResponse) Success() SignInDataResponse {
	r.Status = 200
	r.Message = "签到成功"
	return r
}

func (r SignInDataResponse) Fail() SignInDataResponse {
	r.Status = 500
	r.Message = "签到失败,请联系管理员"
	return r
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

type SignPerDayReq struct {
	Type    string `json:"type"`
	Qq      string `json:"qq"`
	Group   string `json:"group"`
	MsgType int    `json:"msgType"`
}

type SignPerDayRes struct {
	SignData any    `json:"signData"`
	ViewUrl  string `json:"viewUrl"`
	Image    string `json:"image"`
}
