import { NextResponse } from 'next/server';

export async function GET(request: Request) {
  const token = request.headers.get('Authorization');

  if (!token) {
    return NextResponse.json({ message: 'Unauthorized' }, { status: 401 });
  }

  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}/users/get-all`,
      {
        method: 'GET',
        headers: {
          Authorization: token,
          'Content-Type': 'application/json',
        },
      }
    );

    const responseText = await response.text();
    let data;
    try {
      data = JSON.parse(responseText);
    } catch (parseError) {
      console.error('Failed to parse response as JSON:', responseText);
      return NextResponse.json(
        { message: 'Invalid response from server', details: responseText },
        { status: 500 }
      );
    }

    if (!response.ok) {
      throw new Error(
        data.message || `Failed to fetch users: ${response.statusText}`
      );
    }

    return NextResponse.json(data);
  } catch (error) {
    console.error('Error in GET request:', error);
    return NextResponse.json(
      {
        message: error instanceof Error ? error.message : 'An error occurred',
        details: error,
      },
      { status: 500 }
    );
  }
}
