import type { Metadata } from "next";
import { Providers } from "@/hooks/providers";

export const metadata: Metadata = {
  title: "UECS Navi",
  description: "",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <head>
        <title>UECS Navi</title>
      </head>
      <body>
        <div id="root">
          <Providers>{children}</Providers>
        </div>
      </body>
    </html>
  );
}
