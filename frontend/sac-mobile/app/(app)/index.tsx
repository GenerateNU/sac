import { useAuthStore } from '@/hooks/use-auth';
import React from 'react';
import { Button, Text, View } from 'react-native';

const Home = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={logout} title="Logout" />
            <Text>Home</Text>
        </View>
    );
};

export default Home;
