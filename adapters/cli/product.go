package cli

import (
	"fmt"
	"github.com/jvveiga/tests-arch-hexagonal/app"
)

func Run(service app.ProductServiceInterface, action string, productID string, productName string, price float32) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %d",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
	default:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s \nName: %s \nPrice: %f \nStatus: %d",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}
	return result, nil
}
