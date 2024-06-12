import { Tag } from '@generatesac/lib';
import axios from 'axios';

export const URL = 'http://127.0.0.1:8080/api/v1';

export const getTags = async (): Promise<Tag[]> => {
    try {
        const tags = await axios.get(`${URL}/tags`);
        return tags.data;
    } catch (error) {
        console.error('Failed to get tags', error);
        throw error;
    }
};
