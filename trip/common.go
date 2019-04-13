package trip

import (
	"log"

	"github.com/tkanos/gonfig"
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
