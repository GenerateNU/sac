import React from 'react';
import { ScrollView } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import { Button } from '@/components/button';

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
