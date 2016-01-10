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
		indexOfPlus := 0
		for indexOfPlus < len(previous) {
			if (previous[indexOfPlus] == 0) {
				break
			}
			indexOfPlus++
		}
		if (indexOfPlus >= 4) {
			return 0
		}
		return rand() % 4
	}
}