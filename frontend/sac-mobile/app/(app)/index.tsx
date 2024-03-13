import React from 'react';
import { Text, View} from 'react-native';

import { useAuthStore } from '@/hooks/use-auth';
import { Button } from '@/components/button';

const Home = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={logout}>Logout</Button>
            <Text>Home</Text>
        </View>
    );
};

export default Home;
