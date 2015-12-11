package atomas
import (
	"testing"
	"github.com/assertgo/assert"
	"net/http/httptest"
	"net/http"
	"bytes"
)

func TestShouldReturnGameFromStatus(t *testing.T) {
	assert := assert.New(t)
	game := NewGame("uuid", nonRandomElement)
	recorder := httptest.NewRecorder()
	CreateGetGameHandler(map[string]GameDTO{"uuid": game})(recorder, request())
	assert.That(recorder.Body.String()).IsEqualTo(ToJsonString(game))
}

func request() *http.Request {
	req, _ := http.NewRequest("GET", "/game/uuid", bytes.NewReader([]byte("request")))
	return req
}