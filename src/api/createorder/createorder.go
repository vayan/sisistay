package createorder

import (
	"net/http"

	"github.com/vayan/sisistay/src/model"
)

type Controller struct {
	OrderStorage model.OrderStorage
}

func CreateController(dataStore model.OrderStorage) http.Handler {
	return &Controller{OrderStorage: dataStore}
}

func (controller *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpCode, body := controller.handleRequest(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(body)
}

func (controller *Controller) handleRequest(request *http.Request) (int, []byte) {
	return 200, []byte{}
}
