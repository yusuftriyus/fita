package model

import (
	"context"
)

func (c *Checkout) CalculateTotal(ctx context.Context) {
	for _, item := range c.Items {
		subsubtotal := float64(item.Qty-item.DiscQty) * item.Price
		discountTotal := float64(item.Qty-item.DiscQty) * item.DiscPerc * item.Price
		subtotal := subsubtotal - discountTotal

		c.TotalPrice += subtotal
	}
}
