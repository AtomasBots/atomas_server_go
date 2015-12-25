package atomas

type GameDTO struct {
	Id    string `json:"id"`
	Ip    string `json:"ip"`
	Board []int `json:"board"`
	Next  int `json:"next"`
	Round int `json:"round"`
	Score int `json:"score"`
}