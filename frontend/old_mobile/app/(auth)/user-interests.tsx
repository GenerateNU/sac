import React from 'react';
import { Text, View } from 'react-native';

import { StatusBar } from 'expo-status-bar';

import UserInterestsForm from './_components/user-interest-form';

const UserInterests = () => {
    return (
        <View className="px-[8%] pt-[4%]">
            <StatusBar style="dark" />
            <Text className="text-5xl pt-[6%] pb-[4%] font-bold">
                What are you interested in?
            </Text>
            <UserInterestsForm />
        </View>
    );
};

export default UserInterests;
