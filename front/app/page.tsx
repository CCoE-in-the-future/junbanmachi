"use client";

import Header from "@/components/Header";
import UserList from "@/components/UserList";
import WaitTimeDisplay from "@/components/WaitTimeDisplay";
import { useUsers } from "@/hooks/useUsers";

export default function Home() {
  const { users, waitTime } = useUsers();

  return (
    <div className="container mx-auto p-6 max-w-lg">
      <Header />
      <main>
        <h1 className="text-3xl font-bold text-center mb-8 text-blue-600">
          順番待ちシステム 利用者向け
        </h1>
        <WaitTimeDisplay waitTime={waitTime} />
        <UserList users={users} />
      </main>
    </div>
  );
}
