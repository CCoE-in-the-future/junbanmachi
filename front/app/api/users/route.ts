import { NextResponse } from "next/server";
import Users from "@/app/lib/userSample";

export async function GET() {
  return NextResponse.json(Users);
}

export async function POST(req: Request) {
  const body = await req.json();
  const { name } = body;

  const newUser = { id: Users.length + 1, name, arrivalTime: new Date() };
  Users.push(newUser);

  return NextResponse.json(newUser, { status: 201 });
}

export async function DELETE(req: Request) {
  const body = await req.json();
  const { id } = body;
  const newUsers = Users.filter((user) => user.id !== id);

  return NextResponse.json(newUsers);
}
