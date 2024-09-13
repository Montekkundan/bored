import { toast } from 'sonner';
import { useState } from 'react';
import { useRouter } from 'next/navigation';

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@repo/ui/components/ui/dialog';
import { Input } from '@repo/ui/components/ui/input';
import { Button } from '@repo/ui/components/ui/button';
import { useCreateBoringspaceModal } from '../store/use-create-boringspace-modal';
import { Textarea } from '@repo/ui/components/ui/textarea';

export const CreateBoringspaceModal = () => {
  const router = useRouter();
  const [open, setOpen] = useCreateBoringspaceModal();
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [isPending, setIsPending] = useState(false);

  const handleClose = () => {
    setOpen(false);
    setName('');
    setDescription('');
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsPending(true);

    try {
      const response = await fetch('/api/boringspaces', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify({ name, description }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.message || 'Failed to create boring space');
      }

      toast.success('Boringspace created');
      router.push(`/boringspace/${data.data.id}`);
      handleClose();
    } catch (error) {
      toast.error(error instanceof Error ? error.message : 'An error occurred');
    } finally {
      setIsPending(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add a boringspace</DialogTitle>
          <DialogDescription className="hidden"></DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input
            value={name}
            onChange={(e) => setName(e.target.value)}
            disabled={isPending}
            required
            autoFocus
            minLength={3}
            placeholder="Boringspace name e.g. 'Work', 'Personal', 'Home'"
          />
          <Textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            disabled={isPending}
            required
            placeholder="Description of your boring space"
            rows={3}
          />
          <div className="flex justify-end">
            <Button type="submit" disabled={isPending}>
              {isPending ? 'Creating...' : 'Create'}
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};
