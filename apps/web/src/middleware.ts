import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

const isPublicPage = (pathname: string) => {
  const publicPages = ['/auth'];
  return publicPages.some((page) => pathname.startsWith(page));
};

const isAuthenticated = (request: NextRequest) => {
  const token = request.cookies.get('token');
  console.log('isAuthenticated token', token);
  return !!token;
};

export function middleware(request: NextRequest) {
  const pathname = request.nextUrl.pathname;

  if (!isPublicPage(pathname) && !isAuthenticated(request)) {
    return NextResponse.redirect(new URL('/auth', request.url));
  }

  if (isPublicPage(pathname) && isAuthenticated(request)) {
    return NextResponse.redirect(new URL('/', request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
};
