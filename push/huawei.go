package push

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"
)

var (
	secret       string
	push         string
	sessionToken string
)

type HuaWeiPush struct {
}

/**
 * 初始化应用的服务端推送key
 */
func (hw HuaWeiPush) Init(_secret string, _push string) {
	secret = _secret
	push = _push
}

/**
 * SB华为，优秀的(坑)
 * url编码:
 *        Ctx
 *         token
 *         payload
 */
func (hw HuaWeiPush) Push(title string, content string, users []string) {
	appId := StringToInt(push)

	btsCtx, err := json.Marshal(NspCtx{Ver: 1, AppId: appId})
	if err != nil {
		log.Println(err)
	}
	btsCtxCode := url.QueryEscape(string(btsCtx))
	accessTokenUrl := fmt.Sprintf("https://api.push.hicloud.com/pushsend.do?nsp_ctx=%s", btsCtxCode)

	token := hw.huaWeiAccToken()
	token = url.QueryEscape(token)

	userBts, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	userBody := string(userBts)

	hwPush := HWPush{
		Hps: Hps{
			Msg: Msg{
				Type: 3,
				Body: Body{
					Title:   title,
					Content: content,
				},
				Action: Action{
					Type:  1,
					Param: Param{Intent: "##Intent;action=com.huawei.push.action.test;package=com.tokeninfo;end"},
				},
			},
		},
	}
	pushBts, err := json.Marshal(hwPush)
	payload := url.QueryEscape(string(pushBts))

	data := fmt.Sprintf("access_token=%s&nsp_svc=openpush.message.api.send&nsp_ts=%d&device_token_list=%s&payload=%s", token, time.Now().Unix(), userBody, payload)
	Post(accessTokenUrl, nil, data, nil)
}

/**
 * 可用Token
 */
func (hw HuaWeiPush) huaWeiAccToken() string {
	var value string
	if sessionToken == "" {
		value = hw.tokenRequest()
	} else {
		var access HWAccessToken
		err := json.Unmarshal([]byte(sessionToken), &access)
		if err != nil {
			log.Print(err)
		}
		if access.Expire.Sub(time.Now()) > time.Duration(0) {
			value = access.Token
		} else {
			value = hw.tokenRequest()
		}
	}
	return value
}

/**
 * 获取华为(服务端) 可用Token
 */
func (hw HuaWeiPush) tokenRequest() string {
	accessTokenUrl := "https://login.cloud.huawei.com/oauth2/v2/token"

	params := make(map[string]string)
	params["grant_type"] = "client_credentials"
	params["client_secret"] = secret
	params["client_id"] = push

	var result HWTokenResult
	PostForm(accessTokenUrl, nil, params, &result)

	accessToken := result.AccessToken
	expireTime := time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	saveToken := HWAccessToken{
		Token:  accessToken,
		Expire: expireTime,
	}

	value, err := json.Marshal(saveToken)
	if err != nil {
		log.Print("[tokenRequest] ", err)
	}

	sessionToken = string(value)
	return accessToken
}
