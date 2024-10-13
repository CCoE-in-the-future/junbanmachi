"use client";

import { useState, useEffect } from "react";

type User = {
  id: number;
  name: string;
  arrivalTime: Date;
};

export default function Home() {
  const [users, setUsers] = useState<User[]>([]);
  const [newUserName, setNewUserName] = useState("");
  const [waitTime, setWaitTime] = useState<number | null>(null);

  useEffect(() => {
    fetchUsers();
    fetchWaitTime();
  }, []);

  const fetchUsers = async () => {
    const res = await fetch("/api/users");
    const data = await res.json();
    setUsers(data);
  };

  const fetchWaitTime = async () => {
    const res = await fetch("/api/wait-time");
    const data = await res.json();
    setWaitTime(data.waitTime);
  };

  const addUser = async () => {
    const res = await fetch("/api/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name: newUserName }),
    });
    if (res.ok) {
      setNewUserName("");
      fetchUsers();
    }
  };

  const deleteUser = async (id: number) => {
    const res = await fetch("/api/users", {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id }),
    });
    if (res.ok) {
      const data = await res.json();
      setUsers(data);
    }
  };

  return (
    <div className="container mx-auto p-6 max-w-lg">
      <h1 className="text-3xl font-bold text-center mb-8 text-blue-600">
        順番待ちシステム
      </h1>

      <div className="mb-6">
        <input
          type="text"
          value={newUserName}
          onChange={(e) => setNewUserName(e.target.value)}
          placeholder="名前を入力"
          className="w-full p-3 border rounded-lg shadow-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-400"
        />
        <button
          onClick={addUser}
          className="w-full mt-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold p-3 rounded-lg shadow-lg transition duration-300"
        >
          順番に登録
        </button>
      </div>

      <div className="mb-6 text-center">
        <h2 className="text-xl font-medium text-gray-500">
          待ち時間予測:{" "}
          <span className="font-bold text-blue-600">
            {waitTime !== null ? `${waitTime} 分` : "取得中..."}
          </span>
        </h2>
      </div>

      <div className="bg-white p-6 rounded-lg shadow-lg">
        <h2 className="text-xl font-semibold text-gray-800 mb-4">
          現在の待ち順
        </h2>
        <ul className="space-y-2">
          {users.map((user) => (
            <li
              key={user.id}
              className="flex justify-between items-center p-3 border rounded-lg shadow-sm"
            >
              <span className="font-medium text-gray-800">{user.name}</span>
              <span className="text-gray-600">
                {new Date(user.arrivalTime).toLocaleTimeString()}
              </span>
              <button
                onClick={() => deleteUser(user.id)}
                className="ml-4 bg-red-500 text-white p-2 rounded-lg"
              >
                削除
              </button>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
