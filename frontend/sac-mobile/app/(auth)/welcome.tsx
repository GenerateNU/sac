import React from 'react';
import { Image, Text, View } from 'react-native';

import { router } from 'expo-router';

import { Button } from '@/components/button';

const Welcome = () => {
    return (
        <View className="flex-col flex-1 bg-[#f5f5f5] px-[5%] pb-5 gap-y-5 justify-evenly">
            <Image
                source={{ uri: 'https://i.imgur.com/pLtAoSt.png' }}
                height={275}
                className="rounded-2xl"
            />
            <View>
                <Text className="flex text-5xl font-semibold text-black mb-3">
                    Welcome to Student Activity Calendar
                </Text>
                <Text className="text-2xl text-black leading-8 pl-[1%]">
                    Discover, follow, and join all the clubs & events
                    Northeastern has to offer
                </Text>
            </View>
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
