package strategy

import "testing"

func TestExamplePayByCash(t *testing.T) {
	ctx := NewPaymentContext("Ada", "", 123, &Cash{})
	ctx.Pay()
	// Output:
	// Pay $123 to Ada by cash
}

func TestExamplePayByBank(t *testing.T) {
	ctx := NewPaymentContext("Bob", "0002", 888, &Bank{})
	ctx.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
}
