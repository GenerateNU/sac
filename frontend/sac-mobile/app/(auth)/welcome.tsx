import React from 'react';
import { Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';

const Welcome = () => {
    return (
        <View className="flex-col bg-neutral-500 pb-[8%] px-[8%]">
            <View className="bg-neutral-400 h-[38%] w-ful rounded-xl mt-[2%]" />
            <Text className="text-white flex flex-wrap text-5xl font-bold pt-[15%] pb-[3%]">
                Welcome to Student Activity Calendar
            </Text>
            <Text className="text-2xl text-white leading-8 pl-[1%] pb-[10%]">
                Discover, follow, and join all the clubs & events Northeastern
                has to offer
            </Text>
            <View className="flex-row justify-end mb-[10%]">
                <Button
                    size="lg"
                    variant="default"
                    onPress={() => router.replace('/(auth)/login')}
                >
                    Get Started
                </Button>
            </View>
        </View>
    );
};

export default Welcome;
