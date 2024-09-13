import { useState } from 'react';
import { FcGoogle } from 'react-icons/fc';
import { FaGithub } from 'react-icons/fa';
import { TriangleAlert } from 'lucide-react';

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

interface SignUpCardProps {
  setState: (state: SignInFlow) => void;
}

export const SignUpCard = ({ setState }: SignUpCardProps) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [phoneNumber, setPhoneNumber] = useState('');
  const [error, setError] = useState('');
  const [pending, setPending] = useState(false);

  const onPasswordSignUp = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (password !== confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    if (password.length < 8) {
      setError('Password must be at least 8 characters long');
      return;
    }

    setError('');
    setPending(true);

    try {
      const response = await fetch('/api/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          username: name,
          email,
          password,
          phone_number: phoneNumber,
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.message || 'Registration failed');
      }

      setState('signIn');
    } catch (error) {
      setError(
        error instanceof Error
          ? error.message
          : 'An error occurred during registration'
      );
    } finally {
      setPending(false);
    }
  };

  const onProviderSignUp = (value: 'github' | 'google') => {
    setPending(true);
    setPending(false);
  };

  return (
    <Card className="w-full h-full p-8">
      <CardHeader className="px-0 pt-0">
        <CardTitle>Sign up to continue</CardTitle>
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
      <CardContent className="space-y-4 px-0 pb-0">
        <form onSubmit={onPasswordSignUp} className="space-y-2.5">
          <Input
            disabled={pending}
            value={name}
            onChange={({ target }) => setName(target.value)}
            placeholder="Full name"
            required
            autoComplete="name"
          />
          <Input
            disabled={pending}
            value={email}
            onChange={({ target }) => setEmail(target.value)}
            placeholder="Email"
            type="email"
            required
            autoComplete="email"
          />
          <Input
            disabled={pending}
            value={phoneNumber}
            onChange={({ target }) => setPhoneNumber(target.value)}
            placeholder="Phone number"
            type="tel"
            required
            autoComplete="tel"
          />
          <Input
            disabled={pending}
            value={password}
            onChange={({ target }) => setPassword(target.value)}
            placeholder="Password"
            type="password"
            required
            autoComplete="new-password"
          />
          <Input
            disabled={pending}
            value={confirmPassword}
            onChange={({ target }) => setConfirmPassword(target.value)}
            placeholder="Confirm password"
            type="password"
            required
            autoComplete="new-password"
          />
          <Button
            type="submit"
            className="w-full"
            size={'lg'}
            disabled={pending}
          >
            {pending ? 'Signing up...' : 'Continue'}
          </Button>
        </form>
        <Separator />
        <div className="flex flex-col gap-y-2.5">
          <Button
            type="button"
            disabled={pending}
            onClick={() => onProviderSignUp('google')}
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
            onClick={() => onProviderSignUp('github')}
            variant={'outline'}
            size={'lg'}
            className="w-full relative"
          >
            <FaGithub className="size-5 absolute top-3 left-2.5" />
            Continue with Github
          </Button>
        </div>
        <p className="text-xs text-muted-foreground">
          Already have an account?{' '}
          <Button
            type="button"
            variant={'link'}
            onClick={() => setState('signIn')}
            className="text-sky-700 hover:underline cursor-pointer p-0"
          >
            Sign in
          </Button>
        </p>
      </CardContent>
    </Card>
  );
};
