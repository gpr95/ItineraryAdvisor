# Itinerary advisor

Studies project for spatial databases.

### Prerequisite
- Go (1.12.1)

### Install
#### Backend
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
$ go run main.go
```

#### Frontend

There are two options:
- Run separately 
```
$ cd frontend
$ npm install
$ npm start
# Frontend will be avaliable on port :3000 and requests will be passed to go backend at :8000
```

- Run combined 
```
$ cd frontend
$ npm install
$ npm dist
# Frontend will be avaliable on :8000 - make sure that go backend is running
```