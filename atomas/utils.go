package atomas

func Insert(slice []int, element int, index int) []int {
	return append(slice[:index], append([]int{element}, slice[index:]...)...)
}

func Max(a int, b int) int {
	if (a > b) {
		return a
	}else {
		return b
	}
}