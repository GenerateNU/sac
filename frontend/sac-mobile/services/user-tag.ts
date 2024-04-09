import { Tag } from '@/types/tag';
import { User } from '@/types/user';

import { api } from './api';

/**
 * Create association between user and a tag.
 * @param userId ID of the user.
 * @param tag Tag.
 * @returns User.
 */
export const CreateUserTags = async (
    userId: string,
    tagsId: string[]
): Promise<void> => {
    try {
        const tagsRequestBody = { tags: tagsId };
        await api.post<User>(
            `/users/${userId}/tags`,
            JSON.stringify(tagsRequestBody)
        );
    } catch (error) {
        console.log('Error associating users to a tag');
        throw new Error('Error associating users to a tag');
    }
};
