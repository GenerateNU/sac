import React from 'react';
import {
    Keyboard,
    KeyboardAvoidingView,
    Platform,
    ScrollView,
    Text,
    TouchableWithoutFeedback,
    View
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { LoginForm } from './_components/login-form';
import { Wordmark } from '@/components/wordmark';

const Login = () => {
    return (
        <>
            <TouchableWithoutFeedback
                onPress={Keyboard.dismiss}
                accessible={false}
            >
                <SafeAreaView
                    className="bg-neutral-500 h-[100%]"
                    edges={['top']}
                >
                    <View className="flex-1">
                        <View className="px-[8%] pb-[10%]">
                            <View className="pt-[1%]">
                                <Wordmark color={'default'} />
                            </View>
                            <View className="pt-[16.5%] pb-[7%]">
                                <Text className="text-5xl font-bold text-white">
                                    Let's go
                                </Text>
                            </View>
                            <Text className="text-white pt-[0.8%] text-lg leading-6">
                                Discover, follow, and join all the clubs &
                                events Northeastern has to offer
                            </Text>
                        </View>
                        <View className="bg-white pt-[13%] pb-[2%] flex-1 rounded-tl-3xl rounded-tr-3xl px-[8%]">
                            <LoginForm />
                            <View className="mt-[9%] flex-row justify-center">
                                <Text className="font-bold">
                                    Not a student?
                                </Text>
                                <Text>
                                    {' '}
                                    Continue as a{' '}
                                    <Text className="underline">guest</Text>
                                </Text>
                            </View>
                        </View>
                    </View>
                </SafeAreaView>
            </TouchableWithoutFeedback>
        </>
    );
};

export default Login;
