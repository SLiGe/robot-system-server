package v1

type PoemRequest struct {
	Keyword string `json:"keyword"`
	Num     int    `json:"num"`
	Type    int    `json:"type"`
	Rhyme   int    `json:"rhyme"`
}

type PoemResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Keyword string   `json:"keyword"`
		Poem    string   `json:"poem"`
		List    []string `json:"list"`
	} `json:"data"`
	Time  int    `json:"time"`
	Usage int    `json:"usage"`
	LogId string `json:"log_id"`
}
