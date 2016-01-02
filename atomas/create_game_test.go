package atomas

import (
	"testing"
	"github.com/assertgo/assert"
	"net/http/httptest"
	"net/http"
	"bytes"
)

func TestShouldReturnNewGameWithItsUUID(t *testing.T) {
	assert := assert.New(t)
	recorder := httptest.NewRecorder()
	CreateCreateGameHandler(map[string]GameDTO{}, nonRandomUUID, nonRandomElement)(recorder, requestNewGame())
	assert.That(recorder.Body.String()).IsEqualTo(ToJsonString(NewGame("uuid", "", "", nonRandomElement)))
}

func nonRandomUUID() string {
	return "uuid"
}

func nonRandomElement(_ int) int {
	return 1
}

func TestShouldUseProvidedElementGenerator(t *testing.T) {
	assert := assert.New(t)
	game := NewGame("uuid", "", "", elementsFromSlice([]int{1, 2, 3, 4, 5, 6, 7}))
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
	CreateCreateGameHandler(status, nonRandomUUID, nonRandomElement)(recorder, requestNewGame())
	assert.That(status["uuid"]).IsEqualTo(NewGame("uuid", "", "", nonRandomElement))
}

func TestShouldSaveToMapWithName(t *testing.T) {
	assert := assert.New(t)
	recorder := httptest.NewRecorder()
	status := map[string]GameDTO{}
	CreateCreateGameHandler(status, nonRandomUUID, nonRandomElement)(recorder, requestNewGameWitName())
	assert.That(status["uuid"]).IsEqualTo(NewGame("uuid", "", "name", nonRandomElement))
}

func requestNewGame() *http.Request {
	req, _ := http.NewRequest("GET", "/new_game", bytes.NewReader([]byte("request")))
	return req
}

func requestNewGameWitName() *http.Request {
	req, _ := http.NewRequest("GET", "/new_game?name=name", bytes.NewReader([]byte("request")))
	return req
}