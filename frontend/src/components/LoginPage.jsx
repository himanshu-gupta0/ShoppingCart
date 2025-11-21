import React, { useState } from 'react'
import { login } from '../api'

const formStyle = {
  display: 'flex',
  flexDirection: 'column',
  gap: '12px',
  maxWidth: '300px'
}

const inputStyle = {
  padding: '8px',
  fontSize: '14px'
}

const buttonStyle = {
  padding: '10px',
  fontSize: '14px',
  cursor: 'pointer'
}

export default function LoginPage({ onLoginSuccess }) {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e) => {
    e.preventDefault()
    setLoading(true)
    try {
      const data = await login(username, password)
      if (data && data.token) {
        onLoginSuccess(data.token)
      } else {
        window.alert('Invalid username/password')
      }
    } catch (err) {
      window.alert('Invalid username/password')
    } finally {
      setLoading(false)
    }
  }

  return (
    <form style={formStyle} onSubmit={handleSubmit}>
      <h2>Login</h2>
      <input
        style={inputStyle}
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        required
      />
      <input
        style={inputStyle}
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
      />
      <button style={buttonStyle} type="submit" disabled={loading}>
        {loading ? 'Logging in...' : 'Login'}
      </button>
    </form>
  )
}
