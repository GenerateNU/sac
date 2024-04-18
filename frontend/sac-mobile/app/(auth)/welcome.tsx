import React from 'react';
import { Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';

const Welcome = () => {
    return (
        <View className="flex-col flex-1 bg-neutral-500 px-[5%] gap-y-5">
            <View className="bg-neutral-400 h-[38%] w-full rounded-xl" />
            <Text className="flex flex-wrap text-5xl font-bold text-white">
                Welcome to Student Activity Calendar
            </Text>
            <Text className="text-2xl text-white leading-8 pl-[1%]">
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
