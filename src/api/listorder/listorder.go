package listorder

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

	page, pageErr := strconv.Atoi(vars["page"])
	limit, limitErr := strconv.Atoi(vars["limit"])

	if pageErr != nil || limitErr != nil || page == 0 {
		return http.StatusBadRequest, model.SerializedErrorResponse("INVALID_PARAMS")
	}

	orders := c.OrderStorage.List(page, limit)

	return 200, apiutils.Serialize(orders)
}
