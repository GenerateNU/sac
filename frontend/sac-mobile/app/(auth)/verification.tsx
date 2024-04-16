import React from 'react';
import { Text, View } from 'react-native';

import { StatusBar } from 'expo-status-bar';

import { VerificationForm } from './_components/verification-form';

const Verification = () => {
    return (
        <View className="h-full bg-white">
            <StatusBar style="dark" />
            <View className="px-[8%] pb-[9%]">
                <View className="pt-[9%] pb-[7.5%]">
                    <Text className="text-5xl font-bold text-gray-500">
                        Verify your email
                    </Text>
                </View>
                <Text className="text-lg leading-6 text-gray-500 ">
                    Please enter the verification code sent to your email.
                </Text>
            </View>
            <View className=" px-[8%] pt-[13%] pb-[20%]">
                <VerificationForm />
            </View>
        </View>
    );
};

export default Verification;
