package bucket

func bucketSizer(s option) func(totalItem int) int {
	if s.MinSize <= 0 {
		s.MinSize = 1
	}

	if s.ProcessCount <= 0 {
		s.ProcessCount = 1
	}

	return func(totalItem int) int {
		bucketSize := totalItem / s.ProcessCount

		if totalItem%s.ProcessCount != 0 {
			bucketSize++
		}

		if bucketSize < s.MinSize {
			bucketSize = s.MinSize
		} else if s.MaxSize > 0 && bucketSize > s.MaxSize {
			bucketSize = s.MaxSize
		}

		return bucketSize
	}
}
