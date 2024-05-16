import { baseApi } from '@generatesac/lib';
import { configureStore } from '@reduxjs/toolkit'
import { useDispatch, useSelector, useStore } from 'react-redux';
import { clubReducer } from '../app/state/club/clubSlice';
import { recruitmentReducer } from '../app/state/recruitment/recruitmentSlice';

export const makeStore = () => {
  return configureStore({
    reducer: {
      [baseApi.reducerPath]: baseApi.reducer,
      clubReducer: clubReducer,
      recruitmentReducer: recruitmentReducer,
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
