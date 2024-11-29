package bucket

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Bucket[T any] struct {
	// ProcessCount is the maximum number of processes that will be run concurrently.
	//   - Default is 1, which means the data will be processed sequentially.
	ProcessCount int

	// MaxSize is the maximum size of the bucket, bucket size calculated by len(data) / ProcessCount.
	//   - Default no limit.
	MaxSize int
	// MinSize is the minimum size of the bucket, if the bucket size is less than MinSize, it will be set to MinSize.
	//   - Default is 1, which means the minimum size of the bucket is 1.
	MinSize int

	// Callback is the function that will be called for each bucket.
	Callback func(context.Context, []T) error
}

func New[T any](fn func(context.Context, []T) error, opts ...Option) *Bucket[T] {
	o := apply(opts)

	return &Bucket[T]{
		ProcessCount: o.ProcessCount,
		MaxSize:      o.MaxSize,
		MinSize:      o.MinSize,
		Callback:     fn,
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

	processCount := b.ProcessCount
	bucketSize := getBucketSize(len(data), b.MinSize, b.MaxSize, processCount)

	if processCount == 1 {
		// bucketing data and call function
		for i := 0; i < len(data); i += bucketSize {
			index := i

			if err := b.Callback(ctx, data[index:min(index+bucketSize, len(data))]); err != nil {
				return err
			}
		}

		return nil
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(processCount)

	// bucketing data and call function
	for i := 0; i < len(data); i += bucketSize {
		index := i

		g.Go(func() error {
			return b.Callback(ctx, data[index:min(index+bucketSize, len(data))])
		})
	}

	return g.Wait()
}

func getBucketSize(totalItem, minSize, maxSize, processCount int) int {
	bucketSize := totalItem / processCount

	if totalItem%processCount != 0 {
		bucketSize++
	}

	if bucketSize < minSize {
		bucketSize = minSize
	} else if maxSize > 0 && bucketSize > maxSize {
		bucketSize = maxSize
	}

	return bucketSize
}
