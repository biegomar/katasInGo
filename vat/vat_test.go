package vat

import (
	"fmt"
	"testing"
)

func ExampleRenderInvoice() {
	invoice := RenderInvoice(100, .19)
	fmt.Println(invoice)
	// Output:
	// Netto: 100.00 €
	// VAT (0.19): 19.00 €
	// Total: 119.00 €
}

func TestRenderInvoice(t *testing.T) {
	testCases := []struct {
		price    float64
		vatRate  float64
		expected string
	}{
		{100.00, 0.19, "Netto: 100.00 €\nVAT (0.19): 19.00 €\nTotal: 119.00 €"},
		{100.00, .21, "Netto: 100.00 €\nVAT (0.21): 21.00 €\nTotal: 121.00 €"},
		{99.99, 0.19, "Netto: 99.99 €\nVAT (0.19): 19.00 €\nTotal: 118.99 €"},
	}

	for _, tc := range testCases {
		actual := RenderInvoice(tc.price, tc.vatRate)
		if actual != tc.expected {
			t.Fatalf("[%.2f, %.2f]\nExpected:\n%s\ngot:\n%s", tc.price, tc.vatRate, tc.expected, actual)
		}
	}
}
