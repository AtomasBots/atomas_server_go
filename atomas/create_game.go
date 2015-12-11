package atomas

import (
	"net/http"
	"fmt"
)

func CreateCreateGameHandler(games map[string]GameDTO, nextUUID func() string, randomElement func(int) int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := nextUUID()
		newGame := NewGame(uuid, randomElement)
		games[uuid] = newGame
		fmt.Fprint(w, ToJsonString(newGame))
	}
}

func NewGame(uuid string, randomElement func(int) int) GameDTO {
	return GameDTO{
		Id: uuid,
		Board: []int{randomElement(0), randomElement(0), randomElement(0), randomElement(0), randomElement(0), randomElement(0)},
		Next: randomElement(0),
		Round: 0,
		Score: 0,
	}
}