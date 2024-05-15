import React from 'react';
import { FlatList, Text, View } from 'react-native';

// import { useClubs } from '@/hooks/use-club';
// import { useEvents } from '@/hooks/use-event';
import { Club, Event, useClubsQuery, useEventsQuery } from '@sac/lib';

import { HomepageFAQ } from '@/lib/const';
import { isClub, isEvent, isFAQ } from '@/lib/utils';
import { FAQ } from '@/types/faq';

import { ClubHomePageCard } from './club-homepage-card';
import { EventHomePageCard } from './event-homepage-card';
import { FAQHomePageCard } from './faq-homepage-card';

export type HomepageItem = Event | Club | FAQ;

const HomepageList = () => {
    const {
        data: events,
        isLoading: eIsLoading,
        error: eError
    } = useEventsQuery({});
    const {
        data: clubs,
        isLoading: cIsLoading,
        error: cError
    } = useClubsQuery({});

    if (eIsLoading || cIsLoading) {
        return <Text>Loading...</Text>;
    }

    if (eError || cError) {
        return <Text>Error: </Text>;
    }

    const getHomepageItems = () => {
        const allItems: HomepageItem[] = [];
        if (events !== undefined) {
            allItems.push(...events);
        }
        if (clubs !== undefined) {
            allItems.push(...clubs);
        }
        allItems.push(...HomepageFAQ);
        return allItems.sort(() => Math.random() - 0.5);
    };

    const renderItems = (items: HomepageItem[]) => {
        return items.map((item: HomepageItem, index) => {
            if (isClub(item)) {
                const club = item as Club;
                if (club.name === 'SAC') {
                    return null;
                }
                return <ClubHomePageCard key={index} club={club} />;
            } else if (isEvent(item)) {
                return <EventHomePageCard key={index} event={item as Event} />;
            } else if (isFAQ(item)) {
                return <FAQHomePageCard key={index} faq={item as FAQ} />;
            }
        });
    };

    return (
        <View className="px-[6%] pt-[3%]">
            <FlatList
                data={renderItems(getHomepageItems())}
                renderItem={({ item }) => <>{item}</>}
                keyExtractor={(_, index) => index.toString()}
                showsVerticalScrollIndicator={false}
            />
        </View>
    );
};

export { HomepageList };
