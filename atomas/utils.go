package atomas

func Insert(slice []int, element int, index int) []int {
	return append(slice[:index], append([]int{element}, slice[index:]...)...)
}