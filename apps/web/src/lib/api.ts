async function refreshAccessToken(): Promise<string | null> {
  try {
    const response = await fetch('/api/auth/refresh', {
      method: 'POST',
      credentials: 'include', // to include the refresh token cookie
    });

    if (response.ok) {
      const data = await response.json();
      localStorage.setItem('token', data.token);
      return data.token;
    }
  } catch (error) {
    console.error('Failed to refresh token:', error);
  }

  // If refresh failed, clear the stored token and user
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  return null;
}

export async function fetchWithAuth(url: string, options: RequestInit = {}) {
  const token = localStorage.getItem('token');

  if (token) {
    options.headers = {
      ...options.headers,
      Authorization: `Bearer ${token}`,
    };
  }

  let response = await fetch(url, options);

  if (response.status === 401) {
    // Token has expired, try to refresh it
    const newToken = await refreshAccessToken();
    if (newToken) {
      // Retry the request with the new token
      options.headers = {
        ...options.headers,
        Authorization: `Bearer ${newToken}`,
      };
      response = await fetch(url, options);
    }
  }

  return response;
}
