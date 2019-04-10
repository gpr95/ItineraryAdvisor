package main

import (
	"github.com/tkanos/gonfig"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func getSecrets() (Configuration, error) {
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.secret.json", &configuration)
	return configuration, err
}
