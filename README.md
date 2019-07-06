<p align="center">
<img width="300" alt="awesomelogo" src="https://user-images.githubusercontent.com/2945291/60754604-bedc1e00-9fe3-11e9-8c75-663934a2d903.png">
</p>

### Awesome API for Order Management

Hello :wave:

### Requirement

* Go 1.12
* Docker

### Get Started

Create a `.env` file at the root of the repo with
`GOOGLE_API_KEY=mygreatsecret`

`echo "GOOGLE_API_KEY=replaceme" > .env`

run `./start.sh` to run the docker-compose file to start a PostgreSQL
instance and a Go HTTP API.

API is available at `http://localhost:8080`

To run tests and coverage clone the repo and do `go test ./... -cover`


### Context for Reviewer

#### Dependencies

I tried to keep then minimal:

* `gorilla/mux` for HTTP routing
* `Gomega/Ginkgo` for testing
* `gorm` as an ORM
* Go Client for Google Maps Services
