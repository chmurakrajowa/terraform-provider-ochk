package ochk

func transformToStringSlice(size int, mapFunc func(idx int) string) []string {
	resultSlice := make([]string, size)
	for i := 0; i < size; i++ {
		resultSlice[i] = mapFunc(i)
	}

	return resultSlice
}

func findIndexOfFirstMatch(size int, matchFunc func(idx int) bool) int {
	for i := 0; i < size; i++ {
		if matchFunc(i) {
			return i
		}
	}

	return -1
}

func slicesEqual(lhsSize int, rhsSize int, matchFunc func(lhsIdx int, rhsIdx int) bool) bool {
	if lhsSize != rhsSize {
		return false
	}

	for i := 0; i < lhsSize; i++ {
		if !matchFunc(i, i) {
			return false
		}
	}

	return true
}
