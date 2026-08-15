// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"
	"database/sql"
	"errors"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	bloom404 = errors.New("bloom:404")
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// todo: add your logic here and delete this line

	//布隆过滤器,不存在的短链接直接404即可
	exist, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("l.svcCtx.Filter.Exists failed!", logx.LogField{Key: "err", Value: err.Error()})
	} else if !exist {
		return nil, bloom404
	}

	result, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortUrl, Valid: true})
	if err != nil {
		logx.Errorw("l.svcCtx.ShortUrlModel.FindOneBySurl failed!", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	lUrl := result.Lurl
	return &types.ShowResponse{
		LongUrl: lUrl.String,
	}, nil
}
