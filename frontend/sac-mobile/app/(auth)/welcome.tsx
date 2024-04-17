import React from 'react';
import { Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';

const Welcome = () => {
    return (
        <View className="flex-col bg-neutral-500 pb-[14%] px-[8%]">
            <View className="bg-neutral-400 h-[45%] w-ful rounded-xl mt-[5%]" />
                <Text className="text-white flex flex-wrap text-5xl font-bold pt-[20%] pb-[3%]">
                    Welcome to Student Activity Calendar
                </Text>
            <Text className="text-2xl text-white leading-8 pl-[1%] pb-[10%]">
                Discover, follow, and join all the clubs & events Northeastern
                has to offer
            </Text>
            <View className="flex-row justify-end">
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
