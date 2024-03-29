import { useAuthStore } from "@/hooks/use-auth";
import { API_BASE_URL } from "@/lib/const";
import { experationTime, tokenCache } from "@/lib/utils";
import axios, { InternalAxiosRequestConfig } from "axios";
import { Tokens } from "@/types/user";

export const api = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        "Content-Type": "application/json",
    },
    withCredentials: true,
    timeout: 10000,
});

api.interceptors.request.use(
    async (config) => {
        logRoute(config)

        if (config.url?.includes("/auth/login") && config.method === "POST" ||
            config.url?.includes("/users/") && config.method === "POST") {
            return config;
        }

        const refreshToken = await tokenCache.getToken("refreshToken");
        if (!refreshToken) {
            console.log("No refresh token")
            return config;
        }

        const accessToken = await tokenCache.getToken("accessToken");
        if (!accessToken) {
            console.log("No access token")
            return config;
        }

        logRoute(config, accessToken, refreshToken);

        config.headers.Authorization = `Bearer ${accessToken}`;
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

api.interceptors.request.use(
    async (config) => {
        const store = useAuthStore.getState();
        const accessToken = store.tokens?.accessToken;
        if (!accessToken) {
            return config;
        }

        const bufferTime = experationTime(accessToken, 30000);

        if (bufferTime <= 0) {
            try {
                const refreshToken = store.tokens?.refreshToken;
                if (!refreshToken) {
                    throw new Error("No refresh token");
                }

                const newTokens = await refreshAccessToken(refreshToken);

                store.setTokens({ accessToken: newTokens.accessToken, refreshToken: newTokens.refreshToken });
                config.headers.Authorization = `Bearer ${newTokens.accessToken}`;
                return config;
            } catch (error) {
                // TODO: some sort of alert or notification to the user
                store.signOut();
            }
        }

        config.headers.Authorization = `Bearer ${accessToken}`;
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

const refreshAccessToken = async (refreshToken: string) => {
    try {
        const response = await api.post<Tokens>("/auth/refresh", {
            refresh_token: refreshToken
        });
        return response.data;
    } catch (error) {
        throw new Error("Error refreshing access token");
    }
}


const logRoute = (config: InternalAxiosRequestConfig<any>, accessToken?: string, refreshToken?: string) => {
    if (!accessToken || !refreshToken) {
        const currentTime = Intl.DateTimeFormat("en-US", { timeStyle: "medium" }).format(new Date());
        console.log(`[${currentTime}] ${config.url} ${config.method?.toUpperCase()} ${JSON.stringify(config.data || {})}`);
    }

    const currentTime = Intl.DateTimeFormat("en-US", { timeStyle: "medium" }).format(new Date());
    const accessBufferTime = experationTime(accessToken!, 30000);
    const refreshBufferTime = experationTime(refreshToken!, 30000);
    const accessTimeLeft = Math.floor(accessBufferTime / 1000); // convert to seconds
    const refreshTimeLeft = Math.floor(refreshBufferTime / 1000); // convert to seconds
    const accessMinutes = Math.floor(accessTimeLeft / 60); // convert to minutes
    const accessSeconds = accessTimeLeft % 60; // remaining seconds
    const refreshDays = Math.floor(refreshTimeLeft / 86400); // convert to days
    const refreshHours = Math.floor((refreshTimeLeft % 86400) / 3600); // remaining hours
    const refreshMinutes = Math.floor((refreshTimeLeft % 3600) / 60); // remaining minutes
    const refreshSeconds = refreshTimeLeft % 60; // remaining seconds

    console.log(`[${currentTime}] ${config.url} ${config.method?.toUpperCase()} ${JSON.stringify(config.data || {})}`)
    console.log(`[${currentTime}] Access token time left: ${accessMinutes} minutes ${accessSeconds} seconds Refresh token time left: ${refreshDays} days ${refreshHours} hours ${refreshMinutes} minutes ${refreshSeconds} seconds`);
}
