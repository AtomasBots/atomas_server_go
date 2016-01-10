package atomas
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestShouldReturn0After4Non0Elements(t *testing.T) {
	assert := assert.New(t)
	assert.That(CreateElementGeneratorArray(just1)([]int{1, 1, 1, 1})).IsEqualTo(0)
}

func TestShouldReturnRandomElementWhenLessThan4Elements(t *testing.T) {
	assert := assert.New(t)
	assert.That(CreateElementGeneratorArray(just1)([]int{1, 1, 1})).IsEqualTo(1)
}

func TestShouldReturnRandomElementWhenLessThan4ElementsAreBefore0(t *testing.T) {
	assert := assert.New(t)
	assert.That(CreateElementGeneratorArray(just1)([]int{1, 1, 1, 0})).IsEqualTo(1)
}

func just1() int {
	return 1
}