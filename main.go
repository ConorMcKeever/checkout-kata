package checkout

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
