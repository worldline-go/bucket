package bucket_test

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/worldline-go/bucket"
)

func TestBucket(t *testing.T) {
	t.Run("one process count", func(t *testing.T) {
		v := 0
		b := bucket.New(func(ctx context.Context, t []int) error {
			for _, i := range t {
				v += i
			}

			return nil
		})

		if err := b.Process(context.Background(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}); err != nil {
			t.Fatal(err)
		}

		if v != 55 {
			t.Errorf("got %d, want 55", v)
		}
	})

	t.Run("nil data", func(t *testing.T) {
		v := 0
		b := bucket.New(func(ctx context.Context, t []int) error {
			for _, i := range t {
				v += i
			}

			return nil
		})

		if err := b.Process(context.Background(), nil); err != nil {
			t.Fatal(err)
		}

		if v != 0 {
			t.Errorf("got %d, want 0", v)
		}
	})

	t.Run("return error multi", func(t *testing.T) {
		v := int64(0)
		b := bucket.New(func(ctx context.Context, t []int) error {
			for _, i := range t {
				if i == 3 {
					return fmt.Errorf("some error")
				}

				atomic.AddInt64(&v, int64(i))
			}

			return nil
		}, bucket.WithProcessCount(10))

		if err := b.Process(context.Background(), []int{1, 2, 3, 4, 5}); err == nil {
			t.Fatal("want error, got nil")
		} else {
			if err.Error() != "some error" {
				t.Errorf("got %s, want some error", err.Error())
			}
		}
	})

	t.Run("return error", func(t *testing.T) {
		v := 0
		b := bucket.New(func(ctx context.Context, t []int) error {
			for _, i := range t {
				if i == 3 {
					return fmt.Errorf("some error")
				}

				v += i
			}

			return nil
		})

		if err := b.Process(context.Background(), []int{1, 2, 3, 4, 5}); err == nil {
			t.Fatal("want error, got nil")
		} else {
			if err.Error() != "some error" {
				t.Errorf("got %s, want some error", err.Error())
			}
		}
	})
}
