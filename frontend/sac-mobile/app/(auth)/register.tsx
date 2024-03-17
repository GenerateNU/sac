import React from 'react';
import { ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { Button } from '@/components/button';
import Wordmark from '@/components/wordmark';

import RegistrationForm from './_components/registration-form';

const Register = () => {
    return (
        <SafeAreaView className="bg-neutral-500" edges={['top']}>
            <ScrollView>
                <View className="px-[8%] pb-[9%]">
                    <View className="flex flex-row justify-between mx-auto w-full items-center pt-[3%] pb-[5.5%]">
                        <Wordmark 
                        textColor="text-white"/>
                        <Button
                            onPress={() => router.push('/(auth)/login')}
                            variant="secondary"
                            size="sm"
                        >
                            Login
                        </Button>
                    </View>
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
        </SafeAreaView>
    );
};

export default Register;
