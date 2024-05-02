'use client'

import { useRef } from "react";
import { AppStore, makeStore } from "./store";
import { Provider } from "react-redux";

type StoreProviderProps = {
  children: React.ReactNode;
}

/**
 * The Redux store provider component. This makes the store accessible globally across the repository
 */
export default function StoreProvider({ children }: StoreProviderProps) {
    const storeRef = useRef<AppStore>();

    // Instantiate store singleton:
    if (!storeRef.current) {
        storeRef.current = makeStore()
      }
    return (
        <Provider store={storeRef.current}>
            {children}
        </Provider>
    )
}
