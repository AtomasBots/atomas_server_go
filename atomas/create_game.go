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

func NewGame(uuid string, ipAddress string, name string, randomElement func(int) int) GameDTO {
	return GameDTO{
		Id: uuid,
		Ip: ipAddress,
		Name: name,
		Board: []int{randomElement(0), randomElement(0), randomElement(0), randomElement(0), randomElement(0), randomElement(0)},
		Next: randomElement(0),
		Round: 0,
		Score: 0,
	}
}