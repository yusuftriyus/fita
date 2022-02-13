package graph

var RemoveCart = `
mutation remoteCart {
  removeCart
}
`

var AddMacbook = `
mutation addMacbook {
  addToCart(input: {sku: "43N23P", name: "Macbook Pro", qty: 1, price: 5399.99}) {
    items {
      sku
      name
      qty
      price
    }
  }
}
`

var AddRaspberry = `
mutation addRaspberry {
  addToCart(input: {sku: "234234", name: "Raspberry Pi B", qty: 1, price: 30.00}) {
    items {
      sku
      name
      qty
      price
    }
  }
}
`

var AddGoogleHome = `
mutation addGoogleHome {
  addToCart(input: {sku: "120P90", name: "Google Home", qty: 3, price: 49.99}) {
    items {
      sku
      name
      qty
      price
    }
  }
}
`

var AddAlexaSpeaker = `
mutation addAlexaSpeaker {
  addToCart(input: {sku: "A304SD", name: "Alexa Speaker", qty: 3, price: 109.50}) {
    items {
      sku
      name
      qty
      price
    }
  }
}
`

var Checkout = `
query checkout {
  checkout {
    items {
      sku
      name
      qty
      price
      reward_sku
      reward_qty
      disc_qty
      disc_perc
    }
    total_price
  }
}
`
