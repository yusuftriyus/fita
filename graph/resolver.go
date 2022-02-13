package graph

import (
	"github.com/yusuftriyus/fita/graph/generated"
	"github.com/yusuftriyus/fita/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	checkout *model.Checkout
	cart     *model.Cart
	promos1  []model.Promo1
	promos2  []model.Promo2
	promos3  []model.Promo3
}

func New() generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{
			checkout: &model.Checkout{},
			cart:     &model.Cart{},
			promos1: []model.Promo1{
				{QualifierSku: "43N23P", QualifierQty: 1, RewardSku: "234234", RewardQty: 1},
			},
			promos2: []model.Promo2{
				{QualifierSku: "120P90", QualifierQty: 3, DiscQty: 1},
			},
			promos3: []model.Promo3{
				{QualifierSku: "A304SD", QualifierQty: 3, DiscPerc: 0.1},
			},
		},
	}
	return c
}
