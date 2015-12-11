package atomas

import (
	"testing"
	"github.com/assertgo/assert"
	"net/http/httptest"
)

func TestShouldReturnNewGameWithItsUUID(t *testing.T) {
	assert := assert.New(t)
	recorder := httptest.NewRecorder()
	CreateCreateGameHandler(map[string]GameDTO{}, nonRandomUUID, nonRandomElement)(recorder, nil)
	assert.That(recorder.Body.String()).IsEqualTo(ToJsonString(NewGame("uuid", nonRandomElement)))
}

func nonRandomUUID() string {
	return "uuid"
}

func nonRandomElement(_ int) int {
	return 1
}

func TestShouldUseProvidedElementGenerator(t *testing.T) {
	assert := assert.New(t)
	game := NewGame("uuid", elementsFromSlice([]int{1, 2, 3, 4, 5, 6, 7}))
	assert.That(game.Board).IsEqualTo([]int{1, 2, 3, 4, 5, 6})
	assert.That(game.Next).IsEqualTo(7)
	assert.That(game.Score).IsEqualTo(0)
	assert.That(game.Round).IsEqualTo(0)
}

func elementsFromSlice(slice []int) func(int) int {
	i := -1
	return func(_ int) int {
		i++
		return slice[i]
	}
}

func TestShouldSaveToMap(t *testing.T) {
	assert := assert.New(t)
	recorder := httptest.NewRecorder()
	status := map[string]GameDTO{}
	CreateCreateGameHandler(status, nonRandomUUID, nonRandomElement)(recorder, nil)
	assert.That(status["uuid"]).IsEqualTo(NewGame("uuid", nonRandomElement))
}