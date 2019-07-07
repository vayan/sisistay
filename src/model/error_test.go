package model_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vayan/sisistay/src/model"
)

func TestErrorModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Error Suite")
}

var _ = Describe("Error", func() {
	Describe(".SerializedErrorResponse()", func() {
		It("serialize an error API response", func() {
			Expect(model.SerializedErrorResponse("great error")).To(MatchJSON(`{"error":"great error"}`))
		})
	})
})
