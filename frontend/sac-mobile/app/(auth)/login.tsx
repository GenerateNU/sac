import React from 'react';
import { Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import Wordmark from '@/components/wordmark';
import LoginForm from './_components/login-form';

const Login = () => {
    return (
        <SafeAreaView className="bg-neutral-500 h-[100%]" edges={['top']}>
            <View className="flex-1">
                <View className="px-[8%] pb-[10%]">
                    <View className="pt-[1%]">
                        <Wordmark textColor="text-white"/>
                    </View>
                    <View className="pt-[9.5%] pb-[6%]">
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
            </View>
        </SafeAreaView>
    );
};

export default Login;
