import React from 'react';
import { Text, View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import { Link } from 'expo-router';

const Home = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={logout}>Logout</Button>
            <Text>Home</Text>
            <Link href={`/club/${1}`}>Club</Link>
            <Link href={`/club/modal/${1}`}>Modal</Link>

        </View>
    );
};

export default Home;
