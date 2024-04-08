package tests

import (
	"checkout-kata/pkg"
	"testing"
)

func TestInput(t *testing.T) {
	checkout := pkg.NewCheckout()

	err := checkout.Scan("1")
	if err == nil {
		t.Errorf("Expected error for numeric input, got nil")
	}
	err = checkout.Scan("AB")
	if err == nil {
		t.Errorf("Expected error for multi character input, got nil")
	}
	err = checkout.Scan("A")
	if err != nil {
		t.Errorf("Expected no error for correct single character input, got %v", err)
	}
}

// func TestTotals(t *testing.T) {
// 	checkEquals(t, "", 0)
// 	checkEquals(t, "A", 50)
// 	checkEquals(t, "AB", 80)
// 	checkEquals(t, "CDBA", 115)
// }

// func checkEquals(t *testing.T, items string, expectedTotal int) {
// 	// code for checking total
// }

// func TestIncremental(t *testing.T) {
// 	co := NewCheckout()
// 	checkTotal(t, co, 0)
// 	co.Scan("A")
// 	checkTotal(t, co, 50)
// 	co.Scan("B")
// 	checkTotal(t, co, 80)
// 	co.Scan("A")
// 	checkTotal(t, co, 130)
// 	co.Scan("A")
// 	checkTotal(t, co, 160)
// 	co.Scan("B")
// 	checkTotal(t, co, 175)
// }

// func checkTotal(t *testing.T, co *Checkout, expectedTotal int) {
// 	// code for checking incremental total
// }
