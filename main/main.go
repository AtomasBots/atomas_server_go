package main
import (
	"net/http"
	"log"
	"github.com/AtomasBots/atomas_server_go/atomas"
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

var version string

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	games := map[string]atomas.GameDTO{}
	elementGenerator := atomas.CreateElementGenerator(rand.Int)
	http.HandleFunc("/new_game", atomas.CreateCreateGameHandler(games, nextUUID, elementGenerator))
	http.HandleFunc("/game/", atomas.CreateGetGameHandler(games))
	http.HandleFunc("/move/", atomas.CreateMoveHandler(games, elementGenerator))
	http.HandleFunc("/highscore", atomas.CreateHighScoreHandler(games))
	http.HandleFunc("/version", atomas.CreateVersionHandler(version))
	port := os.Args[1]
	log.Fatal(http.ListenAndServe(port, nil))
}