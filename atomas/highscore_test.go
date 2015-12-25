package atomas
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestGetHighScoresShouldReturnInCorrectOrder(t *testing.T) {
	assert := assert.New(t)
	highScores := GetHighScores(mapOf(game1(), game2(), game3(), game4(), game5()))
	assert.That(highScores[0]).IsEqualTo(game5())
	assert.That(highScores[1]).IsEqualTo(game4())
	assert.That(highScores[2]).IsEqualTo(game3())
	assert.That(highScores[3]).IsEqualTo(game2())
	assert.That(highScores[4]).IsEqualTo(game1())
}

func TestGetHighScoresShouldReturnInCorrectOrder2(t *testing.T) {
	assert := assert.New(t)
	highScores := GetHighScores(mapOf(game1(), game3(), game4(), game2(), game5()))
	assert.That(highScores[0]).IsEqualTo(game5())
	assert.That(highScores[1]).IsEqualTo(game4())
	assert.That(highScores[2]).IsEqualTo(game3())
	assert.That(highScores[3]).IsEqualTo(game2())
	assert.That(highScores[4]).IsEqualTo(game1())
}

func mapOf(games ... GameDTO) map[string]GameDTO {
	gamesMap := map[string]GameDTO{}
	for _, game := range games {
		gamesMap[game.Id] = game
	}
	return gamesMap
}

func game1() GameDTO {
	return GameDTO{Id:"1", Score:1}
}
func game2() GameDTO {
	return GameDTO{Id:"2", Score:2}
}
func game3() GameDTO {
	return GameDTO{Id:"3", Score:3}
}
func game4() GameDTO {
	return GameDTO{Id:"4", Score:4}
}
func game5() GameDTO {
	return GameDTO{Id:"5", Score:5}
}
func game6() GameDTO {
	return GameDTO{Id:"6", Score:6}
}