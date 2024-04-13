import { clubsApi } from "@/api/club";
import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";

import { eventsApi } from "./api/event";

export const store = configureStore({
  reducer: {
    [clubsApi.reducerPath]: clubsApi.reducer,
    [eventsApi.reducerPath]: eventsApi.reducer,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(clubsApi.middleware, eventsApi.middleware),
});

setupListeners(store.dispatch);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
