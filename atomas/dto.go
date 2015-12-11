package atomas

type GameDTO struct {

}

type GameWithIdDTO struct {
	Id   string `json:"id"`
	Game GameDTO `json:"game"`
}