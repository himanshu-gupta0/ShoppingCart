import React, { useState } from 'react'
import LoginPage from './components/LoginPage'
import ItemsPage from './components/ItemsPage'

const appContainerStyle = {
  fontFamily: 'system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif',
  padding: '24px',
  maxWidth: '800px',
  margin: '0 auto'
}

export default function App() {
  const [token, setToken] = useState(() => localStorage.getItem('token') || '')
  const [view, setView] = useState(() => (token ? 'items' : 'login'))

  const handleLoginSuccess = (newToken) => {
    setToken(newToken)
    localStorage.setItem('token', newToken)
    setView('items')
  }

  const handleLogout = () => {
    setToken('')
    localStorage.removeItem('token')
    setView('login')
  }

  return (
    <div style={appContainerStyle}>
      <h1>Shopping Cart</h1>
      {view === 'login' && <LoginPage onLoginSuccess={handleLoginSuccess} />}
      {view === 'items' && <ItemsPage token={token} onLogout={handleLogout} />}
    </div>
  )
}
