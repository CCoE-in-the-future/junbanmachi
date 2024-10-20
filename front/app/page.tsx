"use client";

import { useState, useEffect } from "react";

type User = {
  id: string;
  name: string;
  numberPeople: number;
  waitStatus: boolean;
  arrivalTime: Date;
};

export default function Home() {
  const [users, setUsers] = useState<User[]>([]);
  const [newUserName, setNewUserName] = useState<string>("");
  const [newUserNumberPeople, setNewUserNumberPeople] = useState<
    number | string
  >("");
  const [waitTime, setWaitTime] = useState<number | null>(null);

  const fetchUsers = async () => {
    const res = await fetch("http://localhost:8080/api/users");
    const data = await res.json();

    const sortedData = data
      .map((user: User) => ({
        ...user,
        arrivalTime: new Date(user.arrivalTime),
      }))
      .sort(
        (a: User, b: User) => a.arrivalTime.getTime() - b.arrivalTime.getTime()
      );

    setUsers(sortedData);
  };

  const fetchWaitTime = async () => {
    const res = await fetch("http://localhost:8080/api/wait-time");
    const data = await res.json();
    setWaitTime(data.waitTime);
  };

  const addUser = async () => {
    const res = await fetch("http://localhost:8080/api/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: newUserName,
        numberPeople: newUserNumberPeople,
      }),
    });
    if (res.ok) {
      setNewUserName("");
      setNewUserNumberPeople("");
      fetchUsers();
      fetchWaitTime();
    }
  };

  const deleteUser = async (id: string) => {
    const res = await fetch("http://localhost:8080/api/users", {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id }),
    });
    if (res.ok) {
      fetchUsers();
      fetchWaitTime();
    }
  };

  const updateUser = async (id: string) => {
    const res = await fetch("http://localhost:8080/api/users", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id }),
    });
    if (res.ok) {
      fetchUsers();
      fetchWaitTime();
    }
  };

  useEffect(() => {
    fetchUsers();
    fetchWaitTime();
  }, []);

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
        <input
          type="number"
          value={newUserNumberPeople}
          onChange={(e) => setNewUserNumberPeople(Number(e.target.value))}
          placeholder="人数を入力"
          min="1"
          className="w-full p-3 border rounded-lg shadow-sm mt-3 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-400"
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
              <span className="font-medium text-gray-800">
                {user.waitStatus ? "待ち" : "入店"}
              </span>
              <span className="font-medium text-gray-800">
                {`${user.numberPeople}人`}
              </span>
              <span className="text-gray-600">
                {new Date(user.arrivalTime).toLocaleTimeString()}
              </span>
              <button
                onClick={() => updateUser(user.id)}
                className="ml-4 bg-green-500 text-white p-2 rounded-lg"
              >
                入店
              </button>
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
