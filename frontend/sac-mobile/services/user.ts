import { AxiosError } from 'axios';

import { api } from '@/services/api';
import { Club } from '@/types/club';
import { ErrorResponse } from '@/types/error';
import { User } from '@/types/user';
import { uuid } from '@/types/uuid';

/**
 * Gets the user from the access token.
 *
 * @returns The user that was retrieved.
 * @throws Error if the user cannot be retrieved.
 * @see User
 */
const getCurrentUser = async (): Promise<User> => {
    return api
        .get<User>('/users/me')
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error getting current user');
        });
};

/**
 * Associates a user with a tag.
 *
 * @param userID The ID of the user.
 * @param tagIDs The IDs of the tags.
 * @throws Error if the user cannot be associated with the tag.
 * @see User
 * @see Tag
 */
const createUserTags = async (
    userID: uuid,
    tagIDs: uuid[]
): Promise<void> => {
    return api
        .post<void>(`/users/${userID}/tags`, { tags: tagIDs })
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error associating users to a tag');
        });
}

/**
 * Updates the user's information.
 *
 * @param userID The ID of the user.
 * @param updateUserBody The updated user information.
 * @returns The updated user.
 * @throws Error if the user cannot be updated.
 * @see User
 */
const updateUser = async (
    userID: uuid,
    updateUserBody: {
        graduation_year: number;
        college: string;
        graduation_cycle: string;
    }
): Promise<User> => {
    return api
        .patch<User>(`/users/${userID}`, updateUserBody)
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error updating user');
        });
};

/**
 * Fetches the clubs that a user is following.
 *
 * @param userID The ID of the user.
 * @returns A Promise that resolves to the clubs that the user is following or rejects with an error.
 * @throws Error if the clubs cannot be fetched.
 * @see Club
 */
const fetchUserFollowing = async (userID: uuid): Promise<Club[]> => {
    return api
        .get<Club[]>(`/users/${userID}/follower`)
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error fetching followed clubs');
        });
};

/**
 * Creates a user following a club.
 *
 * @param userID the ID of the user
 * @param clubID the ID of the club
 * @returns A Promise that resolves to void or rejects with an error
 * @throws Error if the user cannot follow the club
 * @see Club
 * @see User
 */
const createUserFollowing = async (
    userID: uuid,
    clubID: uuid
): Promise<void> => {
    return api
        .post<void>(`/users/${userID}/follower/${clubID}`)
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error following clubs');
        })
}

/**
 * Deletes a user following a club.
 *
 * @param userID the ID of the user
 * @param clubID the ID of the club
 * @returns A Promise that resolves to void or rejects with an error
 * @throws Error if the user cannot unfollow the club
 * @see Club
 * @see User
 */
const deleteUserFollowing = async (
    userID: uuid,
    clubID: uuid
): Promise<void> => {
    return api
        .delete<void>(`/users/${userID}/follower/${clubID}`)
        .then((response) => response.data)
        .catch((error: AxiosError<ErrorResponse>) => {
            console.error(error);
            throw new Error('Error unfollowing clubs');
        });
};

export {
    getCurrentUser,
    createUserTags,
    updateUser,
    fetchUserFollowing,
    createUserFollowing,
    deleteUserFollowing
};
