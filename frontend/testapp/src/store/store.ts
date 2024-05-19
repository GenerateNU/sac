import { authApi, baseApi } from "@generatesac/lib";
import authReducer, {
	setError,
	setIsAuthenticated,
	setToken
} from "@generatesac/lib/src/store/authSlice";
import userReducer from "@generatesac/lib/src/store/userSlice";
import {
	Middleware,
	MiddlewareAPI,
	configureStore,
	isRejectedWithValue
} from "@reduxjs/toolkit";
import { useDispatch, useSelector } from "react-redux";

// log every api call and the api endpoint
const logger: Middleware = (api: MiddlewareAPI) => (next) => (action) => {
	console.log("api", api);
	console.log("action", action);

	if (isRejectedWithValue(action)) {
		console.error("rejectedWithValue", action);
	}

	return next(action);
};

const storeToken = (token: string) => {
	localStorage.setItem("access_token", token);
};

const authMiddleware: Middleware<{}, RootState> =
	(store) => (next) => (action) => {
		if (authApi.endpoints.login.matchFulfilled(action)) {
			const token = authApi.endpoints.login.select(action.requestId)(
				store.getState()
			).token;
			storeToken(token);
		}

		if (authApi.endpoints.register.matchFulfilled(action)) {
			const token = authApi.endpoints.register.select(action.requestId)(
				store.getState()
			).token;
			storeToken(token);
		}

		if (authApi.endpoints.logout.matchFulfilled(action)) {
			localStorage.removeItem("access_token");
		}

		return next(action);
	};

export const store = configureStore({
	reducer: {
		[baseApi.reducerPath]: baseApi.reducer,
		auth: authReducer,
		user: userReducer
	},

	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(logger, baseApi.middleware)
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = useDispatch.withTypes<AppDispatch>();
export const useAppSelector = useSelector.withTypes<RootState>();
