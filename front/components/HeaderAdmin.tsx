"use client";

import { useRouter } from "next/navigation";

export default function HeaderAdmin() {
  const router = useRouter();

  const handleSignOut = () => {
    alert("サインアウトしました！");
    router.push("/");
  };

  return (
    <header className="flex justify-between items-center mb-8">
      <button
        onClick={handleSignOut}
        className={"px-4 py-2 rounded bg-red-500 text-white"}
      >
        管理者サインアウト
      </button>
    </header>
  );
}
