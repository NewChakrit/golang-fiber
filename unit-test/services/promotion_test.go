package services_test

import (
	"testing"
	"unit-test/repositories"
	"unit-test/services"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type TestCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []TestCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "not applied 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
		// {name: "not applied 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 30}, // Test False
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arrange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promoService := services.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoService.CalculateDiscount(c.amount)
			expected := c.expected

			// Assert
			assert.Equal(t, expected, discount)
		})

	}

}
