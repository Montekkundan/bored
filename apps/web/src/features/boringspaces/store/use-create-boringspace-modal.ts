import { atom, useAtom } from 'jotai';

const modalState = atom(false);

export const useCreateBoringspaceModal = () => {
  return useAtom(modalState);
};
