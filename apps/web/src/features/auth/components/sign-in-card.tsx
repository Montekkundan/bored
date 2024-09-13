import { useState } from 'react';
import { FcGoogle } from 'react-icons/fc';
import { FaGithub } from 'react-icons/fa';
import { TriangleAlert } from 'lucide-react';
import { useRouter } from 'next/navigation';

import { SignInFlow } from '../types';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@repo/ui/components/ui/card';
import { Input } from '@repo/ui/components/ui/input';
import { Button } from '@repo/ui/components/ui/button';
import { Separator } from '@repo/ui/components/ui/separator';
import { LoginResponse } from '@/types/user';

interface SignInCardProps {
  setState: (state: SignInFlow) => void;
}

export const SignInCard = ({ setState }: SignInCardProps) => {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [pending, setPending] = useState(false);

  const signIn = async (
    provider: 'password' | 'github' | 'google',
    credentials?: { email: string; password: string; flow: SignInFlow }
  ) => {
    setPending(true);
    setError('');
    try {
      let response;
      if (provider === 'password') {
        response = await fetch('/api/auth/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: credentials!.email,
            password: credentials!.password,
          }),
        });
      } else {
        // Handle OAuth providers (this is a placeholder, implement actual OAuth flow)
        response = await fetch(`/api/auth/${provider}`, {
          method: 'POST',
        });
      }
      const data: LoginResponse = await response.json();
      if (!response.ok) {
        throw new Error(data.message || 'An error occurred during login');
      }
      // Store user data in localStorage
      localStorage.setItem('user', JSON.stringify(data.data.user));
      router.push('/');
    } catch (err) {
      setError(
        err instanceof Error ? err.message : 'An error occurred during login'
      );
      throw err;
    } finally {
      setPending(false);
    }
  };

  const onPasswordSignIn = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    signIn('password', { email, password, flow: 'signIn' }).catch(() => {
      setError('Invalid email or password');
    });
  };

  const onProviderSignIn = (value: 'github' | 'google') => {
    signIn(value).catch((err) => {
      setError(`Failed to sign in with ${value}: ${err.message}`);
    });
  };

  return (
    <Card className="w-full h-full p-8">
      <CardHeader className="px-0 pt-0">
        <CardTitle>Login to continue</CardTitle>
        <CardDescription>
          Use your email or other service to continue
        </CardDescription>
      </CardHeader>
      {!!error && (
        <div className="bg-destructive/15 p-3 rounded-md flex items-center gap-x-2 text-sm text-destructive mb-6">
          <TriangleAlert className="size-4" />
          <p>{error}</p>
        </div>
      )}
      <CardContent className="space-y-5 px-0 pb-0">
        <form onSubmit={onPasswordSignIn} className="space-y-2.5">
          <Input
            disabled={pending}
            value={email}
            onChange={({ target }) => setEmail(target.value)}
            placeholder="Email"
            type="email"
            required
          />
          <Input
            disabled={pending}
            value={password}
            onChange={({ target }) => setPassword(target.value)}
            placeholder="Password"
            type="password"
            required
          />
          <Button
            type="submit"
            className="w-full"
            size={'lg'}
            disabled={pending}
          >
            Continue
          </Button>
        </form>
        <Separator />
        <div className="flex flex-col gap-y-2.5">
          <Button
            type="button"
            disabled={pending}
            onClick={() => onProviderSignIn('google')}
            variant={'outline'}
            size={'lg'}
            className="w-full relative"
          >
            <FcGoogle className="size-5 absolute top-3 left-2.5" />
            Continue with Google
          </Button>
          <Button
            type="button"
            disabled={pending}
            onClick={() => onProviderSignIn('github')}
            variant={'outline'}
            size={'lg'}
            className="w-full relative"
          >
            <FaGithub className="size-5 absolute top-3 left-2.5" />
            Continue with Github
          </Button>
        </div>
        <p className="text-xs text-muted-foreground">
          Don&apos;t have an account?{' '}
          <Button
            type="button"
            variant={'link'}
            onClick={() => setState('signUp')}
            className="text-sky-700 hover:underline cursor-pointer p-0"
          >
            Sign up
          </Button>
        </p>
      </CardContent>
    </Card>
  );
};
