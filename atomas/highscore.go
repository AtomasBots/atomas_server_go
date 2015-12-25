package atomas
import (
	"net/http"
	"fmt"
)

func CreateHighScoreHandler(games map[string]GameDTO) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, ToJsonString(GetHighScores(games)))
	}
}

func GetHighScores(games map[string]GameDTO) [5]GameDTO {
	var highScores [5]GameDTO
	for _, value := range games {
		for index, saved := range highScores {
			if (value.Score >= saved.Score) {
				if (index <= 3) {
					highScores[4] = highScores[3]
				}
				if (index <= 2) {
					highScores[3] = highScores[2]
				}
				if (index <= 1) {
					highScores[2] = highScores[1]
				}
				if (index <= 0) {
					highScores[1] = highScores[0]
				}
				highScores[index] = value
				break
			}
		}
	}
	return highScores
}
