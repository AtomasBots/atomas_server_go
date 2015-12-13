package atomas
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestShouldNotChangeBoardWithoutSpecialElements(t *testing.T) {
	assert := assert.New(t)
	score, board := EvaluateBoard([]int{1, 2, 3})
	assert.That(score).IsEqualTo(0)
	assert.That(board).IsEqualTo([]int{1, 2, 3})
}

func TestShouldChangeBoardWhenPlusIsSurroundedBySameElements(t *testing.T) {
	assert := assert.New(t)
	score, board := EvaluateBoard([]int{1, 0, 1})
	assert.That(score).IsEqualTo(2)
	assert.That(board).IsEqualTo([]int{2})
}

func TestShouldSumElementsCorrectly(t *testing.T) {
	assert := assert.New(t)
	score, board := EvaluateBoard([]int{2, 0, 2})
	assert.That(score).IsEqualTo(4)
	assert.That(board).IsEqualTo([]int{3})
}

func TestShouldLeaveElementsNotChangedInSumProcess(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{3, 2, 0, 2, 1})
	assert.That(board).IsEqualTo([]int{3, 3, 1})
}

func TestShouldNotChangeBoardWhenPlusIsSurroundedByDifferentElements(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{1, 0, 2})
	assert.That(board).IsEqualTo([]int{1, 0, 2})
}

func TestShouldSeeBoardAsLoop(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{1, 1, 0})
	assert.That(board).IsEqualTo([]int{2})
}

func TestShouldSeeBoardAsLoop2(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{0, 1, 1})
	assert.That(board).IsEqualTo([]int{2})
}

func TestShouldRemoveRemoveSingleElement(t *testing.T) {
	assert := assert.New(t)
	assert.That(Remove([]int{1, 2, 3, 4}, 0)).IsEqualTo([]int{2, 3, 4})
}
func TestShouldRemoveRemoveTwoElementsElement(t *testing.T) {
	assert := assert.New(t)
	assert.That(Remove([]int{1, 2, 3, 4}, 0, 1)).IsEqualTo([]int{3, 4})
}
func TestShouldRemoveRemoveTwoElementsInReverseOrderElement(t *testing.T) {
	assert := assert.New(t)
	assert.That(Remove([]int{1, 2, 3, 4}, 1, 0)).IsEqualTo([]int{3, 4})
}