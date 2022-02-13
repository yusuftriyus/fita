package graph

import (
	"context"
	"github.com/yusuftriyus/fita/graph/model"
)

func AdjustQty(ctx context.Context, cItem *model.CheckoutItem, checkout *model.Checkout) {
	for _, item := range checkout.Items {
		if cItem.Sku == item.RewardSku {
			cItem.Qty -= item.RewardQty
			break
		}
	}
}

func GetPromo1(ctx context.Context, item *model.CheckoutItem, promo1 []model.Promo1) {
	for _, promo := range promo1 {
		if item.Sku == promo.QualifierSku && item.Qty >= promo.QualifierQty {
			item.RewardSku = promo.RewardSku
			item.RewardQty = promo.RewardQty
			break
		}
	}
}

func GetPromo2(ctx context.Context, item *model.CheckoutItem, promos2 []model.Promo2) {
	for _, promo := range promos2 {
		if item.Sku == promo.QualifierSku && item.Qty >= promo.QualifierQty {
			item.DiscQty = promo.DiscQty
			break
		}
	}
}

func GetPromo3(ctx context.Context, item *model.CheckoutItem, promos3 []model.Promo3) {
	for _, promo := range promos3 {
		if item.Sku == promo.QualifierSku && item.Qty >= promo.QualifierQty {
			item.DiscPerc = promo.DiscPerc
			break
		}
	}
}
