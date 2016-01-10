package atomas

import (
	"net/http"
	"fmt"
)

func CreateCreateGameHandler(games map[string]GameDTO, nextUUID func() string, randomElement func(int) int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := nextUUID()
		newGame := NewGame(uuid, r.RemoteAddr, r.FormValue("name"), randomElement)
		games[uuid] = newGame
		fmt.Fprint(w, ToJsonString(newGame))
	}
}