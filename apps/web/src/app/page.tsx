"use client"

import { useState, useEffect } from 'react'
import { Input } from "@repo/ui/components/ui/input"
import { Label } from "@repo/ui/components/ui/label"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@repo/ui/components/ui/card"
import { Button } from '@repo/ui/components/ui/button'

interface User {
  id: number
  username: string
  email: string
  roles: string[]
  created_at: string
  bio: string
  interests: string[] | null
  latitude: number
  longitude: number
  profile_picture: string
  cover_photo: string
  social_links: string
  oauth_providers: string[] | null
  audio_enabled: boolean
  video_enabled: boolean
  email_verified: boolean
  phone_number: string
  phone_verified: boolean
  two_factor_enabled: boolean
  reward_points: number
  updated_at: string
}

interface LoginResponse {
  data: {
    token: string
    user: User
  }
  message: string
  status: string
}

export default function UserAuthPage() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [token, setToken] = useState<string | null>(null)
  const [currentUser, setCurrentUser] = useState<User | null>(null)
  const [users, setUsers] = useState<User[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const storedToken = localStorage.getItem('token')
    const storedUser = localStorage.getItem('user')
    
    if (storedToken && storedUser) {
      setToken(storedToken)
      setCurrentUser(JSON.parse(storedUser))
    }
    setLoading(false)
  }, [])

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    try {
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })
      const data: LoginResponse = await response.json()
      if (!response.ok) {
        throw new Error(data.message || 'An error occurred during login')
      }
      setToken(data.data.token)
      setCurrentUser(data.data.user)
      localStorage.setItem('token', data.data.token)
      localStorage.setItem('user', JSON.stringify(data.data.user))
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred during login')
    } finally {
      setLoading(false)
    }
  }

  const handleLogout = async () => {
    setLoading(true)
    try {
      const response = await fetch('/api/auth/logout', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })
      const data = await response.json()
      if (!response.ok) {
        throw new Error(data.message || 'An error occurred during logout')
      }
      setToken(null)
      setCurrentUser(null)
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred during logout')
    } finally {
      setLoading(false)
    }
  }

  const fetchUsers = async () => {
    if (!token) {
      setError('Please login first')
      return
    }
    setLoading(true)
    try {
      const response = await fetch(`/api/users`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      if (!response.ok) {
        throw new Error('Failed to fetch users')
      }
      const data = await response.json()
      setUsers(data.data)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred')
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <div>Loading...</div>
  }

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">User Authentication</h1>
      {!currentUser ? (
        <Card className="w-full max-w-md mx-auto">
          <CardHeader>
            <CardTitle>Login</CardTitle>
            <CardDescription>Enter your credentials to access the user list.</CardDescription>
          </CardHeader>
          <form onSubmit={handleLogin}>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
              </div>
            </CardContent>
            <CardFooter>
              <Button type="submit" disabled={loading}>
                {loading ? 'Logging in...' : 'Login'}
              </Button>
            </CardFooter>
          </form>
        </Card>
      ) : (
        <div className="space-y-4">
          <p className="text-lg">Welcome, {currentUser.username}!</p>
          <Button onClick={fetchUsers} disabled={loading}>
            {loading ? 'Fetching Users...' : 'Get Users'}
          </Button>
          <Button onClick={handleLogout} disabled={loading}>
            {loading ? 'Logging out...' : 'Logout'}
          </Button>
        </div>
      )}
      {error && <p className="text-red-500 mt-4">{error}</p>}
      {users.length > 0 && (
        <div className="mt-8">
          <h2 className="text-xl font-bold mb-4">All Users</h2>
          <ul className="space-y-4">
            {users.map((user) => (
              <li key={user.id} className="border p-4 rounded-lg">
                <h3 className="text-lg font-semibold">{user.username}</h3>
                <p>Email: {user.email}</p>
                <p>Roles: {user.roles.join(', ')}</p>
                <p>Created at: {new Date(user.created_at).toLocaleString()}</p>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  )
}