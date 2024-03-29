package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vayan/sisistay/src/api/createorder"
	"github.com/vayan/sisistay/src/api/listorder"
	"github.com/vayan/sisistay/src/api/takeorder"
	"github.com/vayan/sisistay/src/model"
	"github.com/vayan/sisistay/src/service"
)

type Config struct {
	Port         string
	OrderStorage model.OrderStorage
	RouteFetcher service.RouteFetcher
	router       *mux.Router
}

func (c Config) InitDB() {
	c.OrderStorage.Migrate()
}

func (c *Config) CreateRoute() {
	c.router = mux.NewRouter()

	c.router.Handle(
		"/orders",
		createorder.CreateController(c.OrderStorage, c.RouteFetcher),
	).Methods("POST")

	c.router.Handle(
		"/orders/{orderID}",
		takeorder.CreateController(c.OrderStorage),
	).Methods("PATCH")

	c.router.Handle(
		"/orders",
		listorder.CreateController(c.OrderStorage),
	).Methods("GET").
		Queries("page", "{page:.*}", "limit", "{limit:.*}")

}

func (c *Config) GetHTTPHandler() http.Handler {
	return c.router
}

func (c Config) ListenAndServe() error {
	c.CreateRoute()

	server := &http.Server{
		Handler:      c.router,
		Addr:         ":" + c.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server.ListenAndServe()
}
