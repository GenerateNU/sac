import React from 'react';
import { SectionList, View } from 'react-native';

import { StatusBar } from 'expo-status-bar';

import { HomepageList } from '@/app/(app)/(tabs)/_components/homepage';
import { useAuthStore } from '@/hooks/use-auth';

import { FollowingHeader } from './_components/following-header';
import { useAuth } from '@/context/AuthContext';


const HomePage = () => {
    const { user } = useAuthStore();

    return (
        <>
            <StatusBar style="dark" />
            <SectionList
                sections={[
                    {
                        data: [
                            <FollowingHeader
                                key={user?.id}
                                id={user?.id || ''}
                            />
                        ]
                    },
                    {
                        data: [<HomepageList key={user?.id} />]
                    }
                ]}
                keyExtractor={(_, index) => index.toString()}
                renderItem={({ item }) => <>{item}</>}
                renderSectionHeader={() => <View className="px-[6%] pt-[3%]" />}
                stickySectionHeadersEnabled={false}
            />
        </>
    );
};

export default HomePage;
