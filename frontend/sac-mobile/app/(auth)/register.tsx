import React from 'react';
import { KeyboardAvoidingView, ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { RegistrationForm } from './_components/registration-form';

const Register = () => {
    return (
        <SafeAreaView className="bg-neutral-500" edges={['top']}>
            <KeyboardAvoidingView behavior={'padding'}>
                <ScrollView>
                    <View className="px-[8%] pb-[9%]">
                        <View className="pt-[9%] pb-[7.5%]">
                            <Text className="text-5xl font-bold text-white">
                                Sign up
                            </Text>
                        </View>
                        <Text className="text-lg leading-6 text-white">
                            Discover, follow, and join all the clubs & events
                            Northeastern has to offer
                        </Text>
                    </View>
                    <View className="bg-white px-[8%] pt-[13%] rounded-t-3xl">
                        <RegistrationForm />
                    </View>
                </ScrollView>
            </KeyboardAvoidingView>
        </SafeAreaView>
    );
};

export default Register;
