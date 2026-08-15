package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShortUrlMapModel = (*customShortUrlMapModel)(nil)

type ShortUrlItem struct {
	ID   uint64 `db:"id"`
	SURL string `db:"surl"`
}

type (
	// ShortUrlMapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortUrlMapModel.
	ShortUrlMapModel interface {
		shortUrlMapModel
		FindSurlsAfterID(ctx context.Context, lastID uint64, pageSize int) ([]ShortUrlItem, error)
	}

	customShortUrlMapModel struct {
		*defaultShortUrlMapModel
	}
)

// NewShortUrlMapModel returns a model for the database table.
func NewShortUrlMapModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ShortUrlMapModel {
	return &customShortUrlMapModel{
		defaultShortUrlMapModel: newShortUrlMapModel(conn, c, opts...),
	}
}

// FindSurlsAfterID returns one page of valid short URLs after lastID.
func (m *customShortUrlMapModel) FindSurlsAfterID(
	ctx context.Context,
	lastID uint64,
	pageSize int,
) ([]ShortUrlItem, error) {
	if pageSize <= 0 {
		return nil, fmt.Errorf("page size must be greater than zero")
	}

	var items []ShortUrlItem
	query := fmt.Sprintf(`select id, surl
		from %s
		where id > ?
		  and is_del = 0
		  and surl is not null
		  and surl != ''
		order by id
		limit ?`, m.table)

	if err := m.QueryRowsNoCacheCtx(ctx, &items, query, lastID, pageSize); err != nil {
		return nil, err
	}

	return items, nil
}
