import { View, Text } from 'react-native';
import React from 'react';
import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';

const Profile = () => {
    const { signOut } = useAuthStore();

    const handleSignOut = async () => {
        signOut();
    };

    return (
        <View className='items-center justify-center flex-1 gap-5'>
            <Button onPress={handleSignOut}>Sign Out</Button>
        </View>
    );
}

export default Profile;