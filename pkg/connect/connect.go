package connect

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var client = &http.Client{
	Transport: &http.Transport{DisableKeepAlives: true},
	Timeout:   2 * time.Second,
}

// 判断url是否能请求通
func Get(url string) bool {
	resp , err := client.Get(url); 
	if err != nil {
		logx.Errorw("connect client.Get failed", logx.LogField{Key: "err", Value: err.Error()})
		return false
	}
	resp.Body.Close()
	return resp.StatusCode==http.StatusOK
}
