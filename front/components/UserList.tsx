import { User } from "@/types/user";

type Props = {
  users: User[];
  onDelete: (id: string) => void;
  onUpdate: (id: string) => void;
};

export default function UserList({ users, onDelete, onUpdate }: Props) {
  return (
    <div className="bg-white p-6 rounded-lg shadow-lg">
      <h2 className="text-xl font-semibold text-gray-800 mb-4">現在の待ち順</h2>
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
            <span className="font-medium text-gray-800">{`${user.numberPeople}人`}</span>
            <span className="text-gray-600">
              {new Date(user.arrivalTime).toLocaleTimeString()}
            </span>
            <button
              onClick={() => onUpdate(user.id)}
              className="ml-4 bg-green-500 text-white p-2 rounded-lg"
            >
              入店
            </button>
            <button
              onClick={() => onDelete(user.id)}
              className="ml-4 bg-red-500 text-white p-2 rounded-lg"
            >
              削除
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
