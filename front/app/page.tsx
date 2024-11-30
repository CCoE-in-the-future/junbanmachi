"use client";

import UserForm from "@/components/UserForm";
import UserList from "@/components/UserList";
import WaitTimeDisplay from "@/components/WaitTimeDisplay";
import { useUsers } from "@/hooks/useUsers";

export default function Home() {
  const { users, waitTime, handleAddUser, handleDeleteUser, handleUpdateUser } =
    useUsers();

  return (
    <div className="container mx-auto p-6 max-w-lg">
      <h1 className="text-3xl font-bold text-center mb-8 text-blue-600">
        順番待ちシステム
      </h1>
      <UserForm onSubmit={handleAddUser} />
      <WaitTimeDisplay waitTime={waitTime} />
      <UserList
        users={users}
        onDelete={handleDeleteUser}
        onUpdate={handleUpdateUser}
      />
    </div>
  );
}
