"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import HeaderAdmin from "@/components/HeaderAdmin";
import UserForm from "@/components/UserForm";
import UserListAdmin from "@/components/UserListAdmin";
import WaitTimeDisplay from "@/components/WaitTimeDisplay";
import { useUsers } from "@/hooks/useUsers";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export default function Home() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const router = useRouter();
  const { users, waitTime, handleAddUser, handleDeleteUser, handleUpdateUser } =
    useUsers();

  useEffect(() => {
    const checkAuthState = async () => {
      try {
        const response = await fetch(`${API_BASE_URL}/auth-status`, {
          method: "GET",
          credentials: "include",
        });

        if (!response.ok) {
          throw new Error("Failed to authenticate");
        }

        const data = await response.json();

        if (data.status === "authenticated") {
          setIsAuthenticated(true);
        } else {
          throw new Error("Not authenticated");
        }
      } catch (error) {
        console.error("Authentication error:", error);
        router.push("/");
      }
    };

    checkAuthState();
  }, [router]);

  if (!isAuthenticated) {
    return (
      <div className="flex justify-center items-center h-screen">
        <p className="text-xl font-bold text-gray-700">読み込み中...</p>
      </div>
    );
  }

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
