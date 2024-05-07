package logic

import (
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
)

type AiLogic interface {
	SaveHiddenPoem(poem []*model.QrHiddenPoem)
}
type aiLogic struct {
	*Logic
}

func NewAiLogic(logic *Logic) AiLogic {
	return &aiLogic{
		Logic: logic,
	}
}

func (l *aiLogic) SaveHiddenPoem(poem []*model.QrHiddenPoem) {
	user := "system"
	for _, hiddenPoem := range poem {
		hiddenPoem.CreateBy = &user
		hiddenPoem.UpdateBy = &user
	}
	err := query.QrHiddenPoem.Create(poem...)
	if err != nil {
		return
	}
}
