package logic

import "robot-system-server/pkg/log"

type Logic struct {
	logger *log.Logger
}

func NewLogic(logger *log.Logger) *Logic {
	return &Logic{
		logger: logger,
	}
}
