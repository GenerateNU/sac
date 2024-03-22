import React from 'react';
import { Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import Svg, {Path, Circle, Line} from 'react-native-svg';

const Profile = () => {
    const { logout } = useAuthStore();
    return (
        <View className="items-center justify-center flex-1">
            <Button onPress={logout}>Logout</Button>
            <Button onPress={() => router.push('/(app)/')}>Home</Button>
        </View>
    );
};

export default Profile;
