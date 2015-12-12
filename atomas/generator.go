package atomas

func CreateElementGenerator(rand func() int) func(int) int {
	return func(round int) int {
		return rand() % 3 + 1
	}
}