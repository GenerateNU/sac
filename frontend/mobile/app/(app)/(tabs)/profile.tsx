import React from 'react';
import { SafeAreaView, Text, View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';

const Profile = () => {
    const { onLogout, user } = useAuthStore();

    const handleSignOut = async () => {
        onLogout();
    };

    return (
        <SafeAreaView className="justify-center flex-1">
            <View className="px-[8%] gap-5 flex-col items-center">
                <Text className="text-lg">Welcome {user?.first_name}</Text>
                <Button onPress={handleSignOut}>Sign Out</Button>
            </View>
        </SafeAreaView>
    );
};

export default Profile;
