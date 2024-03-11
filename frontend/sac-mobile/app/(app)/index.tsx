import React from 'react';
import { Text, View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';

const Home = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={logout} size={'sm'}>
                Logout
            </Button>
            <Text>Home</Text>
        </View>
    );
};

export default Home;
