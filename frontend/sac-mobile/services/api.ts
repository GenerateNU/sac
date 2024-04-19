import axios, { InternalAxiosRequestConfig } from 'axios';

import { API_BASE_URL } from '@/lib/const';
import { tokenCache } from '@/lib/utils';

// TODO: Add a refresh token interceptor.
// TODO: Add a retry interceptor.

const api = axios.create({
    baseURL: API_BASE_URL,
    withCredentials: true,
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json'
    }
});

const addAuthorizationHeaderInterceptor = async (
    config: InternalAxiosRequestConfig
) => {
    if (
        (config.url?.includes('/auth/login') && config.method === 'POST') ||
        (config.url?.includes('/users/') && config.method === 'POST')
    ) {
        return config;
    }

    const refreshToken = await tokenCache.getToken('refreshToken');
    if (!refreshToken) {
        console.log('No refresh token');
        return config;
    }

    const accessToken = await tokenCache.getToken('accessToken');
    if (!accessToken) {
        console.log('No access token');
        return config;
    }

    const currentTime = new Date().toLocaleTimeString();
    console.log(
        `[${currentTime}] ${config.url} ${config.method?.toUpperCase()} ${JSON.stringify(config.data || {})}`
    );

    config.headers.Authorization = `Bearer ${accessToken}`;
    return config;
};

// export const refreshTokenInterceptor = async (config: InternalAxiosRequestConfig) => {
//     const accessToken = await tokenCache.getToken('accessToken');
//     if (!accessToken) return config;

//     const bufferTime = experationTime(accessToken, 30000);

//     if (bufferTime <= 0) {
//         try {
//             const refreshToken = await tokenCache.getToken('refreshToken');
//             if (!refreshToken) throw new Error('No refresh token');

//             const newTokens = await refreshAccessToken(refreshToken);

//             config.headers.Authorization = `Bearer ${newTokens.accessToken}`;
//             return config;
//         } catch (error) {
//             console.log('Error refreshing access token:', error);
//             // signOut();
//             return config;
//         }
//     }

//     config.headers.Authorization = `Bearer ${accessToken}`;
//     return config;
// }

api.interceptors.request.use(addAuthorizationHeaderInterceptor, (error) =>
    Promise.reject(error)
);
// api.interceptors.request.use(refreshTokenInterceptor, (error) => Promise.reject(error));

// export const unauthorizedInterceptor = async (error: AxiosError) => {
//     if (error.response?.status !== 401) return Promise.reject(error);

//     const refreshToken = await tokenCache.getToken('refreshToken');
//     if (!refreshToken) return Promise.reject(error);

//     try {
//         const newTokens = await refreshAccessToken(refreshToken);
//         const config = error.config as InternalAxiosRequestConfig;
//         config.headers.Authorization = `Bearer ${newTokens.accessToken}`;
//         return axios.request(config);
//     } catch (error) {
//         return Promise.reject(error);
//     }
// }

// axios.interceptors.response.use(undefined, async (error: AxiosError) => {
//     if (error.response?.status !== 401) return Promise.reject(error);

//     const refreshToken = await tokenCache.getToken('refreshToken');
//     if (!refreshToken) return Promise.reject(error);

//     try {
//         const newTokens = await refreshAccessToken(refreshToken);
//         const config = error.config as InternalAxiosRequestConfig;
//         config.headers.Authorization = `Bearer ${newTokens.accessToken}`;
//         return axios.request(config);
//     } catch (error) {
//         return Promise.reject(error);
//     }
// });

export { api, addAuthorizationHeaderInterceptor };
