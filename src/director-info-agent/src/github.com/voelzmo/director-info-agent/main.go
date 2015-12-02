package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/voelzmo/bosh-director-client/director"
	"github.com/voelzmo/director-info-agent/mongo"
)

type Arguments struct {
	Target          string
	RootCAPath      string
	ClientName      string
	ClientSecret    string
	MongoDBAddress  string
	MongoDBUser     string
	MongoDBPassword string
	MongoDBName     string
}

func main() {
	args, err := parseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	director := director.NewDirector(args.Target, args.RootCAPath, args.ClientName, args.ClientSecret)

	ticker := time.NewTicker(30 * time.Second)
	quitChannel := make(chan struct{})
	defer close(quitChannel)

	for {
		fmt.Println("Waiting for a channel signal...")
		select {
		case <-ticker.C:
			fmt.Println("received ticker signal")
			deployments := director.Deployments()
			mongo.PushData(args.MongoDBAddress, args.MongoDBUser, args.MongoDBPassword, args.MongoDBName, &mongo.DeploymentInfo{director.Status().UUID, deployments})
			fmt.Printf("Here is the current list of deployments: '%s'", deployments)
		case <-quitChannel:
			fmt.Println("received quit signal")
			ticker.Stop()
			return
		}
	}

}

func parseArgs(args []string) (Arguments, error) {
	expectedNumberOfArgs := 9
	if len(args) != expectedNumberOfArgs {
		return Arguments{}, fmt.Errorf("parseArgs: Wrong number of arguments, expected %v, but got %v\nUsage: bosh-director-client <director URL> <root CA path> <oauth client name> <oauth client secret>", expectedNumberOfArgs-1, len(args)-1)
	}
	return Arguments{args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8]}, nil
}
