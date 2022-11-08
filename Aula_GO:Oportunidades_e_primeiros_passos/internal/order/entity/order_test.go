package entity

import(
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenSouldReceiveAnError(t *testing.T) {
		order  := Order {}
		assert.Error(t, order.IsValid(), "invalid id")
} 

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenSouldReceiveAnError(t *testing.T) {
	order  := Order {ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
} 

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenSouldReceiveAnError(t *testing.T) {
	order  := Order {ID: "123", Price: 100}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func  TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{
		ID: "123",
		Price: 100,
		Tax: 10,
		FinalPrice: 110,
	}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)
}

func TestGivenAnPriceAndTax_WhenICallCalculatePrice_THenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculatePrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}

