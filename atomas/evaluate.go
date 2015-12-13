package atomas

func EvaluateBoard(board []int) (int, []int) {
	for index, element := range board {
		if element == 0 {
			if (isSurroundingSame(board, index)) {
				board[index] = board[(index + 1) % len(board)] + 1
				board = Remove(board, index - 1, index + 1)
			}
		}
	}
	return 0, board
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
