package v1

import "robot-system-server/internal/model"

type GetFortuneRequest struct {
	QQ         string `json:"qq"`
	IsOne      int    `json:"isOne"`
	IsGroup    int    `json:"isGroup"`
	IsIntegral int    `json:"isIntegral"`
	GroupNum   string `json:"groupNum"`
}

type GetFortuneResponse struct {
	Status       int              `json:"status"`
	Message      string           `json:"message"`
	DataResponse *model.QrFortune `json:"dataResponse"`
}
