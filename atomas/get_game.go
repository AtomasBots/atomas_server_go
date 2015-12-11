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
		fmt.Fprint(w, ToJsonString(games[gameId]))
	}
}