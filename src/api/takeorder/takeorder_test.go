package takeorder_test

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

func TestTakeOrderApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "take Suite")
}

var _ = Describe("TakeOrder", func() {
	var apiConfig api.Config

	BeforeEach(func() {
		apiConfig = api.Config{
			OrderStorage: test.OrderMockDB{
				FakeID: 1,
			},
			RouteFetcher: test.MockRouteFetcher{
				Distance: 2,
			},
		}
	})

	JustBeforeEach(func() {
		apiConfig.CreateRoute()
	})

	Describe("PATCH /orders/:id", func() {
		Context("With correct request payload", func() {
			var request *http.Request

			BeforeEach(func() {
				request, _ = http.NewRequest(
					"PATCH",
					"/orders/1",
					strings.NewReader(`{"status":"TAKEN"}`),
				)
			})

			It("returns success if order can be taken", func() {
				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(200))
				Expect(response.Body.String()).To(MatchJSON(`{"status":"SUCCESS"}`))
			})

			Context("when order cannot be taken", func() {
				BeforeEach(func() {
					apiConfig.OrderStorage = test.OrderMockDB{
						TakeError: errors.New("already taken"),
					}
				})

				It("returns error", func() {
					handler := apiConfig.GetHTTPHandler()

					response := httptest.NewRecorder()
					handler.ServeHTTP(response, request)

					Expect(response.Code).To(Equal(400))
					Expect(response.Body.String()).To(MatchJSON(`{"error":"CANNOT_BE_TAKEN"}`))
				})
			})
		})

		Context("With incorrect request", func() {
			It("returns an error if no payload", func() {
				var request, _ = http.NewRequest(
					"PATCH",
					"/orders/1",
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
					"PATCH",
					"/orders/1",
					strings.NewReader(`{"im":"invalid"}`),
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(400))
				Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
			})

			It("returns an error if incorrect status", func() {
				var request, _ = http.NewRequest(
					"PATCH",
					"/orders/1",
					strings.NewReader(`{"status":"STOLEN"}`),
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(400))
				Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
			})

			It("returns an error if incorrect id", func() {
				var request, _ = http.NewRequest(
					"PATCH",
					"/orders/imnotanid",
					strings.NewReader(`{"status":"TAKEN"}`),
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
