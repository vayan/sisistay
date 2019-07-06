package createorder_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vayan/sisistay/src/api"
	"github.com/vayan/sisistay/src/test"
)

func TestCreateOrderApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cli Suite")
}

var _ = Describe("CreateOrder", func() {
	var apiConfig = api.Config{
		OrderStorage: test.OrderMockDB{},
	}

	BeforeEach(func() {
		apiConfig.CreateRoute()
	})

	Describe("POST /orders", func() {
		var request, _ = http.NewRequest(
			"POST",
			"/orders",
			strings.NewReader(`{}`),
		)

		Context("With payload", func() {
			It("returns the correct response", func() {
				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(200))
				Expect(response.Body.String()).To(Equal(""))
			})
		})
	})
})
