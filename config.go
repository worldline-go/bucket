package bucket

type Config struct {
	// MinSize is the minimum size of the bucket. Default is 1.
	//   - If the bucket size is less than MinSize, it will be set to MinSize.
	MinSize int `cfg:"min_size" json:"min_size"`
	// MaxSize is the maximum size of the bucket.
	MaxSize int `cfg:"max_size" json:"max_size"`
	// ProcessCount is the number of processes that will be run concurrently.
	ProcessCount int `cfg:"process_count" json:"process_count"`
}

// ToOption will convert Config to Option.
func (c Config) ToOption() Option {
	return func(o *option) {
		o.MinSize = c.MinSize
		o.MaxSize = c.MaxSize
		o.ProcessCount = c.ProcessCount
	}
}
