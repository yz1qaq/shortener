package sequence

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Redis struct {
	client *redis.Redis
	key    string
}

func NewRedis(client *redis.Redis, key string) *Redis {
	return &Redis{
		client: client,
		key:    key,
	}
}

func (r *Redis) Next(ctx context.Context) (uint64, error) {
	value, err := r.client.IncrCtx(ctx, r.key)
	if err != nil {
		logx.Errorw("r.client.IncrCtx", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	return uint64(value), nil
}
