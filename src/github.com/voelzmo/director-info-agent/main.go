package main

import (
	"fmt"
	"log"
	"os"

	"github.com/voelzmo/bosh-director-client/director"
)

type Arguments struct {
	Target       string
	RootCAPath   string
	ClientName   string
	ClientSecret string
}

func main() {
	args, err := parseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	director := director.NewDirector(args.Target, args.RootCAPath, args.ClientName, args.ClientSecret)
	director.Deployments()

}

func parseArgs(args []string) (Arguments, error) {
	expectedNumberOfArgs := 5
	if len(args) != expectedNumberOfArgs {
		return Arguments{}, fmt.Errorf("parseArgs: Wrong number of arguments, expected %v, but got %v\nUsage: bosh-director-client <director URL> <root CA path> <oauth client name> <oauth client secret>", expectedNumberOfArgs-1, len(args)-1)
	}
	return Arguments{args[1], args[2], args[3], args[4]}, nil
}
