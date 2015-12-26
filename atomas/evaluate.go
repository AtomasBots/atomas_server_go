package atomas
import (
	"container/list"
)

const (
	PLUS_SIGN  int = iota
)

func EvaluateBoard(arrayBoard []int) (int, []int) {
	score := 0
	board := toList(arrayBoard)
	for e := board.Front(); e != nil; e = e.Next() {
		if e.Value == PLUS_SIGN {
			if (isSurroundingSame(board, e)) {
				score += nextWithLoop(board, e).Value.(int) * 2
				e.Value = nextWithLoop(board, e).Value.(int) + 1
				board = removeNeighbours(board, e)
			}
		}
	}
	return score, toArray(board)
}
func removeNeighbours(board *list.List, element *list.Element) *list.List {
	newBoard := list.New()
	for e := board.Front(); e != nil; e = e.Next() {
		if (e != prevWithLoop(board, element) && e != nextWithLoop(board, element)) {
			newBoard.PushBack(e.Value.(int))
		}
	}
	return newBoard
}

func nextWithLoop(board *list.List, element *list.Element) *list.Element {
	if (element.Next() != nil ) {
		return element.Next()
	}else {
		return board.Front()
	}
}

func prevWithLoop(board *list.List, element *list.Element) *list.Element {
	if (element.Prev() != nil ) {
		return element.Prev()
	}else {
		return board.Back()
	}
}

func isSurroundingSame(board *list.List, element *list.Element) bool {
	return nextWithLoop(board, element).Value == prevWithLoop(board, element).Value
}

func toList(board []int) *list.List {
	result := list.New()
	for _, element := range board {
		result.PushBack(int(element))
	}
	return result
}

func toArray(board *list.List) []int {
	array := make([]int, size(board))
	i := 0
	for e := board.Front(); e != nil; e = e.Next() {
		array[i] = e.Value.(int)
		i += 1
	}
	return array
}

func size(list *list.List) int {
	count := 0
	for e := list.Front(); e != nil; e = e.Next() {
		count += 1
	}
	return count
}