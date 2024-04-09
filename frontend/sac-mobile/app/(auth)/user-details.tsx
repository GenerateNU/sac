import { Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { Wordmark } from '@/components/wordmark';

import UserDetailForm from './_components/user-details-form';

const UserDetails = () => {
    return (
        <SafeAreaView>
            <View className="px-[8%] pb-[9%]">
                <Wordmark color={'primary'} />
                <Text className="font-bold text-5xl pt-[15%] pb-[10%]">
                    Let's learn more about you
                </Text>
                <UserDetailForm />
            </View>
        </SafeAreaView>
    );
};

export default UserDetails;
