package urltool

import (
	"errors"
	"net/url"
	"path"

	"github.com/zeromicro/go-zero/core/logx"
)

//获取长链接的最后一节内容
/*
Scheme   = "https"
Host     = "www.liwenzhou.com"
Path     = "/posts/go/golang-menu/"
RawQuery = "id=10&name=yz1"
Fragment = "chapter2"
*/
func GetBasePath(LongUrl string) (string, error) {
	u, err := url.Parse(LongUrl)
	if err != nil {
		logx.Error("url.Parse falied!", logx.LogField{Key: "lurl", Value: LongUrl}, logx.LogField{Key: "err", Value: err.Error()})
		return "", err
	}
	if len(u.Host) == 0 {
		return "", errors.New("host 为空!")
	}
	basePath := path.Base(u.Path)
	return basePath, nil
}

func GetScheme(LongUrl string) (string, error) {
	u, err := url.Parse(LongUrl)
	if err != nil {
		logx.Error("url.Parse falied!", logx.LogField{Key: "lurl", Value: LongUrl}, logx.LogField{Key: "err", Value: err.Error()})
		return "", err
	}
	if len(u.Host) == 0 {
		return "", errors.New("host 为空!")
	}
	Scheme := u.Scheme
	return Scheme, nil
}
