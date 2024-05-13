package logic

import (
	"errors"
	"gorm.io/gorm"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"strings"
	"sync"
)

type AiLogic interface {
	SaveHiddenPoem(poem []*model.QrHiddenPoem)
	SaveAlHiddenPoem(req *v1.PoemRequest, poem string)
	SaveTxHiddenPoem(req *v1.PoemRequest, res v1.TxPoemResponse)
}
type aiLogic struct {
	*Logic
	mu sync.Mutex
}

func NewAiLogic(logic *Logic) AiLogic {
	return &aiLogic{
		Logic: logic,
	}
}

func (l *aiLogic) SaveAlHiddenPoem(req *v1.PoemRequest, poem string) {
	var saveArray = make([]*model.QrHiddenPoem, 1)
	saveArray[0] = &model.QrHiddenPoem{
		Poem:        &poem,
		PoemKeyword: &req.Keyword,
		PoemNum:     int32(req.Num),
		PoemType:    int32(req.Type),
		PoemRhyme:   int32(req.Rhyme),
	}
	l.SaveHiddenPoem(saveArray)
}

func (l *aiLogic) SaveTxHiddenPoem(req *v1.PoemRequest, res v1.TxPoemResponse) {
	var saveArray = make([]*model.QrHiddenPoem, len(res.Result.List))
	for i := 0; i < len(res.Result.List); i++ {
		saveArray[i] = &model.QrHiddenPoem{
			Poem:        &res.Result.List[i].Content,
			PoemKeyword: &req.Keyword,
			PoemNum:     int32(req.Num),
			PoemType:    int32(req.Type),
			PoemRhyme:   int32(req.Rhyme),
		}
	}
	l.SaveHiddenPoem(saveArray)
}

func (l *aiLogic) SaveHiddenPoem(poem []*model.QrHiddenPoem) {
	user := "system"
	p := query.QrHiddenPoem
	l.mu.Lock()
	defer l.mu.Unlock()
	var saveArray []*model.QrHiddenPoem
	for _, hiddenPoem := range poem {
		*hiddenPoem.PoemKeyword = strings.TrimSpace(*hiddenPoem.PoemKeyword)
		*hiddenPoem.Poem = strings.TrimSpace(*hiddenPoem.Poem)
		_, err := p.Select(p.ID).Where(p.Poem.Eq(*hiddenPoem.Poem), p.PoemKeyword.Eq(*hiddenPoem.PoemKeyword), p.PoemType.Eq(hiddenPoem.PoemType), p.PoemNum.Eq(hiddenPoem.PoemNum), p.PoemRhyme.Eq(hiddenPoem.PoemRhyme)).First()
		hiddenPoem.CreateBy = &user
		hiddenPoem.UpdateBy = &user
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			saveArray = append(saveArray, hiddenPoem)
		}
	}
	if len(saveArray) > 0 {
		_ = p.Create(saveArray...)
	}
}
