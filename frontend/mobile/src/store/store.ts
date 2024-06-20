import { useDispatch, useSelector, useStore } from 'react-redux';

import { baseApi } from '@generatesac/lib';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { combineReducers, configureStore } from '@reduxjs/toolkit';
import {
    FLUSH,
    PAUSE,
    PERSIST,
    PURGE,
    REGISTER,
    REHYDRATE,
    persistStore
} from 'redux-persist';
import persistReducer from 'redux-persist/es/persistReducer';
import autoMergeLevel2 from 'redux-persist/es/stateReconciler/autoMergeLevel2';
import { thunk } from 'redux-thunk';

import clubReducer from '@/src/store/slices/clubSlice';
import eventReducer from '@/src/store/slices/eventSlice';
import userReducer from '@/src/store/slices/userSlice';
import globalReducer from './slices/globalSlice';

const rootReducer = combineReducers({
    global: globalReducer,
    user: userReducer,
    event: eventReducer,
    club: clubReducer,
    [baseApi.reducerPath]: baseApi.reducer
});

type RootReducer = ReturnType<typeof rootReducer>;

const persistanceConfig = {
    key: 'root',
    storage: AsyncStorage,
    stateReconciler: autoMergeLevel2,
    blacklist: [baseApi.reducerPath],
    whitelist: ['user', 'event', 'club', 'global']
};

const persistedReducer = persistReducer<RootReducer>(
    persistanceConfig,
    rootReducer
);

export const store = configureStore({
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({
            serializableCheck: {
                ignoredActions: [
                    FLUSH,
                    REHYDRATE,
                    PAUSE,
                    PERSIST,
                    PURGE,
                    REGISTER
                ]
            }
        })
            .concat(baseApi.middleware)
            .concat(thunk),

    reducer: persistedReducer,
    devTools: process.env.NODE_ENV !== 'production'
});

export const persistor = persistStore(store);

// Redux types:
export type AppStore = typeof store;
export type RootState = ReturnType<AppStore['getState']>;
export type AppDispatch = AppStore['dispatch'];

// Typed Redux interactive methods:
export const useAppDispatch = useDispatch.withTypes<AppDispatch>();
export const useAppSelector = useSelector.withTypes<RootState>();
export const useAppStore = useStore.withTypes<AppStore>();
