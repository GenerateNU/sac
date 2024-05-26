'use client'

import { Provider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import { persistor, store } from "./store";

type StoreProviderProps = {
  children: React.ReactNode;
}

/**
 * The Redux store provider component. This makes the store accessible globally across the repository
 */
export default function StoreProvider({ children }: StoreProviderProps) {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        {children}
      </PersistGate>
    </Provider>
  )
}
