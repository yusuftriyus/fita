package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/yusuftriyus/fita/graph/generated"
	"github.com/yusuftriyus/fita/graph/model"
)

func (r *mutationResolver) AddPromo1(ctx context.Context, input model.NewPromo1) (*model.Promo1, error) {
	newPromo := model.Promo1{
		QualifierSku: input.QualifierSku,
		QualifierQty: input.QualifierQty,
		RewardSku:    input.RewardSku,
		RewardQty:    input.RewardQty,
	}
	r.promos1 = append(r.promos1, newPromo)
	return &newPromo, nil
}

func (r *mutationResolver) AddPromo2(ctx context.Context, input model.NewPromo2) (*model.Promo2, error) {
	newPromo := model.Promo2{
		QualifierSku: input.QualifierSku,
		QualifierQty: input.QualifierQty,
		DiscQty:      input.DiscQty,
	}
	r.promos2 = append(r.promos2, newPromo)
	return &newPromo, nil
}

func (r *mutationResolver) AddPromo3(ctx context.Context, input model.NewPromo3) (*model.Promo3, error) {
	newPromo := model.Promo3{
		QualifierSku: input.QualifierSku,
		QualifierQty: input.QualifierQty,
		DiscPerc:     input.DiscPerc,
	}
	r.promos3 = append(r.promos3, newPromo)
	return &newPromo, nil
}

func (r *mutationResolver) AddToCart(ctx context.Context, input model.NewItem) (*model.Cart, error) {
	newItem := &model.Item{
		Sku:   input.Sku,
		Name:  input.Name,
		Qty:   input.Qty,
		Price: input.Price,
	}

	for _, item := range r.cart.Items {
		if newItem.Sku == item.Sku {
			item.Qty = newItem.Qty
			return r.cart, nil
		}
	}

	r.cart.Items = append(r.cart.Items, newItem)
	return r.cart, nil
}

func (r *mutationResolver) RemoveCart(ctx context.Context) (bool, error) {
	emptyCart := &model.Cart{}
	r.cart = emptyCart
	return true, nil
}

func (r *queryResolver) Checkout(ctx context.Context) (*model.Checkout, error) {
	r.checkout = &model.Checkout{}
	for _, item := range r.cart.Items {
		checkoutItem := &model.CheckoutItem{
			Sku:   item.Sku,
			Name:  item.Name,
			Qty:   item.Qty,
			Price: item.Price,
		}

		AdjustQty(ctx, checkoutItem, r.checkout)
		if checkoutItem.Qty <= 0 {
			continue
		}

		GetPromo1(ctx, checkoutItem, r.promos1)
		GetPromo2(ctx, checkoutItem, r.promos2)
		GetPromo3(ctx, checkoutItem, r.promos3)

		r.checkout.Items = append(r.checkout.Items, checkoutItem)
	}

	r.checkout.CalculateTotal(ctx)

	return r.checkout, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
