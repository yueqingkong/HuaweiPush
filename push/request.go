package push

import (
	"encoding/json"
	"github.com/go-resty/resty"
	"log"
	"time"
)

// POST
func Post(url string, headers map[string]string, params interface{}, inter interface{}) {
	if headers == nil {
		headers = make(map[string]string)
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	resp, err := resty.SetTimeout(time.Minute * 1).R().
		SetHeaders(headers).
		SetBody(params).
		Post(url)
	log.Print("[Post]", resp.String())

	if inter != nil {
		err = json.Unmarshal(resp.Body(), &inter)
		if err != nil {
			log.Print("[net][PostForm]", err, url)
		}
	}
}

// POST 表单
func PostForm(url string, headers map[string]string, params map[string]string, inter interface{}) {
	if headers == nil {
		headers = make(map[string]string)
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	resp, err := resty.SetTimeout(time.Minute * 1).R().
		SetHeaders(headers).
		SetFormData(params).
		Post(url)

	if err != nil {
		log.Print("[PostForm]", err, "[url]", url)
	}
	log.Print("[PostForm]",resp.String())

	if inter != nil {
		err = json.Unmarshal(resp.Body(), &inter)
		if err != nil {
			log.Print("[net][PostForm]", err, url)
		}
	}
}
