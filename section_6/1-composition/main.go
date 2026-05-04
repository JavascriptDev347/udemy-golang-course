package main

import "fmt"

// composition => Has-A relationship

type Address struct {
	Street, City, State, ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	ID              int
	Name            string
	Email           string
	BillingAddress  Address // composition
	ShippingAddress Address
}

func (c Customer) PrintDetail() {
	fmt.Printf("Customer ID: %d\n", c.ID)
	fmt.Printf("Customer Name: %s\n", c.Name)
	fmt.Printf("Customer Email: %s\n", c.Email)
	fmt.Printf("Customer BillingAddress: %s\n", c.BillingAddress.FullAddress())
}

func main() {
	fmt.Println("Composition in Go")

	customer := Customer{
		ID:    1,
		Name:  "Russel",
		Email: "russel@gmail.com",
		BillingAddress: Address{
			Street:  "Makhmtumkuli 150",
			City:    "Tashkent",
			State:   "Tashkente",
			ZipCode: "100001",
		},
		ShippingAddress: Address{
			Street:  "Makhmtumkuli 151",
			City:    "Tashkent 2",
			State:   "Tashkente 2",
			ZipCode: "100003",
		},
	}
	customer.PrintDetail()

	fmt.Println("Customer with same address")

	customer2 := Customer{
		ID:    2,
		Name:  "Bob",
		Email: "Smith",
		BillingAddress: Address{
			Street:  "Farhad 150",
			City:    "Tashkent",
			State:   "Guzar",
			ZipCode: "100004",
		},
		ShippingAddress: Address{
			Street:  "Alisher Navoi 151",
			City:    "Tashkent 3",
			State:   "Tashkent 3",
			ZipCode: "10009",
		},
	}

	customer2.PrintDetail()
}
