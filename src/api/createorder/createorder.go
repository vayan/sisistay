package createorder

import (
	"net/http"

	"github.com/vayan/sisistay/src/api/apiutils"
	"github.com/vayan/sisistay/src/model"
)

type Controller struct {
	OrderStorage model.OrderStorage
}

type RequestPayload struct {
	Origin      model.Coordinates `json:"origin"`
	Destination model.Coordinates `json:"destination"`
}

func (rp RequestPayload) Valid() bool {
	return rp.Origin.Valid() && rp.Destination.Valid()
}

func CreateController(dataStore model.OrderStorage) http.Handler {
	return &Controller{OrderStorage: dataStore}
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpCode, body := c.handleRequest(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(body)
}

func (c *Controller) handleRequest(request *http.Request) (int, []byte) {
	var requestPayload RequestPayload

	err := apiutils.ParseTo(request.Body, &requestPayload)

	if err != nil || !requestPayload.Valid() {
		return http.StatusBadRequest, apiutils.Serialize(model.ErrorResponse{
			Error: "INVALID_PARAMS",
		})
	}

	order := model.Order{
		Status:         model.OrderUnassigned,
		DistanceMeters: 2, // TODO get from Google API
	}

	c.OrderStorage.Create(&order)

	return 200, apiutils.Serialize(order)
}
