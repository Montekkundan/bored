import useSWR from 'swr';

const fetcher = async (url: string) => {
  const response = await fetch(url, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem('token')}`,
    },
  });
  if (!response.ok) {
    throw new Error("Failed to fetch user's boring spaces");
  }
  return response.json();
};

export const useGetBoringspaces = () => {
  const { data, error, isLoading } = useSWR('/api/users/boringspaces', fetcher);

  return {
    data: data?.data,
    isLoading,
    isError: error,
  };
};
