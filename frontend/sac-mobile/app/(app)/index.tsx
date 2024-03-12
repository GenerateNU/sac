import React from 'react';
import { Text, View} from 'react-native';

import { useAuthStore } from '@/hooks/use-auth';

const Home = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Text onPress={logout}>Logout</Text>
            <Text>Home</Text>
        </View>
    );
};

export default Home;
