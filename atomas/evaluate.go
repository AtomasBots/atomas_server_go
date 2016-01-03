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
	next := nextWithLoop(board, element)
	prev := prevWithLoop(board, element)
	surroundingValue := next.Value.(int)
	score += surroundingValue * 2
	element.Value = Max(surroundingValue, element.Value.(int)) + 1
	board.Remove(prev)
	board.Remove(next)
	if (shouldMergeElements(board, element)) {
		return applyPlusSign(board, element, multiplier + 1, score)
	} else if mergablePlusSign := findMergablePlusSign(board); mergablePlusSign != nil {
		return applyPlusSign(board, mergablePlusSign, multiplier + 1, score)
	}else {
		return score, multiplier, board
	}
}

func shouldMergeElements(board *list.List, element *list.Element) bool {
	return (board.Len() > 2 && isSurroundingSame(board, element) && theyAreProperElement(board, element))
}

func findMergablePlusSign(board *list.List) *list.Element {
	for e := board.Front(); e != nil; e = e.Next() {
		if e.Value == PLUS_SIGN && shouldMergeElements(board, e) {
			return e
		}
	}
	return nil
}

func theyAreProperElement(board *list.List, element *list.Element) bool {
	return nextWithLoop(board, element).Value.(int) > 0
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
	array := make([]int, board.Len())
	i := 0
	for e := board.Front(); e != nil; e = e.Next() {
		array[i] = e.Value.(int)
		i += 1
	}
	return array
}
