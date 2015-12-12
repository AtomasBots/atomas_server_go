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