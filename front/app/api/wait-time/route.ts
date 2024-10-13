import { NextResponse } from "next/server";
import Users from "@/app/lib/userSample";

export async function GET() {
  const userLength = Users.length;
  const estimatedWaitTime = userLength * 20;
  return NextResponse.json({ waitTime: estimatedWaitTime });
}
