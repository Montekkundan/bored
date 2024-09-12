import { NextResponse } from 'next/server';

export async function POST() {
  try {
    const apiResponse = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/rotate-token`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', // to include the refresh token cookie
    });

    const data = await apiResponse.json();

    if (!apiResponse.ok) {
      throw new Error(data.message || 'An error occurred');
    }

    // Set the new refresh token as an HTTP-only cookie
    const newRefreshToken = data.data.token.refresh_token;
    const cookieOptions = {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'strict' as const,
      maxAge: 7 * 24 * 60 * 60, // 7 days
      path: '/',
    };

    const response = NextResponse.json({ token: data.data.token.access_token });
    response.cookies.set('refresh_token', newRefreshToken, cookieOptions);

    return response;
  } catch (error) {
    return NextResponse.json(
      { message: error instanceof Error ? error.message : 'An error occurred' },
      { status: 500 }
    );
  }
}