package bucket

import (
	"testing"
)

func Test_getBucketSize(t *testing.T) {
	type args struct {
		totalItem    int
		minSize      int
		maxSize      int
		processCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal amount",
			args: args{
				totalItem:    10,
				minSize:      1,
				maxSize:      10,
				processCount: 1,
			},
			want: 10,
		},
		{
			name: "3 of 10",
			args: args{
				totalItem:    10,
				minSize:      1,
				maxSize:      10,
				processCount: 3,
			},
			want: 3,
		},
		{
			name: "maxSize 2",
			args: args{
				totalItem:    10,
				minSize:      1,
				maxSize:      2,
				processCount: 3,
			},
			want: 2,
		},
		{
			name: "minSize 100",
			args: args{
				totalItem:    1000,
				minSize:      100,
				processCount: 20,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBucketSize(tt.args.totalItem, tt.args.minSize, tt.args.maxSize, tt.args.processCount); got != tt.want {
				t.Errorf("getBucketSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
