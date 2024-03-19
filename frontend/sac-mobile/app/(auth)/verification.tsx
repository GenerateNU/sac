import { Button } from '@/components/button';
import { router } from 'expo-router';
import React from 'react';
import { ScrollView, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

const Verification = () => {
    return (
        <SafeAreaView className="h-full bg-neutral-500" edges={['top']}>
            <ScrollView>
                <Button onPress={() => router.push('/(auth)/user-details')}>
                    Go to User Details
                </Button>
            </ScrollView>
        </SafeAreaView>
    );
};

export default Verification;
