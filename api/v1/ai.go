package v1

type PoemRequest struct {
	Keyword string `form:"keyword"`
	Num     int    `form:"num"`
	Type    int    `form:"type"`
	Rhyme   int    `form:"rhyme"`
}

type PoemData struct {
	Keyword string   `json:"keyword"`
	Poem    string   `json:"poem"`
	List    []string `json:"list"`
}
type PoemResponse struct {
	Code  int      `json:"code,exclude"`
	Msg   string   `json:"msg,exclude"`
	Data  PoemData `json:"data"`
	Time  int      `json:"time,exclude"`
	Usage int      `json:"usage,exclude"`
	LogId string   `json:"log_id,exclude"`
}

type TxPoemResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		List []struct {
			Content string `json:"content"`
		} `json:"list"`
	} `json:"result"`
}
