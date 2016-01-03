package atomas
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestInsertShouldInsertAtBegin(t *testing.T) {
	assert := assert.New(t)
	assert.That(Insert([]int{2, 3}, 1, 0)).IsEqualTo([]int{1, 2, 3})
}

func TestInsertShouldInsertInMiddle(t *testing.T) {
	assert := assert.New(t)
	assert.That(Insert([]int{2, 3}, 1, 1)).IsEqualTo([]int{2, 1, 3})
}

func TestInsertShouldInsertAtTheEnd(t *testing.T) {
	assert := assert.New(t)
	assert.That(Insert([]int{2, 3}, 1, 2)).IsEqualTo([]int{2, 3, 1})
}

func TestMaxShouldReturnFirst(t *testing.T) {
	assert := assert.New(t)
	assert.ThatInt(Max(2, 1)).IsEqualTo(2)
}

func TestMaxShouldReturnSecond(t *testing.T) {
	assert := assert.New(t)
	assert.ThatInt(Max(1, 2)).IsEqualTo(2)
}