const BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://127.0.0.1:8081/api'

async function getToken() {
  if (typeof window !== 'undefined') {
    return localStorage.getItem('token')
  }
  // For server-side requests, you might want to implement a different way to get the token
  return null
}

async function client(endpoint: string, { body, ...customConfig }: RequestInit = {}) {
  const token = await getToken()
  const headers = { 'Content-Type': 'application/json' }

  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const config: RequestInit = {
    method: body ? 'POST' : 'GET',
    ...customConfig,
    headers: {
      ...headers,
      ...customConfig.headers,
    },
  }

  if (body) {
    config.body = JSON.stringify(body)
  }

  try {
    const response = await fetch(`${BASE_URL}${endpoint}`, config)
    const data = await response.json()

    if (response.ok) {
      return data
    } else {
      return Promise.reject(data)
    }
  } catch (error) {
    return Promise.reject(error)
  }
}

export { client }