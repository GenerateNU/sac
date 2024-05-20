"use client";

import { useEffect } from "react";
import { calendarAuthRedirect } from "@generatesac/lib";
import useGoogle from "../hooks/useGoogle";

export default function Home() {
  const { googleAuthorize } = useGoogle();

  return (
    <div className="flex flex-col items-center justify-between min-h-screen p-24">
      <button onClick={googleAuthorize}>Click Me</button>
    </div>
  );
}
