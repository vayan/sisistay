<p align="center">
<img width="300" alt="awesomelogo" src="https://user-images.githubusercontent.com/2945291/60754604-bedc1e00-9fe3-11e9-8c75-663934a2d903.png">
</p>

### Awesome API for Order Management

[![Build Status](https://travis-ci.com/vayan/sisistay.svg?branch=master)](https://travis-ci.com/vayan/sisistay)
[![codecov](https://codecov.io/gh/vayan/sisistay/branch/master/graph/badge.svg)](https://codecov.io/gh/vayan/sisistay)

Hello nice reviewer(s) :wave:

### Requirement

* Docker

### Get Started

Create a `.env` file at the root of the repo with
`GOOGLE_API_KEY=mygreatsecret`

`echo "GOOGLE_API_KEY=replaceme" > .env`

run `./start.sh` to run the docker-compose file to start a PostgreSQL
instance and a Go HTTP API.

API is available at `http://localhost:8080`

To run the tests clone the repo and do `go test ./...`

You can see the coverage on [codecov](https://codecov.io/gh/vayan/sisistay)


### Context for Reviewer


#### Assumptions

* For the distance I assumed we're using cars
* For the listing I ordered by ID

#### Dependencies

I tried to keep them minimal:

* `gorilla/mux` for HTTP routing
* `Gomega/Ginkgo` for testing
* `gorm` as an ORM
* Go Client for Google Maps Services


#### Improvements

The API controllers/handler could probably be refactored to have a bit less code
duplication

There's also no authentication

#### Misc

Coverage is only at 80%~~ because I didn't had time to do the setup to have
a DB for testing, else I would have unit tested all the database related
functions.

I also included my [postman collection](https://github.com/vayan/sisistay/blob/master/postman_collection.json) I used when developing this.
You guys must have yours but just in case :smile:


Thanks and have a good day :sunflower: !
