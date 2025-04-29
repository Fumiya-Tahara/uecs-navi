"use client";

import { M304IDsProvider } from "@/hooks/m304ids-context";

export function Providers({ children }: { children: React.ReactNode }) {
  return <M304IDsProvider>{children}</M304IDsProvider>;
}
