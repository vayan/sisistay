package model_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vayan/sisistay/src/model"
)

func TestCoordinateModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coordinate Suite")
}

var _ = Describe("Coordinate", func() {
	Describe(".Valid()", func() {
		Context("for valid coordinate", func() {
			It("returns true", func() {
				Expect(model.Coordinates{"-10.1", "3.1"}.Valid()).To(BeTrue())
			})
		})

		Context("for invalid coordinate", func() {
			It("returns false if more than two points", func() {
				Expect(model.Coordinates{"-10.1", "3.1", "3.3"}.Valid()).To(BeFalse())
			})

			It("returns false if less than two points", func() {
				Expect(model.Coordinates{"-10.1"}.Valid()).To(BeFalse())
			})

			It("returns false if lat is not a number", func() {
				Expect(model.Coordinates{"notanumberman", "11"}.Valid()).To(BeFalse())
			})

			It("returns false if long is not a number", func() {
				Expect(model.Coordinates{"11", "meneither"}.Valid()).To(BeFalse())
			})

			It("returns false if lat is out of bound", func() {
				Expect(model.Coordinates{"-91", "11"}.Valid()).To(BeFalse())
				Expect(model.Coordinates{"91", "11"}.Valid()).To(BeFalse())
			})

			It("returns false if long is out of bound", func() {
				Expect(model.Coordinates{"11", "-181"}.Valid()).To(BeFalse())
				Expect(model.Coordinates{"11", "181"}.Valid()).To(BeFalse())
			})
		})
	})
})
