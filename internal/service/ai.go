package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	v1 "robot-system-server/api/v1"
)

type AiService interface {
	Poem(req *v1.PoemRequest)
}

func NewAiService(service *Service) AiService {
	return &aiService{
		Service: service,
	}
}

type aiService struct {
	*Service
}

var (
	poemUrl = "https://v2.alapi.cn/api/ai/poem?token=vn5IJzE4ZEgoYDVg&keyword=%s&num=%d&type=%d&rhyme=%d"
)

func (s *aiService) Poem(req *v1.PoemRequest) {
	reqUrl := fmt.Sprintf(poemUrl, req.Keyword, req.Num, req.Type, req.Rhyme)
	response, err := http.Get(reqUrl)
	if err != nil {
		return
	}
	var poem v1.PoemResponse
	if err = json.NewDecoder(response.Body).Decode(&poem); err != nil {
		return
	}
	//请求过多
	if poem.Code == 429 {

	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

}
