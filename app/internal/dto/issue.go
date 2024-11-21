package dto

type Issue struct {
	ID      string `json:"id"`
	Summary string `json:"summary"`
	Link    string `json:"self"`
}
