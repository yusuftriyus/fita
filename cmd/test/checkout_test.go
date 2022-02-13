package test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
	"github.com/yusuftriyus/fita/graph"
	"github.com/yusuftriyus/fita/graph/generated"
	"github.com/yusuftriyus/fita/graph/model"
	"testing"
)

func TestCheckout(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(graph.New())))

	t.Run("case #1 - checkout promo type 1", func(t *testing.T) {
		var removeCartResp struct {
			RemoveCart bool
		}
		c.MustPost(graph.RemoveCart, &removeCartResp)

		var cartResp struct {
			AddToCart *model.Cart
		}
		c.MustPost(graph.AddMacbook, &cartResp)
		c.MustPost(graph.AddRaspberry, &cartResp)

		var checkoutResp struct {
			Checkout *model.Checkout
		}
		c.MustPost(graph.Checkout, &checkoutResp)
		require.Equal(t, 5399.99, checkoutResp.Checkout.TotalPrice)
	})

	t.Run("case #2 - checkout promo type 2", func(t *testing.T) {
		var removeCartResp struct {
			RemoveCart bool
		}
		c.MustPost(graph.RemoveCart, &removeCartResp)

		var cartResp struct {
			AddToCart *model.Cart
		}
		c.MustPost(graph.AddGoogleHome, &cartResp)

		var checkoutResp struct {
			Checkout *model.Checkout
		}
		c.MustPost(graph.Checkout, &checkoutResp)
		require.Equal(t, 99.98, checkoutResp.Checkout.TotalPrice)
	})

	t.Run("case #3 - checkout promo type 3", func(t *testing.T) {
		var removeCartResp struct {
			RemoveCart bool
		}
		c.MustPost(graph.RemoveCart, &removeCartResp)

		var cartResp struct {
			AddToCart *model.Cart
		}
		c.MustPost(graph.AddAlexaSpeaker, &cartResp)

		var checkoutResp struct {
			Checkout *model.Checkout
		}
		c.MustPost(graph.Checkout, &checkoutResp)
		require.Equal(t, 295.65, checkoutResp.Checkout.TotalPrice)
	})

}
