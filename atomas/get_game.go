package atomas
import (
	"net/http"
	"strings"
	"fmt"
)

func CreateGetGameHandler(games map[string]GameDTO) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		gameId := path[len(path) - 1]
		game := games[gameId]
		if (game.Id == gameId) {
			fmt.Fprint(w, ToJsonString(game))
		}else {
			http.NotFound(w, r)
		}
	}
}