import React from 'react';
import { ScrollView, View } from 'react-native';

import { HomepageList } from '@/app/(app)/(tabs)/_components/homepage';
import { useAuthStore } from '@/hooks/use-auth';

import { FollowingHeader } from './_components/following-header';

const Homepage = () => {
    const { user } = useAuthStore();
    const userId: string = user?.id!;

    return (
        <ScrollView showsVerticalScrollIndicator={false}>
            <View className="pb-[28%]">
                <View className="pt-[6%]">
                    <FollowingHeader id={userId} />
                    <HomepageList />
                </View>
            </View>
        </ScrollView>
    );
};

export default Homepage;
