import { baseApi } from '@generatesac/lib';
import { UnknownAction, configureStore } from '@reduxjs/toolkit'
import { useDispatch, useSelector, useStore } from 'react-redux';
import { clubReducer } from '@/src/app/reducers/clubReducer';
import { Reducer } from 'react';
import { ClubState } from '../app/actions/actionTypes';

export const makeStore = () => {
  return configureStore({
    reducer: {
      [baseApi.reducerPath]: baseApi.reducer,
      club: clubReducer as Reducer<ClubState | undefined, UnknownAction>,
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
