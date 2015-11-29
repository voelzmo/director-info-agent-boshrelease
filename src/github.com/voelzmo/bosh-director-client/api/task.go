package api

type Task struct {
	ID          string `json:"id"`
	State       string `json:"state"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
	Result      string `json:"result"`
	User        string `json:"user"`
}
