package atomas

const (
	PLUS_SIGN  int = iota
)

func EvaluateBoard(board []int) (int, []int) {
	score := 0
	for index, element := range board {
		if element == PLUS_SIGN {
			if (isSurroundingSame(board, index)) {
				score += board[(index + 1) % len(board)] * 2
				board[index] = board[(index + 1) % len(board)] + 1
				board = Remove(board, index - 1, index + 1)
			}
		}
	}
	return score, board
}

func Remove(board []int, indexes ... int) []int {
	if (len(indexes) == 0 ) {
		return board
	}else {
		indexToRemove := (indexes[0] % len(board) + len(board)) % len(board)
		return Remove(append(board[:indexToRemove], board[indexToRemove + 1:]...), decreaseIndexes(indexes[1:], indexToRemove)...)
	}
}

func decreaseIndexes(indexes []int, threshold int) []int {
	for index, value := range indexes {
		if (indexes[index] > threshold) {
			indexes[index] = value - 1
		}
	}
	return indexes
}

func isSurroundingSame(board []int, index int) bool {
	return board[((index - 1) % len(board) + len(board)) % len(board)] == board[(index + 1) % len(board)]
}
