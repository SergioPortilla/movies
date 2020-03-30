package commands

type MovieCommand struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Duration int64  `json:"duration"`
	Synopsis string `json:"synopsis"`
}
