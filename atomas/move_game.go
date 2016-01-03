package atomas
import (
	"net/http"
	"strings"
	"fmt"
	"strconv"
)

const END_OF_GAME = -1000

func CreateMoveHandler(games map[string]GameDTO, randomElement func(int) int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := handleMove(r, games, randomElement)
		w.WriteHeader(response.Code)
		fmt.Fprint(w, response.Response)
	}
}

func handleMove(r *http.Request, games map[string]GameDTO, randomElement func(int) int) *ServerResponse {
	game, err := findGame(r, games)
	if (err != nil) {
		return err
	}
	moveTo, err := extractMoveTo(r, *game)
	if (err != nil) {
		return err
	}
	if (game.Next == END_OF_GAME) {
		return &ServerResponse{"Game over", http.StatusBadRequest }
	} else {
		afterMove := Move(*game, moveTo, randomElement(game.Round))
		games[game.Id] = afterMove
		return &ServerResponse{ToJsonString(afterMove), http.StatusOK }
	}
}

func findGame(r *http.Request, games map[string]GameDTO) (*GameDTO, *ServerResponse) {
	path := strings.Split(r.URL.Path, "/")
	gameId := path[len(path) - 2]
	game := games[gameId]
	if (game.Id != gameId) {
		return nil, &ServerResponse{"Game does not exist", http.StatusNotFound }
	}else {
		return &game, nil
	}
}

func extractMoveTo(r *http.Request, game GameDTO) (int, *ServerResponse) {
	path := strings.Split(r.URL.Path, "/")
	moveTo := path[len(path) - 1]
	i, err := strconv.Atoi(moveTo)
	if (err != nil) {
		return -1, &ServerResponse{fmt.Sprintf("\"%s\" is not an integer", moveTo), http.StatusBadRequest }
	} else if (i < 0 || i > len(game.Board)) {
		return -1, &ServerResponse{"Index out of bounds", http.StatusBadRequest }
	}else {
		return i, nil
	}
}

func Move(game GameDTO, moveTo int, next int) GameDTO {
	newBoard := Insert(game.Board, game.Next, moveTo)
	scoreForMove, newBoard := EvaluateBoard(newBoard)
	if (len(newBoard) > 18) {
		next = END_OF_GAME
	}
	return GameDTO{
		Id:game.Id,
		Ip:game.Ip,
		Board:newBoard,
		Next:next,
		Round:game.Round + 1,
		Score:game.Score + scoreForMove,
	}
}