import React from 'react';
import { SectionList, Text, View } from 'react-native';

import { StatusBar } from 'expo-status-bar';

import { useUserFollowingQuery } from '@sac/lib';

import { useAuthStore } from '@/hooks/use-auth';

import { HappeningToday } from './_components/happening-today';
import { RecommendedClubs } from './_components/recommended-clubs';
import { UpcomingEvents } from './_components/upcoming-events';

const HomePage = () => {
    const { user } = useAuthStore();
    const {
        data: clubs,
        isLoading: cIsLoading,
        error: cError
    } = useUserFollowingQuery(user?.id!);
    return (
        <>
            <StatusBar style="dark" />
            <View>
                {cIsLoading && <Text>Loading...</Text>}
                {cError && <Text>Error...</Text>}
                {clubs && clubs.length === 0 && (
                    <Text>Follow some clubs to see them here!</Text>
                )}
                {clubs && clubs.length > 0 && (
                    <Text>Following Clubs: {clubs.length}</Text>
                )}
                {clubs &&
                    clubs.map((club) => <Text key={club.id}>{club.name}</Text>)}
            </View>
            <View className="px-[5%]">
                <SectionList
                    sections={[
                        {
                            title: 'Happening Today',
                            data: [<HappeningToday />]
                        },
                        {
                            title: 'Upcoming Events',
                            data: [<UpcomingEvents />]
                        },
                        {
                            title: 'Recommended Clubs',
                            data: [<RecommendedClubs />]
                        }
                    ]}
                    keyExtractor={(_, index) => index.toString()}
                    renderItem={({ item }) => <View>{item}</View>}
                    renderSectionHeader={({ section: { title } }) => (
                        <Text className="my-3 text-xl font-light leading-relaxed text-zinc-900">
                            {title}
                        </Text>
                    )}
                    stickyHeaderHiddenOnScroll={true}
                    stickySectionHeadersEnabled={false}
                />
            </View>
        </>
    );
};

export default HomePage;
