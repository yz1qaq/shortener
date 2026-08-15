package sequence

import "context"

type Sequence interface {
	Next(ctx context.Context) (uint64, error)
}
