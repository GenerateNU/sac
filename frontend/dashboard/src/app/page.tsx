"use client";

import useGoogle from "../hooks/useGoogle";

export default function Home() {
  const { googleRedirect } = useGoogle();

  return (
    <main className="flex flex-col items-center justify-between min-h-screen p-24">
      <button onClick={() => googleRedirect()}>Click Me</button>
    </main>
  );
}
