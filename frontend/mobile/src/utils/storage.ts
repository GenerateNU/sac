import * as SecureStore from 'expo-secure-store';

/**
 * METHODS FOR INTERACTING WITH SECURE STORAGE
 */

async function getToken(key: string) {
    try {
        return SecureStore.getItemAsync(key);
    } catch (error) {
        return null;
    }
}

async function saveToken(key: string, value: string) {
    try {
        return SecureStore.setItemAsync(key, value);
    } catch (error) {
        return null;
    }
}

async function deleteToken(key: string) {
    try {
        return SecureStore.deleteItemAsync(key);
    } catch (error) {
        return null;
    }
}

export { getToken, saveToken, deleteToken }