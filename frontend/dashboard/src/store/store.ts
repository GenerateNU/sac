import { configureStore } from '@reduxjs/toolkit'
import { useDispatch, useSelector, useStore } from 'react-redux';
import { baseApi } from '../api/api';

export const makeStore = () => {
  return configureStore({
    reducer: {
      [baseApi.reducerPath]: baseApi.reducer,
    },

    middleware: (getDefaultMiddleware) =>
      getDefaultMiddleware()
      .concat(baseApi.middleware)
  })
};

// Redux types:
export type AppStore = ReturnType<typeof makeStore>
export type RootState = ReturnType<AppStore['getState']>
export type AppDispatch = AppStore['dispatch']

// Typed Redux interactive methods:
export const useAppDispatch = useDispatch.withTypes<AppDispatch>()
export const useAppSelector = useSelector.withTypes<RootState>()
export const useAppStore = useStore.withTypes<AppStore>()