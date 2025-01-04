"use client";

import HeaderAdmin from "@/components/HeaderAdmin";
import UserForm from "@/components/UserForm";
import UserListAdmin from "@/components/UserListAdmin";
import WaitTimeDisplay from "@/components/WaitTimeDisplay";
import { useUsers } from "@/hooks/useUsers";

export default function Home() {
  const { users, waitTime, handleAddUser, handleDeleteUser, handleUpdateUser } =
    useUsers();

  return (
    <div className="container mx-auto p-6 max-w-lg">
      <HeaderAdmin />
      <main>
        <h1 className="text-3xl font-bold text-center mb-8 text-blue-600">
          順番待ちシステム 管理者向け
        </h1>
        <UserForm onSubmit={handleAddUser} />
        <WaitTimeDisplay waitTime={waitTime} />
        <UserListAdmin
          users={users}
          onDelete={handleDeleteUser}
          onUpdate={handleUpdateUser}
        />
      </main>
    </div>
  );
}
