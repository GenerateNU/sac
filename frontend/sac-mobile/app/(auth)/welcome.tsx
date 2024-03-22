import React from 'react';
import { SafeAreaView, Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';
import Wordmark from '@/components/wordmark';

const Welcome = () => {
    return (
        <SafeAreaView className="flex-col mb-[8%] mx-[8%]">
            <Wordmark />
            <View className="bg-gray-500 h-[45%] w-ful rounded-xl mt-[5%]" />
            <Text className="h-[18%] text-6xl font-bold mt-[10%]">
                Welcome to StudCal
            </Text>
            <Text className="text-2xl leading-8 pl-[1%] pb-[10%]">
                Discover, follow, and join all the clubs & events Northeastern
                has to offer
            </Text>
            <View className="flex-row justify-end">
                <Button
                    size="lg"
                    variant="default"
                    onPress={() => router.push('/(app)/homepage')}
                >
                    Get Started
                </Button>
            </View>
        </SafeAreaView>
    );
};

export default Welcome;
