import { useState, useEffect } from "react";
import {
  getUsers,
  addUser,
  deleteUser,
  updateUser,
  getWaitTime,
} from "@/services/api";
import { User } from "@/types/user";

export function useUsers() {
  const [users, setUsers] = useState<User[]>([]);
  const [waitTime, setWaitTime] = useState<number | null>(null);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    const userData = await getUsers();
    const sortedUsers = userData.sort(
      (a, b) =>
        new Date(a.arrivalTime).getTime() - new Date(b.arrivalTime).getTime()
    );
    setUsers(sortedUsers);

    const timeData = await getWaitTime();
    setWaitTime(timeData.waitTime);
  };

  const handleAddUser = async (name: string, numberPeople: number) => {
    await addUser({ name, numberPeople });
    fetchData();
  };

  const handleDeleteUser = async (id: string) => {
    await deleteUser(id);
    fetchData();
  };

  const handleUpdateUser = async (id: string) => {
    await updateUser(id);
    fetchData();
  };

  return { users, waitTime, handleAddUser, handleDeleteUser, handleUpdateUser };
}
