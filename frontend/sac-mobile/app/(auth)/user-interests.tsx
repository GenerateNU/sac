import React from 'react';
import { SafeAreaView, Text, View } from 'react-native';
import { QueryClient, QueryClientProvider, useQuery } from 'react-query';
import Wordmark from '@/components/wordmark';

import UserInterestsForm from './_components/user-interest-form';

const UserInterests = () => {
    return (
        <SafeAreaView>
                <View className="px-[8%] pt-[4%]">
                    <View className="flex flex-row">
                        <Wordmark />
                    </View>
                    <Text className="text-5xl pt-[6%] pb-[5%] font-bold">
                        What are you interested in?
                    </Text>
                    <UserInterestsForm />
                </View>
        </SafeAreaView>
    );
};

export default UserInterests;
