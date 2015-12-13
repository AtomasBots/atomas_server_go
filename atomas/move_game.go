package atomas
import (
	"net/http"
	"strings"
	"fmt"
	"strconv"
)

func CreateMoveHandler(games map[string]GameDTO, randomElement func(int) int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		moveTo := path[len(path) - 1]
		gameId := path[len(path) - 2]
		game := games[gameId]
		if (game.Id == gameId) {
			moveToInt, err := strconv.Atoi(moveTo)
			if (err != nil) {
				http.Error(w, "Bad request", 502)
			}else {
				afterMove := Move(game, moveToInt, randomElement(game.Round))
				games[gameId] = afterMove
				fmt.Fprint(w, ToJsonString(afterMove))
			}
		}else {
			http.NotFound(w, r)
		}
	}
}

func Move(game GameDTO, moveTo int, next int) GameDTO {
	newBoard := append(game.Board[:moveTo], append([]int{game.Next}, game.Board[moveTo:]...)...)
	scoreForMove, newBoard := EvaluateBoard(newBoard)
	return GameDTO{
		Id:game.Id,
		Board:newBoard,
		Next:next,
		Round:game.Round + 1,
		Score:game.Score + scoreForMove,
	}
}