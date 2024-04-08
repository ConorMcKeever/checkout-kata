package pkg

import (
	"errors"
	"log"
	"strings"
)

type ICheckout interface {
	Scan(item string) error
	GetTotalPrice() int
}

type Deal struct {
	quantity int
	price    int
}

type Checkout struct {
	itemPrices map[string]int
	deals      map[string]Deal
	cart       map[string]int
}

func NewCheckout() *Checkout {
	itemPrices := map[string]int{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
	deals := map[string]Deal{
		"A": {quantity: 3, price: 130},
		"B": {quantity: 2, price: 45},
	}
	return &Checkout{itemPrices: itemPrices, deals: deals, cart: make(map[string]int)}
}

func (c *Checkout) Scan(item string) error {
	if len(item) != 1 {
		return errors.New("item should be a single character")
	}
	item = strings.ToUpper(item)

	if _, ok := c.itemPrices[item]; !ok {
		return errors.New("invalid item")
	}
	c.cart[item]++
	log.Printf("Adding item: %v to cart", item)
	return nil
}
