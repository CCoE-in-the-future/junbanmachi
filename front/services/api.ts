const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function getUsers() {
  const res = await fetch(`${API_BASE_URL}/users`);
  return res.json();
}

export async function addUser(data: { name: string; numberPeople: number }) {
  await fetch(`${API_BASE_URL}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
}

export async function deleteUser(id: string) {
  await fetch(`${API_BASE_URL}/users`, {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ id }),
  });
}

export async function updateUser(id: string) {
  await fetch(`${API_BASE_URL}/users`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ id }),
  });
}

export async function getWaitTime() {
  const res = await fetch(`${API_BASE_URL}/wait-time`);
  return res.json();
}
