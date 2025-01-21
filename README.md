# bucket 🪣

[![License](https://img.shields.io/github/license/worldline-go/bucket?color=red&style=flat-square)](https://raw.githubusercontent.com/worldline-go/bucket/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/worldline-go_bucket?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=worldline-go_bucket)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/worldline-go/bucket/test.yml?branch=main&logo=github&style=flat-square&label=ci)](https://github.com/worldline-go/bucket/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/worldline-go/bucket?style=flat-square)](https://goreportcard.com/report/github.com/worldline-go/bucket)
[![Go PKG](https://raw.githubusercontent.com/worldline-go/guide/main/badge/custom/reference.svg)](https://pkg.go.dev/github.com/worldline-go/bucket)

Bucket call a function with limited size of gorutine and data count.  
Size of data calculated by `len(data) / processCount` with minimum and maximum size configurations.

```sh
go get github.com/worldline-go/bucket
```

> This package based on [golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).

## Usage

Bucket accepts a function with signature `func(context.Context, []T) error`.

```go
processBucket := bucket.New(process,
    bucket.WithProcessCount(4),
    bucket.WithMinSize(2),
    bucket.WithMaxSize(100),
)

// 10 items -> 10/4 -> 3 items per bucket, 3,3,3,1 will be processed in 4 gorutine
data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// processBucket returns error of process function
if err := processBucket.Process(context.Background(), data); err != nil {
    return err
}
```

Get original begin index for process batch with `bucket.CtxIndex(ctx)`.
