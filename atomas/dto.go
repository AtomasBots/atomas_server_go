package atomas

type GameDTO struct {
	Id    string `json:"id"`
	Ip    string `json:"ip"`
	Name    string `json:"name,omitempty"`
	Board []int `json:"board"`
	Next  int `json:"next"`
	Round int `json:"round"`
	Score int `json:"score"`
	InitialBoard []int `json:"initial_borad"`
	PreviousElements []int `json:"previous_elements"`
}

func NewGame(uuid string, ipAddress string, name string, randomElement func(int) int) GameDTO {
	board := []int{randomElement(0), randomElement(0), randomElement(0), randomElement(0), randomElement(0), randomElement(0)}
	return GameDTO{
		Id: uuid,
		Ip: ipAddress,
		Name: name,
		Board: board,
		Next: randomElement(0),
		Round: 0,
		Score: 0,
		InitialBoard: board,
		PreviousElements: []int{},
	}
}

func (game GameDTO) NextBoard(newBoard []int, newNext int, scoreForMove int) GameDTO {
	return GameDTO{
		Id:game.Id,
		Name: game.Name,
		Ip:game.Ip,
		Board:newBoard,
		Next:newNext,
		Round:game.Round + 1,
		Score:game.Score + scoreForMove,
		InitialBoard:game.InitialBoard,
		PreviousElements:append([]int{game.Next}, game.PreviousElements...),
	}
}

type ServerResponse struct {
	Response string
	Code     int
}