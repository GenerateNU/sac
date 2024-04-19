import React from 'react';
import { SectionList, View, Text } from 'react-native';

import { StatusBar } from 'expo-status-bar';

import { HomepageList } from '@/app/(app)/(tabs)/_components/homepage';
import { useAuthStore } from '@/hooks/use-auth';
import { HappeningToday } from './_components/happening-today';
import { RecommendedClubs } from './_components/recommended-clubs';


const HomePage = () => {
    const { user } = useAuthStore();

    return (
        <>
            <StatusBar style="dark" />
            <View className="pt-[2%]">
                <Text className="text-[18px] font-bold px-[6%]">Happening Today</Text>
                <HappeningToday />
            </View>
            {/* <View className="pt-[2%]">
                <HomepageList />
            </View> */}
            <View className="pt-[2%]">
                <Text className="text-[18px] font-bold px-[6%]">Recommended Clubs</Text>
                <RecommendedClubs />
            </View>
            {/* <SectionList
                sections={[
                    {
                        data: [<HappeningToday key={user?.id} />],
                        title: 'Happening Today',
                    },
                    // {
                    //     data: [<HomepageList key={user?.id} />],
                    //     title: 'Your Upcoming Events'
                    // },
                    {
                        data: [<RecommendedClubs key={user?.id} />],
                        title: 'Recommended Clubs'
                    }
                ]}
                keyExtractor={(_, index) => index.toString()}
                renderSectionFooter={() => <View className="px-[6%] pb-[3%]" />}
                renderSectionHeader={({ section: { title } }) => <View className="px-[6%] pt-[3%]"><Text className="text-[18px] font-bold">{title}</Text></View>}
                stickySectionHeadersEnabled={false}
                renderItem={({ item }) => item}
            /> */}
        </>
    );
};

export default HomePage;
