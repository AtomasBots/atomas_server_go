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
	score, board := EvaluateBoard([]int{1, PLUS_SIGN, 1})
	assert.That(score).IsEqualTo(2)
	assert.That(board).IsEqualTo([]int{2})
}

func TestShouldSumElementsCorrectly(t *testing.T) {
	assert := assert.New(t)
	score, board := EvaluateBoard([]int{2, PLUS_SIGN, 2})
	assert.That(score).IsEqualTo(4)
	assert.That(board).IsEqualTo([]int{3})
}

func TestShouldSumElementsCorrectlyRecursively(t *testing.T) {
	assert := assert.New(t)
	score, board := EvaluateBoard([]int{1, 2, PLUS_SIGN, 2, 1 })
	assert.That(score).IsEqualTo(12)
	assert.That(board).IsEqualTo([]int{4})
}

func TestShouldLeaveElementsNotChangedInSumProcess(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{3, 2, 0, 2, 1})
	assert.That(board).IsEqualTo([]int{3, 3, 1})
}

func TestShouldNotChangeBoardWhenPlusIsSurroundedByDifferentElements(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{1, PLUS_SIGN, 2})
	assert.That(board).IsEqualTo([]int{1, PLUS_SIGN, 2})
}

func TestShouldSeeBoardAsLoop(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{1, 1, PLUS_SIGN})
	assert.That(board).IsEqualTo([]int{2})
}

func TestShouldSeeBoardAsLoop2(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{PLUS_SIGN, 1, 1})
	assert.That(board).IsEqualTo([]int{2})
}

func TestShouldSeeBoardAsLoopButCountElementsOnlyOnce(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{1, PLUS_SIGN})
	assert.That(board).IsEqualTo([]int{1, PLUS_SIGN})
}

func TestShouldMergeElementsButLeftNonMatching(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{1, 2, 3, 4, PLUS_SIGN, 4, 3, 2})
	assert.That(board).IsEqualTo([]int{1, 7})
}

func TestShouldNotMergePlusElements(t *testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{4, PLUS_SIGN, PLUS_SIGN, PLUS_SIGN, 4})
	assert.That(board).IsEqualTo([]int{4, PLUS_SIGN, PLUS_SIGN, PLUS_SIGN, 4})
}

func TestShouldMergeElementsRecursivelyUsingOtherPluses(t*testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{PLUS_SIGN, 1, PLUS_SIGN, 1, 4, 3, 2, 2})
	assert.That(board).IsEqualTo([]int{3, 4, 3, 2})
}

func TestShouldMergeFirstOfMultiPossibleMerges(t*testing.T) {
	assert := assert.New(t)
	_, board := EvaluateBoard([]int{3, PLUS_SIGN, 3, PLUS_SIGN, 3, 5, 4})
	assert.That(board).IsEqualTo([]int{4, PLUS_SIGN, 3, 5, 4})
}
