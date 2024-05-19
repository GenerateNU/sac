import "./App.css";
import { CreateUserRequestBody, authApi, userApi } from "@generatesac/lib";
import { useForm } from "react-hook-form";
import { useAppDispatch, useAppSelector } from "./store/store";
import { useEffect } from "react";

function App() {
	const { token, error } = useAppSelector((state) => state.auth);
	const { user, error: userError } = useAppSelector((state) => state.user);
	const dispatch = useAppDispatch();
	const {
		register,
		handleSubmit,
		formState: { errors }
	} = useForm<CreateUserRequestBody>();

	const submitForm = (data: CreateUserRequestBody) => {
		// transform email string to lowercase to avoid case sensitivity issues in login
		data.email = data.email.toLowerCase();
		dispatch(authApi.endpoints.login.initiate(data));
	};

	const sendSimpleRequest = () => {
		console.log("[send simple request] user.id", user?.id || null);
		dispatch(userApi.endpoints.users.initiate({ page: 1, limit: 10 }));
	};

	useEffect(() => {
		console.log("user", user);
		console.log("token", token);
		console.log("error", error);
	}, [user, token, error]);

	return (
		<>
			<form onSubmit={handleSubmit(submitForm)}>
				<div>
					<label htmlFor="email">Email</label>
					<input
						id="email"
						type="email"
						{...register("email", {
							required: "Email is required"
						})}
					/>
					{errors.email && <p>{errors.email.message}</p>}
				</div>

				<div>
					<label htmlFor="password">Password</label>
					<input
						id="password"
						type="password"
						{...register("password", {
							required: "Password is required",
							minLength: {
								value: 8,
								message:
									"Password must be at least 8 characters long"
							}
						})}
					/>
					{errors.password && <p>{errors.password.message}</p>}
				</div>

				<button type="submit">Login</button>
			</form>

			<div>
				<button onClick={sendSimpleRequest}>Make User Request</button>
			</div>
		</>
	);
}

export default App;
