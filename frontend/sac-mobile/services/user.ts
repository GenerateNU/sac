import { User } from "@/types/user";
import { AxiosError } from "axios";
import { api } from "./api";
import { ErrorResponse } from "@/types/error";

/**
 * Gets the user from the access token.
 * @param accessToken The access token of the user.
 * @returns The user that was retrieved.
 */
export const getCurrentUser = async (): Promise<User> => {
    try {
        const response = await api.get('/users/me');

        return response.data as User;
    } catch (error: any) {
        const axiosErr: AxiosError = error;
        const response = axiosErr.response;
        if (response) {
            const data = response.data as ErrorResponse;
            throw new Error(data.error);
        } else {
            throw new Error('Error getting current user');
        }
    }
};