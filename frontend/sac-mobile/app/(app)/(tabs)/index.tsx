import React from 'react';
import { SectionList, Text, View } from 'react-native';

import { StatusBar } from 'expo-status-bar';

import { HappeningToday } from './_components/happening-today';
import { RecommendedClubs } from './_components/recommended-clubs';
import { UpcomingEvents } from './_components/upcoming-events';

const HomePage = () => {
    return (
        <>
            <StatusBar style="dark" />
            <SectionList
                sections={[
                    { title: 'Happening Today', data: [<HappeningToday />] },
                    { title: 'Upcoming Events', data: [<UpcomingEvents />] },
                    { title: 'Recommended Clubs', data: [<RecommendedClubs />] }
                ]}
                keyExtractor={(_, index) => index.toString()}
                renderItem={({ item }) => <View>{item}</View>}
                renderSectionHeader={({ section: { title } }) => (
                    <Text className="text-[18px] font-bold px-[6%] py-[3%]">
                        {title}
                    </Text>
                )}
                stickyHeaderHiddenOnScroll={true}
            />
        </>
    );
};

export default HomePage;
