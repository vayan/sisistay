package createorder

import (
	"net/http"

	"github.com/vayan/sisistay/src/api/apiutils"
	"github.com/vayan/sisistay/src/model"
	"github.com/vayan/sisistay/src/service"
)

type controller struct {
	OrderStorage model.OrderStorage
	RouteFetcher service.RouteFetcher
}

type RequestPayload struct {
	Origin      model.Coordinates `json:"origin"`
	Destination model.Coordinates `json:"destination"`
}

func (rp RequestPayload) Valid() bool {
	return rp.Origin.Valid() && rp.Destination.Valid()
}

func CreateController(dataStore model.OrderStorage, routeFetcher service.RouteFetcher) http.Handler {
	return &controller{
		OrderStorage: dataStore,
		RouteFetcher: routeFetcher,
	}
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpCode, body := c.handleRequest(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(body)
}

func (c *controller) handleRequest(request *http.Request) (int, []byte) {
	var requestPayload RequestPayload

	err := apiutils.ParseTo(request.Body, &requestPayload)

	if err != nil || !requestPayload.Valid() {
		return http.StatusBadRequest, model.SerializedErrorResponse("INVALID_PARAMS")
	}

	distance, err := c.RouteFetcher.GetDistance(requestPayload.Origin, requestPayload.Destination)

	if err != nil {
		return http.StatusBadRequest, model.SerializedErrorResponse("NO_ROUTE_AVAILABLE")
	}

	order := model.Order{
		Status:         model.OrderUnassigned,
		DistanceMeters: distance,
	}

	c.OrderStorage.Create(&order)

	return 200, apiutils.Serialize(order)
}
