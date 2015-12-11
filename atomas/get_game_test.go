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
	recorder := httptest.NewRecorder()
	CreateGetGameHandler()(recorder, request())
	assert.That(recorder.Body.String()).IsEqualTo(ToJsonString(NewGame("uuid", nonRandomElement)))
}

func request() *http.Request {
	req, _ := http.NewRequest("GET", "/game/uuid", bytes.NewReader([]byte("request")))
	return req
}