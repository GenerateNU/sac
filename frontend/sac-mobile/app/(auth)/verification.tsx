import { View, Text, ScrollView } from 'react-native'
import React, { useEffect, useState } from 'react'
import { useSignUp } from '@clerk/clerk-expo';
import Spinner from 'react-native-loading-spinner-overlay';
import { Button } from '@/components/button';
import Input from '@/components/input';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { Wordmark } from '@/components/wordmark';
import { router } from 'expo-router';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useAuthStore } from '@/hooks/use-auth';
import { requestVerification } from '@/services/auth';
import VerificationForm from './_components/verification-form';

const Verification = () => {
    return (
        <SafeAreaView className="bg-neutral-500" edges={['top']}>
            <ScrollView>
                <View className="px-[8%] pb-[9%]">
                    <View className="flex flex-row justify-between mx-auto w-full items-center pt-[3%] pb-[5.5%]">
                        <Wordmark />
                    </View>
                    <View className="pt-[9%] pb-[7.5%]">
                        <Text className="text-5xl font-bold text-white">
                            Verify your email
                        </Text>
                    </View>
                    <Text className="text-lg leading-6 text-white ">
                        Please enter the verification code sent to your email.
                    </Text>
                </View>
                <View className="bg-white px-[8%] pt-[13%] pb-[20%] mt-[50%] rounded-t-3xl">
                    <VerificationForm />
                </View>
            </ScrollView>
        </SafeAreaView>
    );
}

export default Verification;
