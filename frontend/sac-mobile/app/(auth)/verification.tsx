import React from 'react';
import { ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { Wordmark } from '@/components/wordmark';

import { VerificationForm } from './_components/verification-form';
import { StatusBar } from 'expo-status-bar';

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
