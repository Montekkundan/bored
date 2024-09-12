import { NextResponse } from 'next/server';
import { cookies } from 'next/headers';

export async function POST() {
  try {
    const cookieStore = cookies();
    const refreshToken = cookieStore.get('refresh_token');

    if (!refreshToken) {
      return NextResponse.json(
        { status: 'fail', message: 'No refresh token provided' },
        { status: 400 }
      );
    }

    const response = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}/auth/logout`,
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Cookie: `refresh_token=${refreshToken.value}`,
        },
      }
    );

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Logout failed');
    }

    cookies().delete('token');
    cookies().delete('refresh_token');

    return NextResponse.json(data);
  } catch (error) {
    console.error('Logout error:', error);
    return NextResponse.json(
      {
        status: 'fail',
        message:
          error instanceof Error
            ? error.message
            : 'An error occurred during logout',
      },
      { status: 500 }
    );
  }
}
