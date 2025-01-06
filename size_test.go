package bucket

import "testing"

func Test_getBucketSize(t *testing.T) {
	tests := []struct {
		name      string
		args      option
		totalItem int
		want      int
	}{
		{
			name: "equal amount",
			args: option{
				MinSize:      1,
				MaxSize:      10,
				ProcessCount: 1,
			},
			totalItem: 10,
			want:      10,
		},
		{
			name: "3 of 10",
			args: option{
				MinSize:      1,
				MaxSize:      10,
				ProcessCount: 3,
			},
			totalItem: 10,
			want:      4,
		},
		{
			name: "maxSize 2",
			args: option{
				MinSize:      1,
				MaxSize:      2,
				ProcessCount: 3,
			},
			totalItem: 10,
			want:      2,
		},
		{
			name: "minSize 100",
			args: option{
				MinSize:      100,
				ProcessCount: 20,
			},
			totalItem: 1000,
			want:      100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := bucketSizer(tt.args)
			if got := size(tt.totalItem); got != tt.want {
				t.Errorf("getBucketSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
