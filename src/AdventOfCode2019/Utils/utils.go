package utils

// FindInSlice finds an integer element in a slice passed by
func FindInSlice(s []int, element int) bool {
	for _, n := range s {
		if element == n {
			return true
		}
	}

	return false
}
