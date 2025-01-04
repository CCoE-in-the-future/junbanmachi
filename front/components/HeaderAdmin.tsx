"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

export default function HeaderAdmin() {
  const [isSignedIn, setIsSignedIn] = useState(false);
  const router = useRouter();

  const handleSignIn = () => {
    setIsSignedIn(true);
    alert("管理者としてサインインしました！");
    router.push("/admin");
  };

  const handleSignOut = () => {
    setIsSignedIn(false);
    alert("サインアウトしました！");
    router.push("/");
  };

  return (
    <header className="flex justify-between items-center mb-8">
      <button
        onClick={isSignedIn ? handleSignOut : handleSignIn}
        className={`px-4 py-2 rounded ${
          isSignedIn ? "bg-red-500 text-white" : "bg-green-500 text-white"
        }`}
      >
        {isSignedIn ? "サインアウト" : "サインイン"}
      </button>
    </header>
  );
}
