package bucket

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Bucket[T any] struct {
	// callback is the function that will be called for each bucket.
	callback func(context.Context, []T) error
	// size return bucket size with the given total item.
	size         func(int) int
	processCount int
}

func New[T any](fn func(context.Context, []T) error, opts ...Option) *Bucket[T] {
	o := apply(opts)
	size := bucketSizer(o)

	return &Bucket[T]{
		callback:     fn,
		size:         size,
		processCount: o.ProcessCount,
	}
}

// Process will process the data concurrently based on the ProcessCount.
//   - If the data is empty, it will return nil.
//   - If the ProcessCount is 1, it will process the data sequentially.
//   - If the ProcessCount is more than 1, it will process the data concurrently.
//   - The bucket size will be calculated by len(data) / ProcessCount, if not divisible, it will be rounded up. 10/3 -> 4 that means 4,4,2 buckets.
//   - If the bucket size is less than MinSize, it will be set to MinSize.
//   - If the bucket size is more than MaxSize, it will be set to MaxSize.
//   - The function will return an error if any of the bucket processing returns an error.
func (b *Bucket[T]) Process(ctx context.Context, data []T) error {
	if len(data) == 0 {
		return nil
	}

	bucketSize := b.size(len(data))

	if b.processCount == 1 || len(data) <= bucketSize {
		// bucketing data and call function
		for i := 0; i < len(data); i += bucketSize {
			index := i

			ctxProcess := withIndex(ctx, index)

			if err := b.callback(ctxProcess, data[index:min(index+bucketSize, len(data))]); err != nil {
				return err
			}
		}

		return nil
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(b.processCount)

	// bucketing data and call function
	for i := 0; i < len(data); i += bucketSize {
		index := i

		ctxProcess := withIndex(ctx, index)

		g.Go(func() error {
			return b.callback(ctxProcess, data[index:min(index+bucketSize, len(data))])
		})
	}

	return g.Wait()
}
