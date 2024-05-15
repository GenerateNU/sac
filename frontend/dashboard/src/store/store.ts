import { baseApi } from "@generatesac/lib";
import { combineReducers, configureStore } from "@reduxjs/toolkit";
import { useDispatch, useSelector, useStore } from "react-redux";
import userReducer from "@/src/store/userSlice";
import storage from "redux-persist/lib/storage";
import persistReducer from "redux-persist/es/persistReducer";
import { thunk } from "redux-thunk";
import {
  FLUSH,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
  REHYDRATE,
  persistStore,
} from "redux-persist";

const rootReducer = combineReducers({
  user: userReducer,
  [baseApi.reducerPath]: baseApi.reducer,
});

type RootReducer = ReturnType<typeof rootReducer>;

const persistanceConfig = {
  key: "root",
  storage,
  blacklist: [baseApi.reducerPath],
};

const persistedReducer = persistReducer<RootReducer>(
  persistanceConfig,
  rootReducer,
);

export const makeStore = () => {
  return configureStore({
    middleware: (getDefaultMiddleware) =>
      getDefaultMiddleware({
        serializableCheck: {
          ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
        },
      })
        .concat(baseApi.middleware)
        .concat(thunk),

    reducer: persistedReducer,
    devTools: process.env.NODE_ENV !== "production",
  });
};

// State persistor:
export const persistor = persistStore(makeStore());

// Redux types:
export type AppStore = ReturnType<typeof makeStore>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];

// Typed Redux interactive methods:
export const useAppDispatch = useDispatch.withTypes<AppDispatch>();
export const useAppSelector = useSelector.withTypes<RootState>();
export const useAppStore = useStore.withTypes<AppStore>();
