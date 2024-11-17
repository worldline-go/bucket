package bucket

type option struct {
	ProcessCount int
	MaxSize      int
	MinSize      int
}

func (o *option) setDefault() {
	if o.ProcessCount <= 0 {
		o.ProcessCount = 1
	}

	if o.MinSize <= 0 {
		o.MinSize = 1
	}
}

type Option func(o *option)

func apply(opts []Option) option {
	var o option
	for _, opt := range opts {
		opt(&o)
	}

	o.setDefault()

	return o
}

// WithProcessCount is an option to divide the data into several parts and process them concurrently.
//   - Sets the maximum number of processes that will be run concurrently.
//   - Default is 1, which means the data will be processed sequentially.
//   - If ProcessCount is less than or equal to 0, it will be set to 1.
func WithProcessCount(processCount int) Option {
	return func(o *option) {
		o.ProcessCount = processCount
	}
}

// WithMaxSize is an option to limit the maximum size of the bucket.
//   - Bucket size automatically calculated by len(data) / ProcessCount.
//   - Default no limit which is 0 or less than 0.
func WithMaxSize(maxSize int) Option {
	return func(o *option) {
		o.MaxSize = maxSize
	}
}

// WithMinSize is an option to set the minimum size of the bucket.
//   - If the bucket size is less than MinSize, it will be set to MinSize.
func WithMinSize(minSize int) Option {
	return func(o *option) {
		o.MinSize = minSize
	}
}
