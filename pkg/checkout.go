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
	item_prices map[string]int
	deals       map[string]Deal
	cart        map[string]int
}

func NewCheckout() *Checkout {
	item_prices := map[string]int{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
	deals := map[string]Deal{
		"A": {quantity: 3, price: 130},
		"B": {quantity: 2, price: 45},
	}
	return &Checkout{item_prices: item_prices, deals: deals, cart: make(map[string]int)}
}

func (c *Checkout) Scan(item string) error {
	if len(item) != 1 {
		return errors.New("item should be a single character")
	}
	item = strings.ToUpper(item)

	if _, ok := c.item_prices[item]; !ok {
		return errors.New("invalid item")
	}
	c.cart[item]++
	log.Printf("Adding item: %v to cart", item)
	return nil
}

func (c *Checkout) GetTotalPrice() int {
	total := 0
	for item, quantity := range c.cart {
		deal, has_deal := c.deals[item]
		item_price := c.item_prices[item]
		if has_deal {
			total += (quantity/deal.quantity)*deal.price +
				(quantity%deal.quantity)*item_price
		} else {
			total += item_price * quantity
		}
	}
	return total
}
