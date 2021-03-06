package atomas
import (
	"container/list"
)

const (
	PLUS_SIGN  int = 0
)

func EvaluateBoard(arrayBoard []int) (int, []int) {
	score, multiplier, board := lookForPossibleCombinations(toList(arrayBoard), 0, 0)
	return score * multiplier, toArray(board)
}

func lookForPossibleCombinations(board *list.List, multiplier int, score int) (int, int, *list.List) {
	mergablePlusSign := findMergablePlusSign(board)
	if (mergablePlusSign == nil) {
		return score, multiplier, board
	}else {
		return applyPlusSign(board, mergablePlusSign, multiplier + 1, score)
	}
}

func applyPlusSign(board *list.List, element *list.Element, multiplier int, score int) (int, int, *list.List) {
	score, board = mergeElementsAround(board, element, score)
	if (shouldMergeElements(board, element)) {
		return applyPlusSign(board, element, multiplier + 1, score)
	}else {
		return lookForPossibleCombinations(board, multiplier, score)
	}
}

func mergeElementsAround(board *list.List, element *list.Element, score int) (int, *list.List) {
	next := nextWithLoop(board, element)
	prev := prevWithLoop(board, element)
	surroundingValue := next.Value.(int)
	element.Value = Max(surroundingValue, element.Value.(int)) + 1
	board.Remove(prev)
	board.Remove(next)
	return score + surroundingValue * 2, board
}

func nextWithLoop(board *list.List, element *list.Element) *list.Element {
	if (element.Next() != nil ) {
		return element.Next()
	}else {
		return board.Front()
	}
}

func findMergablePlusSign(board *list.List) *list.Element {
	for e := board.Front(); e != nil; e = e.Next() {
		if isMergablePlusSign(board, e) {
			return e
		}
	}
	return nil
}

func isMergablePlusSign(board *list.List, e *list.Element) bool {
	return e.Value == PLUS_SIGN && shouldMergeElements(board, e)
}

func shouldMergeElements(board *list.List, element *list.Element) bool {
	if (board.Len() < 3) {
		return false
	}
	next := nextWithLoop(board, element).Value
	prev := prevWithLoop(board, element).Value
	return areSameAndValid(next, prev)
}

func areSameAndValid(next interface{}, prev interface{}) bool {
	return next == prev && next.(int) > 0
}

func prevWithLoop(board *list.List, element *list.Element) *list.Element {
	if (element.Prev() != nil ) {
		return element.Prev()
	}else {
		return board.Back()
	}
}

func toList(board []int) *list.List {
	result := list.New()
	for _, element := range board {
		result.PushBack(int(element))
	}
	return result
}

func toArray(board *list.List) []int {
	array := make([]int, board.Len())
	i := 0
	for e := board.Front(); e != nil; e = e.Next() {
		array[i] = e.Value.(int)
		i += 1
	}
	return array
}
