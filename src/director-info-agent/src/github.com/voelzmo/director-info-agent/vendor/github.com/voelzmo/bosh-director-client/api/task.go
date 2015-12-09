package api

type Task struct {
	ID          int    `json:"id"`
	State       string `json:"state"`
	Description string `json:"description"`
	Timestamp   int    `json:"timestamp"`
	Result      string `json:"result"`
	User        string `json:"user"`
}
