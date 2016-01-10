package atomas

func CreateElementGenerator(rand func() int) func(int) int {
	return func(round int) int {
		if (round == 0) {
			return rand() % 3 + 1
		}else {
			return rand() % 4
		}
	}
}
func CreateElementGeneratorArray(rand func() int) func([]int) int {
	return func(previous []int) int {
		if (indexOfPlus(previous) >= 4) {
			return 0
		}
		return rand() % 4
	}
}

func indexOfPlus(previous []int) int {
	index := 0
	for index < len(previous) {
		if (previous[index] == 0) {
			return index
		}
		index++
	}
	return index
}