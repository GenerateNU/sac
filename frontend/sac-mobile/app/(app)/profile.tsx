import React from 'react';
import { View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';

const Profile = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={logout}>Logout</Button>
        </View>
    );
};

export default Profile;
