package listorder_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vayan/sisistay/src/model"

	"github.com/vayan/sisistay/src/api"
	"github.com/vayan/sisistay/src/test"
)

func TestListOrderApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "list order Suite")
}

var _ = Describe("ListOrder", func() {
	var apiConfig api.Config

	BeforeEach(func() {
		apiConfig = api.Config{
			OrderStorage: test.OrderMockDB{
				FakeID:           1,
				ListOrderResults: []model.Order{},
			},
			RouteFetcher: test.MockRouteFetcher{
				Distance: 2,
			},
		}
	})

	JustBeforeEach(func() {
		apiConfig.CreateRoute()
	})

	Describe("GET /orders", func() {
		Context("With correct request payload", func() {
			var request *http.Request

			BeforeEach(func() {
				request, _ = http.NewRequest(
					"GET",
					"/orders?page=1&limit=10",
					nil,
				)
			})

			Context("with results", func() {
				BeforeEach(func() {
					apiConfig.OrderStorage = test.OrderMockDB{
						ListOrderResults: []model.Order{
							{
								ID:             1,
								DistanceMeters: 2,
								Status:         model.OrderTaken,
							},
							{
								ID:             2,
								DistanceMeters: 2,
								Status:         model.OrderUnassigned,
							},
						},
					}
				})

				It("returns orders", func() {
					handler := apiConfig.GetHTTPHandler()

					response := httptest.NewRecorder()
					handler.ServeHTTP(response, request)

					Expect(response.Code).To(Equal(200))
					Expect(response.Body.String()).To(MatchJSON(`[{"id":1,"distance":2,"status":"TAKEN"},{"id":2,"distance":2,"status":"UNASSIGNED"}]`))
				})
			})

			Context("with no results", func() {
				It("returns empty array", func() {
					handler := apiConfig.GetHTTPHandler()

					response := httptest.NewRecorder()
					handler.ServeHTTP(response, request)

					Expect(response.Code).To(Equal(200))
					Expect(response.Body.String()).To(MatchJSON(`[]`))
				})
			})
		})

		Context("With incorrect request payload", func() {
			It("returns an error if page param is not a number", func() {
				var request, _ = http.NewRequest(
					"GET",
					"/orders?page=a&limit=10",
					nil,
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(400))
				Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
			})

			It("returns an error if page param is zero", func() {
				var request, _ = http.NewRequest(
					"GET",
					"/orders?page=0&limit=10",
					nil,
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(400))
				Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
			})

			It("returns an error if limit param is not a number", func() {
				var request, _ = http.NewRequest(
					"GET",
					"/orders?page=1&limit=notanumbergoddammit",
					nil,
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(400))
				Expect(response.Body.String()).To(MatchJSON(`{"error":"INVALID_PARAMS"}`))
			})

			It("returns an error if no param is given", func() {
				var request, _ = http.NewRequest(
					"GET",
					"/orders",
					nil,
				)

				handler := apiConfig.GetHTTPHandler()

				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(405))
				Expect(response.Body.String()).To(Equal(""))
			})
		})
	})
})
