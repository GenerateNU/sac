import { clubsApi } from "@/api/club";
import { eventsApi } from "@/api/event";
import { usersApi } from "@/api/user";
import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";

export const store = configureStore({
  reducer: {
    [usersApi.reducerPath]: usersApi.reducer,
    [clubsApi.reducerPath]: clubsApi.reducer,
    [eventsApi.reducerPath]: eventsApi.reducer,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware()
      .concat(usersApi.middleware)
      .concat(clubsApi.middleware)
      .concat(eventsApi.middleware),
});

// setupListeners(store.dispatch);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
