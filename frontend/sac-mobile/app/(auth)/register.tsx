import React from 'react';
import {
    KeyboardAvoidingView,
    Platform,
    ScrollView,
    Text,
    View
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import { Wordmark } from '@/components/wordmark';

import { RegistrationForm } from './_components/registration-form';

const Register = () => {
    return (
        <KeyboardAvoidingView
            behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
            className="flex-1"
        >
            <SafeAreaView className="bg-neutral-500" edges={['top']}>
                <ScrollView>
                    <View className="bg-neutral-500">
                        <View className="px-[8%] flex flex-row justify-between mx-auto w-full items-center pb-[5.5%]">
                            <View className="pl-[0.5%]">
                                <Wordmark color={'default'} />
                            </View>
                            <Button
                                onPress={() => router.push('/(auth)/login')}
                                variant="secondary"
                                size="sm"
                            >
                                Login
                            </Button>
                        </View>
                        <View className="px-[8%] pt-[8%] pb-[7.5%]">
                            <Text className="text-5xl font-bold text-white">
                                Sign up
                            </Text>
                        </View>
                        <Text className="px-[8%] mb-[9%] text-lg leading-6 text-white">
                            Discover, follow, and join all the clubs & events
                            Northeastern has to offer
                        </Text>
                        <View className="bg-white px-[8%] pt-[13%] rounded-t-3xl">
                            <RegistrationForm />
                        </View>
                    </View>
                </ScrollView>
            </SafeAreaView>
        </KeyboardAvoidingView>
    );
};

export default Register;
