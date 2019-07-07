package takeorder

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vayan/sisistay/src/api/apiutils"
	"github.com/vayan/sisistay/src/model"
)

type controller struct {
	OrderStorage model.OrderStorage
}

type RequestValidResponse struct {
	Status string `json:"status"`
}

type RequestPayload struct {
	Status model.OrderStatus `json:"status"`
}

func (rp RequestPayload) Valid() bool {
	return rp.Status == model.OrderTaken
}

func CreateController(dataStore model.OrderStorage) http.Handler {
	return &controller{
		OrderStorage: dataStore,
	}
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpCode, body := c.handleRequest(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(body)
}

func (c *controller) handleRequest(request *http.Request) (int, []byte) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["orderID"])

	if err != nil {
		return http.StatusBadRequest, model.SerializedErrorResponse("INVALID_PARAMS")
	}

	var requestPayload RequestPayload

	err = apiutils.ParseTo(request.Body, &requestPayload)

	if err != nil || !requestPayload.Valid() {
		return http.StatusBadRequest, model.SerializedErrorResponse("INVALID_PARAMS")
	}

	err = c.OrderStorage.Take(uint(id))

	if err != nil {
		return http.StatusBadRequest, model.SerializedErrorResponse("CANNOT_BE_TAKEN")
	}

	return 200, apiutils.Serialize(RequestValidResponse{
		Status: "SUCCESS",
	})
}
