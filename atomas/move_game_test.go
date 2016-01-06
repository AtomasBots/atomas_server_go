package atomas
import (
	"testing"
	"github.com/assertgo/assert"
	"net/http/httptest"
	"net/http"
	"bytes"
	"fmt"
)

func TestShouldChangeGame(t *testing.T) {
	assert := assert.New(t)
	games := initialGames()
	CreateMoveHandler(games, nonRandomElement)(httptest.NewRecorder(), moveRequest("uuid", "0"))
	assert.That(games["uuid"].Board).IsEqualTo([]int{5, 1, 2, 3, 4})
}

func TestMoveShouldReturnNotFound(t *testing.T) {
	assert := assert.New(t)
	games := initialGames()
	recorder := httptest.NewRecorder()
	CreateMoveHandler(games, nonRandomElement)(recorder, moveRequest("incorrect uuid", "0"))
	assert.That(recorder.Code).IsEqualTo(404)
	assert.That(recorder.Body.String()).IsEqualTo("Game does not exist")
}

func TestMoveShouldReturnParseException(t *testing.T) {
	assert := assert.New(t)
	games := initialGames()
	recorder := httptest.NewRecorder()
	CreateMoveHandler(games, nonRandomElement)(recorder, moveRequest("uuid", "NAN"))
	assert.That(recorder.Code).IsEqualTo(400)
	assert.That(recorder.Body.String()).IsEqualTo("\"NAN\" is not an integer")
}

func TestMoveShouldReturnOutOfBoundsException(t *testing.T) {
	assert := assert.New(t)
	games := initialGames()
	recorder := httptest.NewRecorder()
	CreateMoveHandler(games, nonRandomElement)(recorder, moveRequest("uuid", "100"))
	assert.That(recorder.Code).IsEqualTo(400)
	assert.That(recorder.Body.String()).IsEqualTo("Index out of bounds")
}

func TestMoveShouldReturnGameOverException(t *testing.T) {
	assert := assert.New(t)
	games := initialGames()
	recorder := httptest.NewRecorder()
	CreateMoveHandler(games, nonRandomElement)(recorder, moveRequest("eog", "0"))
	assert.That(recorder.Code).IsEqualTo(400)
	assert.That(recorder.Body.String()).IsEqualTo("Game over")
}

func TestMoveShouldKeepName(t *testing.T) {
	assert := assert.New(t)
	games := initialGames()
	CreateMoveHandler(games, nonRandomElement)(httptest.NewRecorder(), moveRequest("uuid", "0"))
	assert.That(games["uuid"].Name).IsEqualTo("name")
}

func initialGames() map[string]GameDTO {
	return map[string]GameDTO{
		"uuid" : GameDTO{
			Id:"uuid",
			Name:"name",
			Board:[]int{1, 2, 3, 4},
			Next: 5,
			Round:0,
			Score:0,
		},
		"eog" : GameDTO{
			Id:"eog",
			Board:[]int{1, 2, 3, 4},
			Next: END_OF_GAME,
			Round:0,
			Score:0,
		},
	}
}

func moveRequest(id string, target string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/move/%s/%s", id, target), bytes.NewReader([]byte("request")))
	return req
}