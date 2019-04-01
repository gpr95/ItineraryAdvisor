# Itinerary advisor

Studies project for spatial databases.

### Prerequisite
- Go (1.12.1)

### Install
Put your Google API KEY in file
```bash
$PROJECT_ROOT/config/config.secret.json
```
3trd party libraries:
```
$ go get github.com/tkanos/gonfig
$ go get github.com/tkanos/gonfig
$ go get googlemaps.github.io/maps
```
Run:
```
// Run whole directory
$ go run main.go route.go
```