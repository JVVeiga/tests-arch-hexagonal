package app_test

import (
	"github.com/jvveiga/tests-arch-hexagonal/app"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Product 1"
	product.Status = app.ENABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "Product 1"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 1
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disable", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = app.ENABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = 2
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}
