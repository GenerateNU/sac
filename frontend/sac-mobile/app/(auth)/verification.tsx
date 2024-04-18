import React from 'react';
import { Text, View } from 'react-native';

import { StatusBar } from 'expo-status-bar';

import EmailConfirm from '@/assets/images/svg/email-confirm.svg';

import { VerificationForm } from './_components/verification-form';

const Verification = () => {
    return (
        <View className="h-full bg-white">
            <StatusBar style="dark" />
            <View className="px-[8%] pb-[9%]">
                <View className="items-center pt-[15%]">
                    <EmailConfirm />
                </View>
                <View className="pt-[9%] pb-[7.5%]">
                    <Text className="text-5xl font-bold text-black">
                        Confirm your email address
                    </Text>
                </View>
                <Text className="text-lg leading-6 text-gray-500 ">
                    Enter the 6-digit sent to your email
                </Text>
            </View>
            <View className=" px-[8%] pt-[13%] pb-[20%]">
                <VerificationForm />
            </View>
        </View>
    );
};

export default Verification;
