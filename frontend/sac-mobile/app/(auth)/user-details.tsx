import { SafeAreaView, ScrollView, Text, View } from 'react-native';

import UserDetailForm from './_components/user-details-form';
import { StatusBar } from 'expo-status-bar';

const UserDetails = () => {
    return (
        <SafeAreaView>
            <ScrollView
                className="px-[8%] pb-[9%]">
                <StatusBar style="dark" />
                <Text className="font-bold text-5xl pt-[8%] pb-[10%]">
                    Let's learn more about you
                </Text>
                <View className="pb-[10%]">
                <UserDetailForm />
            </View>
            </ScrollView>
        </SafeAreaView>
    );
};

export default UserDetails;
