'use client';
import { TriangleAlert } from 'lucide-react';

const BoringspaceIdPage = () => {
  return (
    <div className="h-full flex-1 flex items-center justify-center flex-col gap-2">
      <TriangleAlert className="size-6 text-destructive" />
      <span className="text-sm text-muted-foreground">No channel found</span>
    </div>
  );
};

export default BoringspaceIdPage;
