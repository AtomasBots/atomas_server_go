package atomas

import (
	"net/http"
	"fmt"
)

func CreateCreateGameHandler(nextUUID func() string, randomElement func(int) int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, ToJsonString(NewGame(nextUUID(), randomElement)))
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