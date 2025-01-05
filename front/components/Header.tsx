"use client";

import { useRouter } from "next/navigation";

export default function Header() {
  const router = useRouter();

  const handleSignIn = async () => {
    try {
      // バックエンドの /login エンドポイントにリクエストを送信
      window.location.href = "http://localhost:8080/login"; // サーバー側でリダイレクト処理を行う
    } catch (error) {
      console.error("サインイン中にエラーが発生しました:", error);
      alert("サインイン中に問題が発生しました。");
    }

    router.push("/admin");
  };

  return (
    <header className="flex justify-between items-center mb-8">
      <button
        onClick={handleSignIn}
        className={"px-4 py-2 rounded bg-green-500 text-white"}
      >
        管理者サインイン
      </button>
    </header>
  );
}
