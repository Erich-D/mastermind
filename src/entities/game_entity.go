package entities

type Game struct {
	ID     string `json:"id"`
	Answer string `json:"answer"`
	Status bool   `json:"status"`
}
