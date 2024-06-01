"use client";

import { UserState, setUser } from "../store/slices/userSlice";
import { useAppDispatch, useAppSelector } from "../store/store";

export default function Home() {
  const user: UserState = {
    clubId: "123",
    userId: "123",
    email: "",
    clubName: "",
    clubLogo: "",
    accessToken: "",
  };
  const { clubId } = useAppSelector((state) => state.user)
  const dispatch = useAppDispatch();

  return (
    <main>
      <button onClick={() => dispatch(setUser(user))}>Press Me</button>
      SAC Dashboard
      {clubId}
    </main>
  );
}
