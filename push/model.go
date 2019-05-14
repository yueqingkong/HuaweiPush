package push

import "time"

type NspCtx struct {
	Ver   int `json:"ver"`
	AppId int `json:"appId"`
}

type HWAccessToken struct {
	Token  string    `json:"access_token"`
	Expire time.Time `json:"expire_time"`
}

type Body struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

type Param struct {
	Intent string `json:"intent"`
}

type Action struct {
	Type  int   `json:"type"`
	Param Param `json:"param"`
}

type Msg struct {
	Type   int    `json:"type"`
	Body   Body   `json:"body"`
	Action Action `json:"action"`
}

type Hps struct {
	Msg Msg `json:"msg"`
}
type HWPush struct {
	Hps Hps `json:"hps"`
}

type HWTokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
