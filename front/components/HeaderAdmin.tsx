"use client";

export default function HeaderAdmin() {
  const handleSignOut = () => {
    try {
      // バックエンドの /login エンドポイントにリクエストを送信
      window.location.href = "http://localhost:8080/api/logout"; // サーバー側でリダイレクト処理を行う
    } catch (error) {
      console.error("サインアウト中にエラーが発生しました:", error);
      alert("サインアウト中に問題が発生しました。");
    }
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
