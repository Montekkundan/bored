'use client';

import { UserButton } from '@/features/auth/components/user-button';
import { useCreateBoringspaceModal } from '@/features/boringspaces/store/use-create-boringspace-modal';
import { useGetBoringspaces } from '@/hooks/use-get-boringspaces';
import { useRouter } from 'next/navigation';
import { useEffect, useMemo } from 'react';

export default function Home() {
  const router = useRouter();
  const [open, setOpen] = useCreateBoringspaceModal();
  const { data, isLoading } = useGetBoringspaces();

  const boringspaceId = useMemo(() => data?.[0]?.boringspace_id, [data]);

  useEffect(() => {
    if (isLoading) return;
    if (boringspaceId) {
      router.replace(`/boringspace/${boringspaceId}`);
    } else if (!open) {
      setOpen(true);
    }
  }, [boringspaceId, isLoading, open, setOpen, router]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <div className="absolute top-4 right-4">
        <UserButton />
      </div>
      {!boringspaceId && !open && (
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4">Welcome to Boring Spaces</h1>
          <p className="mb-4">
            No boring spaces found. Create one to get started!
          </p>
          <button
            onClick={() => setOpen(true)}
            className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
          >
            Create Boring Space
          </button>
        </div>
      )}
    </div>
  );
}
