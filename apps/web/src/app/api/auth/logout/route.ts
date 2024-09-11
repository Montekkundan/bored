import { NextResponse } from 'next/server';

export async function POST(request: Request) {
  try {
    const token = request.headers.get('Authorization');
    
    if (!token) {
      return NextResponse.json({ status: 'fail', message: 'No token provided' }, { status: 401 });
    }

    //  invalidate the token on your backend
    //  simulate a successful logout

    return NextResponse.json({
      status: 'success',
      message: 'Logout successful'
    });
  } catch (error) {
    return NextResponse.json(
      { status: 'fail', message: error instanceof Error ? error.message : 'An error occurred' },
      { status: 500 }
    );
  }
}