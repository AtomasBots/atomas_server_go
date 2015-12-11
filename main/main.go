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

func main() {
	http.HandleFunc("/new_game", atomas.CreateCreateGameHandler(nextUUID))
	port := os.Args[1]
	log.Fatal(http.ListenAndServe(port, nil))
}