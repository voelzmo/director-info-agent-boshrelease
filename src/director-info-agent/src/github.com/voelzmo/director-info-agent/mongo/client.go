package mongo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/voelzmo/bosh-director-client/api"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DeploymentInfo struct {
	DirectorUUID string
	Deployments  []api.Deployment
}

func PushData(address string, user string, password string, dbname string, deploymentInfo *DeploymentInfo) {
	session, err := mgo.Dial(address)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(dbname).C("deployments")
	err = c.Insert(deploymentInfo)
	if err != nil {
		log.Fatal(err)
	}

	result := DeploymentInfo{}
	err = c.Find(bson.M{"directoruuid": deploymentInfo.DirectorUUID}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	prettyDeployments, _ := json.MarshalIndent(result.Deployments, "", "  ")

	fmt.Println("Deployments: ", prettyDeployments)
}
