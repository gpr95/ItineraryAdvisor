package trip

import (
	"fmt"
	"googlemaps.github.io/maps"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	GoogleAPIKey string
}

func getGoogleClient() *maps.Client {
	configuration, _ := getSecrets()

	var client *maps.Client
	var err error
	if configuration.GoogleAPIKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(configuration.GoogleAPIKey), maps.WithRateLimit(2))
	} else {
		_, err = fmt.Fprintln(os.Stderr, "Please specify an API Key.")
		os.Exit(2)
	}
	check(err)

	return client
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func getSecrets() (Configuration, error) {
	configuration := Configuration{}
	err := gonfig.GetConf("frontend/src/config/config.secret.json", &configuration)
	return configuration, err
}
