package createorder_test

import (
	"errors"
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
	var apiConfig api.Config

	BeforeEach(func() {
		apiConfig = api.Config{
			OrderStorage: test.OrderMockDB{
				FakeID: 1,
			},
			RouteFetcher: test.MockRouteFetcher{
				Distance: 2,
				Error:    nil,
			},
		}
	})

	JustBeforeEach(func() {
		apiConfig.CreateRoute()
	})

	Describe("POST /orders", func() {
		Context("With correct request payload", func() {
			var request *http.Request

			BeforeEach(func() {
				request, _ = http.NewRequest(
					"POST",
					"/orders",
					strings.NewReader(`{"origin":["11.11", "22.22"],"destination":["11.22","22.22"]}`),
				)
			})

			It("returns the correct response", func() {
				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(200))
				Expect(response.Body.String()).To(MatchJSON(`{"id":1,"distance":2,"status":"UNASSIGNED"}`))
			})

			Context("When no route is found", func() {
				BeforeEach(func() {
					apiConfig.RouteFetcher = test.MockRouteFetcher{
						Error: errors.New("cosmic radiation"),
					}
				})

				It("returns the correct error", func() {
					handler := apiConfig.GetHTTPHandler()

					response := httptest.NewRecorder()
					handler.ServeHTTP(response, request)

					Expect(response.Code).To(Equal(400))
					Expect(response.Body.String()).To(MatchJSON(`{"error":"NO_ROUTE_AVAILABLE"}`))
				})
			})
		})

		Context("With invalid request payload", func() {
			It("returns an error if no payload", func() {
				var request, _ = http.NewRequest(
					"POST",
					"/orders",
					nil,
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(400))
				Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
			})

			It("returns an error if invalid payload", func() {
				var request, _ = http.NewRequest(
					"POST",
					"/orders",
					strings.NewReader(`{"foo": "bar""}`),
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(400))
				Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
			})

			Context("With invalid coordinates", func() {
				It("returns an error if there's more than two coordinate for the origin", func() {
					var request, _ = http.NewRequest(
						"POST",
						"/orders",
						strings.NewReader(`{"origin":["11.11", "22.22", "33,33"],"destination":["11.22","22.22"]}`),
					)

					handler := apiConfig.GetHTTPHandler()

					response := httptest.NewRecorder()
					handler.ServeHTTP(response, request)

					Expect(response.Code).To(Equal(400))
					Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
				})

				It("returns an error if there's more than two coordinate for the destination", func() {
					var request, _ = http.NewRequest(
						"POST",
						"/orders",
						strings.NewReader(`{"origin":["11.11", "22.22"],"destination":["11.22","22.22", "33.33"]}`),
					)

					handler := apiConfig.GetHTTPHandler()

					response := httptest.NewRecorder()
					handler.ServeHTTP(response, request)

					Expect(response.Code).To(Equal(400))
					Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
				})
			})
		})
	})
})
