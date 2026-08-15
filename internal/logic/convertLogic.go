// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shortener/model"
	"shortener/pkg/base62"
	"shortener/pkg/md5"
	"shortener/pkg/urltool"

	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/pkg/connect"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// todo: add your logic here and delete this line

	//判断长链接是否有效
	if connect.Get(req.LongUrl) != true {
		return nil, errors.New("链接无效")
	}

	//判读长链接是否已经转过了
	urlMd5 := md5.Sum([]byte(req.LongUrl))
	result, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: urlMd5, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("该链接已被转为短链:%s", result.Surl.String)
		}
		logx.Error("svcCtx.ShortUrlModel.FindOneByMd5 falied!", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	//判断是否这个长链接是一个被存在的短链接,避免循环

	baseurl, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Error("urltool.GetBasePath falied!", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	su, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: baseurl, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("%s:该链接已为短链接", su.Surl.String)
		}
		logx.Error("svcCtx.ShortUrlModel.FindOneBySurl falied!", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	var (
		seq   uint64
		short string
	)

	for {
		seq, err = l.svcCtx.Sequence.Next(l.ctx)
		if err != nil {
			logx.Error("svcCtx.Sequence.Next falied!", logx.LogField{Key: "err", Value: err.Error()})
			return nil, err
		}
		short = base62.Encode(seq)

		if _, ok := l.svcCtx.ShortUrlBlackMap[short]; !ok {
			break
		}
	}
	//存储
	if _, err = l.svcCtx.ShortUrlModel.Insert(l.ctx, &model.ShortUrlMap{
		Lurl: sql.NullString{String: req.LongUrl, Valid: true},
		Md5:  sql.NullString{String: urlMd5, Valid: true},
		Surl: sql.NullString{String: short, Valid: true},
	}); err != nil {
		logx.Error("l.svcCtx.ShortUrlModel.Insert falied!", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	//加入到布隆过滤器中
	err = l.svcCtx.Filter.AddCtx(l.ctx,[]byte(short))
	if err!=nil{
		logx.Error("l.svcCtx.Filter.Add falied!", logx.LogField{Key: "err", Value: err.Error()})
	}

	shortUrl := l.svcCtx.Config.ShortDomain + "/" + short
	return &types.ConvertResponse{
		ShortUrl: shortUrl,
	}, nil
}
