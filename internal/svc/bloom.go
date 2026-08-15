package svc

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

const bloomInitPageSize = 1000

// InitBloomFilter loads all existing valid short URLs into the bloom filter.
func (s *ServiceContext) InitBloomFilter(ctx context.Context) error {
	var (
		lastID uint64
		total  int
	)

	for {
		items, err := s.ShortUrlModel.FindSurlsAfterID(ctx, lastID, bloomInitPageSize)
		
		if err != nil {
			return fmt.Errorf("query short urls after id %d: %w", lastID, err)
		}

		if len(items) == 0 {
			logx.Infof("bloom filter initialization completed, total=%d", total)
			return nil
		}

		for _, item := range items {
			if err := s.Filter.AddCtx(ctx, []byte(item.SURL)); err != nil {
				return fmt.Errorf("add short url %q to bloom filter: %w", item.SURL, err)
			}
		}

		lastID = items[len(items)-1].ID
		total += len(items)
		logx.Infof("loaded %d short urls into bloom filter, lastID=%d", total, lastID)
	}
}
