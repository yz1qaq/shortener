// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"shortener/internal/config"

	"shortener/model"
	"shortener/sequence"
	"strings"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	ShortUrlModel    model.ShortUrlMapModel
	Sequence         sequence.Sequence
	ShortUrlBlackMap map[string]Empty

	Filter *bloom.Filter
}

type Empty struct{}

func NewServiceContext(c config.Config) *ServiceContext {
	var seq sequence.Sequence

	switch strings.ToLower(strings.TrimSpace(c.Sequence.Type)) {
	case "mysql":
		seq = sequence.NewMySQL(c.Sequence.MySQL.DSN)

	case "redis":
		rds := redis.MustNewRedis(c.Sequence.Redis.RedisConf)
		seq = sequence.NewRedis(rds, c.Sequence.Redis.Key)

	default:
		panic("unsupported sequence type: " + c.Sequence.Type)
	}

	m := make(map[string]Empty, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = Empty{}
	}

	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) { r.Type = redis.NodeType })
	bitsSet := bloom.New(store, "bloom_filter", 20*(1<<20))

	return &ServiceContext{
		Config: c,
		ShortUrlModel: model.NewShortUrlMapModel(
			sqlx.NewMysql(c.ShortUrlDB.DSN), c.CacheRedis),
		Sequence:         seq,
		ShortUrlBlackMap: m,
		Filter:           bitsSet,
	}
}
