package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorWhenIdIsBlank(t *testing.T) {
	order := Order{}
	// err := order.Validate()
	// if err == nil {
	// 	t.Errorf("Expected an error, got nil")
	// }
	assert.Error(t, order.Validate(), "Invalid id")
}

func TestIfItGetsAnErrorWhenPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "Invalid Price")
}

func TestIfItGetsAnErrorWhenTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "Invalid tax")
}

func TestWithAllValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
