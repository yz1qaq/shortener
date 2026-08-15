// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ShortUrlDB struct {
		DSN string
	}
	Sequence struct {
		Type  string
		MySQL struct {
			DSN string
		}
		Redis redis.RedisKeyConf
	}
	BaseString string

	ShortUrlBlackList []string

	ShortDomain string

	CacheRedis cache.CacheConf
}
