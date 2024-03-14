import React from 'react';
import { Pressable, SafeAreaView, StyleSheet, Text, View } from 'react-native';

import { router } from 'expo-router';

import Wordmark from '@/components/Wordmark';
import { Button } from '@/components/button';

const Welcome = () => {
    const redirect = () => {
        router.push('/(auth)/login');
    };

    return (
        <SafeAreaView className="flex-col mb-[8%] mx-[8%]">
            <Wordmark />
            <View className="bg-gray-500 h-[45%] w-ful rounded-xl mt-[5%]"></View>
            <Text className="h-[18%] text-6xl font-bold mt-[10%]">Welcome to StudCal</Text>
            <Text className="text-2xl leading-8 pl-[1%] pb-[10%]">
                Discover, follow, and join all the clubs & events Northeastern
                has to offer
            </Text>
            <View className="flex-row justify-end">
                <Button
                    size="lg"
                    variant="default"
                    onPress={() => router.push('/(auth)/login')}
                >Get Started</Button>
            </View>
        </SafeAreaView>
    );
};

export default Welcome;