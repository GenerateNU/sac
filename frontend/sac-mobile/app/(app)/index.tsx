import React from 'react';
import { Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';

const Home = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={logout}>Logout</Button>
            <Button onPress={() => router.push('/(app)/homepage')}>Home</Button>
        </View>
    );
};

export default Home;
