package atomas

import (
	"testing"
	"github.com/assertgo/assert"
	"net/http/httptest"
)

func TestShouldReturnNewGameWithItsUUID(t *testing.T) {
	assert := assert.New(t)
	recorder := httptest.NewRecorder()
	CreateCreateGameHandler(nonRandomUUID)(recorder, nil)
	assert.That(recorder.Body.String()).IsEqualTo(ToJsonString(GameWithIdDTO{Id:"uuid"}))
}

func nonRandomUUID() string {
	return "uuid"
}