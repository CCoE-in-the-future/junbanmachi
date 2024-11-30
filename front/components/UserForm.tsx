import { useState } from "react";

type Props = {
  onSubmit: (name: string, numberPeople: number) => void;
};

export default function UserForm({ onSubmit }: Props) {
  const [name, setName] = useState<string>("");
  const [numberPeople, setNumberPeople] = useState<number | string>("");

  const handleSubmit = () => {
    onSubmit(name, Number(numberPeople));
    setName("");
    setNumberPeople("");
  };

  return (
    <div>
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="名前を入力"
        className="w-full p-3 border rounded-lg shadow-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-400"
      />
      <input
        type="number"
        value={numberPeople}
        onChange={(e) => setNumberPeople(Number(e.target.value))}
        placeholder="人数を入力"
        min="1"
        className="w-full p-3 border rounded-lg shadow-sm mt-3 text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-400"
      />
      <button
        onClick={handleSubmit}
        className="w-full mt-3 bg-blue-500 hover:bg-blue-600 text-white font-semibold p-3 rounded-lg shadow-lg transition duration-300"
      >
        順番に登録
      </button>
    </div>
  );
}
