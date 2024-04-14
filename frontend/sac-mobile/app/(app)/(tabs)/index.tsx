import React from 'react';
import { ScrollView, View } from 'react-native';

import { HomepageList } from '@/app/(app)/(tabs)/_components/homepage';
import { useAuthStore } from '@/hooks/use-auth';
import { useUserFollowing } from '@/hooks/use-user';

import { FollowingHeader } from './_components/following-header';

const Homepage = () => {
    const { user } = useAuthStore();
    const userId: string = user?.id!;
    const {
        data: userFollowingData,
        isLoading,
        error
    } = useUserFollowing(userId);
    return (
        <ScrollView showsVerticalScrollIndicator={false}>
            <View className="pb-[28%]">
                <View className="pt-[2%]">
                    {userFollowingData && userFollowingData.length > 0 && (
                        <View className="pt-[3%] pb-[0.5%]">
                            <FollowingHeader id={userId} />
                        </View>
                    )}
                    <HomepageList />
                </View>
            </View>
        </ScrollView>
    );
};

export default Homepage;
