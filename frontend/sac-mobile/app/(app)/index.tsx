import React from 'react';
import { Text, View } from 'react-native';

import { Button } from '@/components/button';
import { useAuth } from '@clerk/clerk-expo';

const Home = () => {
    const { signOut } = useAuth();

    const handleSignOut = async () => {
        await signOut();
    }

    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={handleSignOut}>Sign Out</Button>
            <Text>Home</Text>
        </View>
    );
};

export default Home;
