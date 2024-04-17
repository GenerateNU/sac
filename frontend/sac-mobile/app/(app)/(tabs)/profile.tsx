import React from 'react';
import { SafeAreaView, Text, View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';

const Profile = () => {
    const { signOut, user } = useAuthStore();

    const handleSignOut = async () => {
        signOut();
    };

    return (
        <SafeAreaView className="flex-1 justify-center">
            <View className="px-[8%] gap-5 flex-col items-center">
                <Text className="text-lg">Welcome {user?.first_name}</Text>
                <Button onPress={handleSignOut}>Sign Out</Button>
            </View>
        </SafeAreaView>
    );
};

export default Profile;
