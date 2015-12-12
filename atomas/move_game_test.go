package atomas
import (
	"testing"
	"github.com/assertgo/assert"
	"net/http/httptest"
	"net/http"
	"bytes"
)

func TestShouldChangeGame(t *testing.T) {
	assert := assert.New(t)
	games := initialGames()
	recorder := httptest.NewRecorder()
	CreateMoveHandler(games, nonRandomElement)(recorder, moveRequest())
	assert.That(games["uuid"].Board).IsEqualTo([]int{5, 1, 2, 3, 4})
}

func initialGames() map[string]GameDTO {
	return map[string]GameDTO{
		"uuid" : GameDTO{
			Id:"uuid",
			Board:[]int{1, 2, 3, 4},
			Next: 5,
			Round:0,
			Score:0,
		},
	}
}

func moveRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/move/uuid/0", bytes.NewReader([]byte("request")))
	return req
}