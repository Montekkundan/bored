import { UserResponse } from '@/types/user';
import useSWR from 'swr';

const fetcher = async (url: string) => {
  const res = await fetch(url);
  if (!res.ok) {
    throw new Error('Failed to fetch user data');
  }
  return res.json();
};

export function useCurrentUser() {
  const { data, error, isLoading, mutate } = useSWR<UserResponse>(
    '/api/auth/user',
    fetcher
  );

  return {
    data,
    isLoading,
    isError: error,
    mutate,
  };
}
