import React from 'react';
import { View } from 'react-native';

import { Button } from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import { useAuth } from '@/context/AuthContext';

const Profile = () => {
<<<<<<< Updated upstream
    const { signOut } = useAuthStore();
=======
    // const { signOut, user } = useAuthStore();
    const { onLogout, user } = useAuth();
>>>>>>> Stashed changes

    const handleSignOut = async () => {
        onLogout!();
    };

    return (
<<<<<<< Updated upstream
        <View className="items-center justify-center flex-1 gap-5">
            <Button onPress={handleSignOut}>Sign Out</Button>
        </View>
=======
        <SafeAreaView className="justify-center flex-1">
            <View className="px-[8%] gap-5 flex-col items-center">
                <Text className="text-lg">Welcome {user?.first_name}</Text>
                <Button onPress={handleSignOut}>Sign Out</Button>
            </View>
        </SafeAreaView>
>>>>>>> Stashed changes
    );
};

export default Profile;
