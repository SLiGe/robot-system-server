package service

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/logic"
	"robot-system-server/internal/model"
	"strings"
)

type AiService interface {
	Poem(ctx *gin.Context, req *v1.PoemRequest) (v1.PoemData, error)
}

func NewAiService(service *Service, aiLogic logic.AiLogic) AiService {
	return &aiService{
		Service: service,
		aiLogic: aiLogic,
	}
}

type aiService struct {
	*Service
	aiLogic logic.AiLogic
}

var (
	poemUrl   = "https://v2.alapi.cn/api/ai/poem?token=vn5IJzE4ZEgoYDVg&keyword=%s&num=%d&type=%d&rhyme=%d"
	txPoemUrl = "https://apis.tianapi.com/cangtoushi/index?key=a3694ec69e8165a2a38ba20299b1d0ca&word=%s&len=%d&type=%d&pat=%d"
)

func (s *aiService) Poem(ctx *gin.Context, req *v1.PoemRequest) (v1.PoemData, error) {
	reqUrl := fmt.Sprintf(poemUrl, req.Keyword, req.Num, req.Type, req.Rhyme)
	encodeReqUrl, _ := netutil.EncodeUrl(reqUrl)
	response, err := http.Get(encodeReqUrl)
	if err != nil {
		return v1.PoemData{}, err
	}
	if response.StatusCode != 200 {
		return s.GetTxPoem(req)
	}
	var poem v1.PoemResponse
	if err = json.NewDecoder(response.Body).Decode(&poem); err != nil {
		return v1.PoemData{}, err
	}
	//请求过多
	if poem.Code == 429 {
		return s.GetTxPoem(req)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	var saveArray = make([]*model.QrHiddenPoem, 1)
	saveArray[0] = &model.QrHiddenPoem{
		Poem:        &poem.Data.Poem,
		PoemKeyword: &req.Keyword,
		PoemNum:     int32(req.Num),
		PoemType:    int32(req.Type),
		PoemRhyme:   int32(req.Rhyme),
	}
	s.aiLogic.SaveHiddenPoem(saveArray)
	return poem.Data, nil
}

func (s *aiService) GetTxPoem(req *v1.PoemRequest) (v1.PoemData, error) {
	reqUrl := fmt.Sprintf(txPoemUrl, req.Keyword, req.Num, req.Type, req.Rhyme)
	encodeReqUrl, _ := netutil.EncodeUrl(reqUrl)
	response, err := http.Get(encodeReqUrl)
	if err != nil {
		return v1.PoemData{}, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	var poem v1.TxPoemResponse
	if err = json.NewDecoder(response.Body).Decode(&poem); err != nil {
		return v1.PoemData{}, err
	}
	if len(poem.Result.List) == 0 {
		return v1.PoemData{}, nil
	}
	poemFirst := poem.Result.List[0].Content
	poemArray := strings.Split(poemFirst, "。")
	if poemArray[len(poemArray)-1] == "" {
		poemArray = poemArray[:len(poemArray)-1]
	}
	var saveArray = make([]*model.QrHiddenPoem, len(poem.Result.List))
	for i, ele := range poem.Result.List {
		saveArray[i] = &model.QrHiddenPoem{
			Poem:        &ele.Content,
			PoemKeyword: &req.Keyword,
			PoemNum:     int32(req.Num),
			PoemType:    int32(req.Type),
			PoemRhyme:   int32(req.Rhyme),
		}
	}
	s.aiLogic.SaveHiddenPoem(saveArray)
	return v1.PoemData{
		Keyword: req.Keyword,
		Poem:    poemFirst,
		List:    poemArray,
	}, nil
}

func (s *aiService) GetOpenAiPoem(ctx *gin.Context, keyword string) (v1.PoemData, error) {
	clientConfig := openai.DefaultConfig("sb-a351f2748d068e0970b5a1bed50a2150")
	clientConfig.BaseURL = "https://api.openai-sb.com/v1"
	client := openai.NewClientWithConfig(clientConfig)
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		MaxTokens:   1024,
		Temperature: 0.7,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "请以“我”、“想”、“你”为第一个字写一首押韵的藏头古诗",
			},
		},
	})
	if err != nil {
		return v1.PoemData{}, err
	}
	if len(resp.Choices) == 0 {
		return v1.PoemData{}, nil
	}
	if resp.Choices[0].Message.Content == "" {
		return v1.PoemData{}, nil
	}
	poemStr := resp.Choices[0].Message.Content
	poemList := strings.Split(poemStr, "\n")
	return v1.PoemData{
		Keyword: keyword,
		Poem:    poemStr,
		List:    poemList,
	}, nil
}
