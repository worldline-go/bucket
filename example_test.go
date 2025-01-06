package bucket_test

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/rakunlabs/bucket"
)

func ExampleBucket() {
	totalCount := int64(0)
	process := func(ctx context.Context, data []int) error {
		// do something with data
		// fmt.Println(data)

		for _, v := range data {
			atomic.AddInt64(&totalCount, int64(v))
		}

		return nil
	}

	processBucket := bucket.New(process,
		bucket.WithProcessCount(4),
		bucket.WithMinSize(2),
		bucket.WithMaxSize(100),
		// or give with config
		bucket.Config{
			ProcessCount: 4,
			MinSize:      2,
			MaxSize:      100,
		}.ToOption(),
	)

	// 10 items -> 10/4 -> 3 items per bucket
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// process data
	if err := processBucket.Process(context.Background(), data); err != nil {
		fmt.Println(err)
	}

	fmt.Println(totalCount)
	// Output:
	// 55
}
