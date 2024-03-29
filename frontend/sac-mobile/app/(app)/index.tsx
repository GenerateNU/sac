import React from 'react';
import { Text, View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';

const Home = () => {
    const { signOut, user } = useAuthStore();

    const handleSignOut = async () => {
        signOut();
    }

    return (
        <View className="items-center justify-center flex-1 gap-5">
            <Text>Welcome {user?.first_name}</Text>
            <Button onPress={handleSignOut}>Sign Out</Button>
        </View>
    );
};

export default Home;
