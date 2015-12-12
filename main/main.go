package main
import (
	"net/http"
	"log"
	"github.com/OrdonTeam/atomas/atomas"
	"github.com/nu7hatch/gouuid"
	"os"
)

func nextUUID() string {
	uuid, err := uuid.NewV4()
	if (err != nil) {
		log.Fatal(err)
	}
	return uuid.String()
}

func elementGenerator(_ int) int {
	return 4
}

func main() {
	games := map[string]atomas.GameDTO{}
	http.HandleFunc("/new_game", atomas.CreateCreateGameHandler(games, nextUUID, elementGenerator))
	http.HandleFunc("/game/", atomas.CreateGetGameHandler(games))
	http.HandleFunc("/move/", atomas.CreateMoveHandler(games, elementGenerator))
	port := os.Args[1]
	log.Fatal(http.ListenAndServe(port, nil))
}