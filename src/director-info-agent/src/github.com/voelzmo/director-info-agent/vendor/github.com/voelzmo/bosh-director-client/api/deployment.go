package api

type Deployment struct {
	Name        string     `json:"name"`
	CloudConfig string     `json:"cloud_config"`
	Releases    []Release  `json:"releases"`
	Stemcells   []Stemcell `json:"stemcells"`
}
