package atomas
import (
	"net/http"
	"strings"
	"fmt"
	"strconv"
	"errors"
)

const END_OF_GAME = -1000

func CreateMoveHandler(games map[string]GameDTO, randomElement func(int) int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handleMove(w, r, games, randomElement)
	}
}

func handleMove(w http.ResponseWriter, r *http.Request, games map[string]GameDTO, randomElement func(int) int) {
	game, err := findGame(r, games)
	if (err) {
		http.Error(w, err.Error(), 404)
	} else {
		moveTo, err := extractMoveTo(r)
		if (err != nil) {
			http.Error(w, err.Error(), 502)
		} else {
			if (game.Next == END_OF_GAME) {
				http.Error(w, "Game over", 502)
			} else {
				afterMove := Move(game, moveTo, randomElement(game.Round))
				games[game.Id] = afterMove
				fmt.Fprint(w, ToJsonString(afterMove))
			}
		}
	}
}

func findGame(r *http.Request, games map[string]GameDTO) (*GameDTO, error) {
	path := strings.Split(r.URL.Path, "/")
	gameId := path[len(path) - 2]
	game := games[gameId]
	if (game.Id != gameId) {
		return nil, errors.New("Game does not exist")
	}else {
		return &game, nil
	}
}

func extractMoveTo(r *http.Request) (i int, err error) {
	path := strings.Split(r.URL.Path, "/")
	moveTo := path[len(path) - 1]
	return strconv.Atoi(moveTo)
}

func Move(game GameDTO, moveTo int, next int) GameDTO {
	newBoard := append(game.Board[:moveTo], append([]int{game.Next}, game.Board[moveTo:]...)...)
	scoreForMove, newBoard := EvaluateBoard(newBoard)
	if (len(newBoard) > 18) {
		next = END_OF_GAME
	}
	return GameDTO{
		Id:game.Id,
		Board:newBoard,
		Next:next,
		Round:game.Round + 1,
		Score:game.Score + scoreForMove,
	}
}