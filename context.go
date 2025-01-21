package bucket

import "context"

type ctxKey string

var IndexKey = ctxKey("index")

// CtxIndex returns the index of original list from the context.
//   - If the index is not found, it will return -1.
func CtxIndex(ctx context.Context) int {
	if v, ok := ctx.Value(IndexKey).(int); ok {
		return v
	}

	return -1
}

func withIndex(ctx context.Context, index int) context.Context {
	return context.WithValue(ctx, IndexKey, index)
}
