import React, { useEffect, useState } from 'react'
import {
  getItems,
  addItemToCart,
  getCarts,
  createOrder,
  getOrders
} from '../api'

const topBarStyle = {
  display: 'flex',
  gap: '12px',
  marginBottom: '16px'
}

const buttonStyle = {
  padding: '8px 12px',
  fontSize: '14px',
  cursor: 'pointer'
}

const itemListStyle = {
  display: 'flex',
  flexWrap: 'wrap',
  gap: '12px'
}

const itemCardStyle = {
  border: '1px solid #ccc',
  padding: '12px',
  cursor: 'pointer',
  minWidth: '120px'
}

export default function ItemsPage({ token, onLogout }) {
  const [items, setItems] = useState([])
  const [loadingItems, setLoadingItems] = useState(false)

  useEffect(() => {
    fetchItems()
  }, [])

  const fetchItems = async () => {
    setLoadingItems(true)
    try {
      const data = await getItems()
      setItems(data.items || [])
    } catch (err) {
      window.alert('Failed to load items')
    } finally {
      setLoadingItems(false)
    }
  }

  const handleItemClick = async (itemId) => {
    try {
      await addItemToCart(token, itemId)
      window.alert('Item added to cart')
    } catch (err) {
      window.alert(err.message || 'Failed to add item to cart')
    }
  }

  const handleShowCart = async () => {
    try {
      const data = await getCarts(token)
      const carts = data.carts || []

      if (!carts.length) {
        window.alert('Cart is empty')
        return
      }

      // Pick the latest open cart, otherwise last cart
      let currentCart = carts.find((c) => c.status === 'open')
      if (!currentCart) {
        currentCart = carts[carts.length - 1]
      }

      const itemsList = (currentCart.cart_items || []).map(
        (ci) => `cart_id: ${ci.cart_id}, item_id: ${ci.item_id}`
      )

      if (!itemsList.length) {
        window.alert('Cart is empty')
        return
      }

      window.alert(itemsList.join('\n'))
    } catch (err) {
      window.alert('Failed to load cart')
    }
  }

  const handleShowOrders = async () => {
    try {
      const data = await getOrders(token)
      const orders = data.orders || []

      if (!orders.length) {
        window.alert('No orders yet')
        return
      }

      const lines = orders.map((o) => `Order ID: ${o.id}`)
      window.alert(lines.join('\n'))
    } catch (err) {
      window.alert('Failed to load orders')
    }
  }

  const handleCheckout = async () => {
    try {
      const data = await getCarts(token)
      const carts = data.carts || []

      if (!carts.length) {
        window.alert('No cart available to checkout')
        return
      }

      let currentCart = carts.find((c) => c.status === 'open')
      if (!currentCart) {
        currentCart = carts[carts.length - 1]
      }

      if (!currentCart || !currentCart.id) {
        window.alert('No valid cart to checkout')
        return
      }

      await createOrder(token, currentCart.id)
      window.alert('Order successful')
    } catch (err) {
      window.alert(err.message || 'Checkout failed')
    }
  }

  return (
    <div>
      <div style={topBarStyle}>
        <button style={buttonStyle} onClick={handleCheckout}>
          Checkout
        </button>
        <button style={buttonStyle} onClick={handleShowCart}>
          Cart
        </button>
        <button style={buttonStyle} onClick={handleShowOrders}>
          Order History
        </button>
        <button
          style={{ ...buttonStyle, marginLeft: 'auto' }}
          onClick={onLogout}
        >
          Logout
        </button>
      </div>

      <h2>Items</h2>
      {loadingItems && <p>Loading items...</p>}
      {!loadingItems && !items.length && <p>No items found.</p>}

      <div style={itemListStyle}>
        {items.map((item) => (
          <div
            key={item.id}
            style={itemCardStyle}
            onClick={() => handleItemClick(item.id)}
          >
            <strong>{item.name}</strong>
            <div>Status: {item.status || 'N/A'}</div>
          </div>
        ))}
      </div>
    </div>
  )
}
