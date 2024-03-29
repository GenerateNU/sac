import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import { View } from 'react-native';

const UserDetails = () => {
    return (
        <SafeAreaView>
            <View className="px-[8%] pb-[9%]">
                <View className="flex-row justify-end pt-[5%]">
                    <Button
                        size="lg"
                        variant="default"
                        onPress={() => router.push('/(app)/')}
                    >
                        Continue
                    </Button>
                </View>
            </View>
        </SafeAreaView>
    );
};

export default UserDetails;