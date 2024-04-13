import React from 'react';
import { KeyboardAvoidingView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { LoginForm } from './_components/login-form';

const Login = () => {
    return (
        <SafeAreaView className="bg-neutral-500 h-[100%]" edges={['top']}>
            <KeyboardAvoidingView behavior={'padding'} className="flex-1">
                <View className="px-[8%] pb-[8%]">
                    <View className="pt-[2%] pb-[1%]">
                        <Text className="text-5xl font-bold text-white">
                            Let's go
                        </Text>
                    </View>
                    <Text className="text-white pt-[0.8%] text-lg leading-6">
                        Discover, follow, and join all the clubs & events
                        Northeastern has to offer
                    </Text>
                </View>
                <View className="bg-white pt-[13%] pb-[2%] flex-1 rounded-tl-3xl rounded-tr-3xl px-[8%]">
                    <LoginForm />
                    <View className="mt-[9%] flex-row justify-center">
                        <Text className="font-bold">Not a student?</Text>
                        <Text>
                            {' '}
                            Continue as a{' '}
                            <Text className="underline">guest</Text>
                        </Text>
                    </View>
                </View>
            </KeyboardAvoidingView>
        </SafeAreaView>
    );
};

export default Login;
