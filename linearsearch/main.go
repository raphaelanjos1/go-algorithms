package linearsearch

func main(list []int, item int) bool {
	for index := 0; index < len(list); index++ {
		if list[index] == item {
			return true
		}
	}
	return false
}
