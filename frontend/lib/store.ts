import { clubsApi } from "@/api/club";
import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";

export const store = configureStore({
  reducer: {
    [clubsApi.reducerPath]: clubsApi.reducer,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(clubsApi.middleware),
});

setupListeners(store.dispatch);
