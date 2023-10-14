package binarysearch

func main(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for low < high {
		middle := low + (high-low)/2
		value := haystack[middle]

		if value == needle {
			return true
		} else if value > needle {
			high = middle
		} else {
			low = middle + 1
		}
	}

	return false
}
