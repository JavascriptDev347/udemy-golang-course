package main

import (
	"fmt"
	"strings"
)

var ProductPrices = map[string]float64{
	"Tshirt": 12.34,
	"MUG":    12.50,
	"Jeans":  89.09,
	"Book":   45.097,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, ok := ProductPrices[itemCode]
	if !ok {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found := ProductPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf(" - Item %s (Sale original price: %.2f, Sale price: %.2f)\n", originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}
		fmt.Printf(" - Item %s (Product not found)\n", itemCode)
		return 0.0, false
	}

	return basePrice, true
}

func main() {

	fmt.Println("Product Prices")
	orderItems := []string{
		"Tshirt",
		"MUG_SALE",
	}

	var subTotal float64

	for _, item := range orderItems {
		price, ok := calculateItemPrice(item)
		if ok {
			subTotal += price
		}
	}

	fmt.Printf("Sub total: %.2f\n", subTotal)
}
