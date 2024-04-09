package v1

type SignInForQqRequest struct {
	QQ     string `json:"qq"`
	Points int64  `json:"points"`
}

type QuerySignInDataRequest struct {
	Id int    `json:"id"`
	QQ string `json:"qq"`
}

type AddSignPointsRequest struct {
	QQ     string `json:"qq"`
	Points int    `json:"points"`
}
