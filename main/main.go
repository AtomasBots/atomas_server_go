package main
import (
	"net/http"
	"log"
	"github.com/OrdonTeam/atomas/atomas"
	"github.com/nu7hatch/gouuid"
	"os"
	"math/rand"
	"time"
)

func nextUUID() string {
	uuid, err := uuid.NewV4()
	if (err != nil) {
		log.Fatal(err)
	}
	return uuid.String()
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	games := map[string]atomas.GameDTO{}
	elementGenerator := atomas.CreateElementGenerator(rand.Int)
	http.HandleFunc("/new_game", atomas.CreateCreateGameHandler(games, nextUUID, elementGenerator))
	http.HandleFunc("/game/", atomas.CreateGetGameHandler(games))
	http.HandleFunc("/move/", atomas.CreateMoveHandler(games, elementGenerator))
	port := os.Args[1]
	log.Fatal(http.ListenAndServe(port, nil))
}