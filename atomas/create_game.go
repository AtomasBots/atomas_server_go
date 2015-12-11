package atomas

import (
	"net/http"
	"fmt"
)

func CreateCreateGameHandler(nextUUID func() string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, ToJsonString(GameWithIdDTO{Id:nextUUID()}))
	}
}
