package api

type UserAuth struct {
	Type    string            `json:"type"`
	Options map[string]string `json:"options"`
}

type Feature struct {
	Status bool              `json:"status"`
	Extras map[string]string `json:"extras"`
}

type Status struct {
	Name               string             `json:"name"`
	UUID               string             `json:"uuid"`
	Version            string             `json:"version"`
	User               string             `json:"user"`
	CPI                string             `json:"cpi"`
	UserAuthentication UserAuth           `json:"user_authentication"`
	Features           map[string]Feature `json:"features"`
}
