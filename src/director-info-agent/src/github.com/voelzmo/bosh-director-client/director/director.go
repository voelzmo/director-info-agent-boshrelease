package director

import (
	"encoding/base64"
	"fmt"

	"github.com/voelzmo/bosh-director-client/api"
)

type Director interface {
	Status() api.Status
	Login() api.Login
	Deployments() []api.Deployment
}

type director struct {
	target       string
	rootCAPath   string
	clientName   string
	clientSecret string
}

func NewDirector(target string, rootCAPath string, clientName string, clientSecret string) Director {
	return &director{target, rootCAPath, clientName, clientSecret}
}

func (d *director) Status() api.Status {
	var status api.Status

	auth := "" // info endpoint doesn't need authorization
	GetClient(d.target, d.rootCAPath, auth).RequestAndParseJSON("GET", "/info", make(map[string]string), nil, &status)

	return status
}

func (d *director) Login() api.Login {
	var login api.Login

	directorStatus := d.Status()
	authURL := directorStatus.UserAuthentication.Options["url"]

	postBody := []byte(`grant_type=client_credentials`)

	userPassword := []byte(fmt.Sprintf("%s:%s", d.clientName, d.clientSecret))
	auth := "Basic " + base64.StdEncoding.EncodeToString(userPassword)
	headers := make(map[string]string)
	headers["content-type"] = "application/x-www-form-urlencoded;charset=utf-8"
	headers["accept"] = "application/json;charset=utf-8"
	GetClient(authURL, d.rootCAPath, auth).RequestAndParseJSON("POST", "/oauth/token", headers, postBody, &login)
	return login
}

func (d *director) Deployments() []api.Deployment {
	var deployments []api.Deployment

	login := d.Login()
	auth := fmt.Sprintf("%s %s", login.TokenType, login.AccessToken)
	GetClient(d.target, d.rootCAPath, auth).RequestAndParseJSON("GET", "/deployments", make(map[string]string), nil, &deployments)

	return deployments
}

//
// func (d *director) Tasks() []api.Task {
// 	var tasks []api.Task
//
// 	directorClient := NewClient(d.rootCAPath)
//
// 	resp, err := directorClient.Get(fmt.Sprintf("%s/tasks", d.target))
// 	if err != nil {
// 		log.Fatal("Error getting director task: %s", err)
// 	}
// 	defer resp.Body.Close()
//
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	json.Unmarshal(body, &tasks)
//
// 	return tasks
// }
