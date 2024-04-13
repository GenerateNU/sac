import React from 'react';
import { Text, View } from 'react-native';

import { useClubs } from '@/hooks/use-club';
import { useEvents } from '@/hooks/use-event';
import { HomepageFAQ } from '@/lib/const';
import { isClub, isEvent, isFAQ, shuffleArray } from '@/lib/utils';
import { Club } from '@/types/club';
import { Event } from '@/types/event';
import { FAQ } from '@/types/faq';

import { ClubHomepageCard } from './club-homepage-card';
import { EventHomePageCard } from './event-homepage-card';
import { FAQHomePageCard } from './faq-homepage-card';

export type HomepageItem = Event | Club | FAQ;

const HomepageList = () => {
    const { data: events } = useEvents();
    const { data: clubs } = useClubs();

    const getHomepageItems = () => {
        const allItems: HomepageItem[] = [];
        if (events !== undefined) {
            allItems.push(...events);
        }
        if (clubs !== undefined) {
            allItems.push(...clubs);
        }
        allItems.push(...HomepageFAQ);
        const shuffledItems = shuffleArray(allItems);
        return shuffledItems;
    };

    return (
        <View className="px-[6%] pt-[3%]">
            {getHomepageItems().map((item: FAQ | Club | Event, index) => {
                if (isClub(item)) {
                    return <ClubHomepageCard key={index} club={item as Club} />;
                } else if (isEvent(item)) {
                    return (
                        <EventHomePageCard key={index} event={item as Event} />
                    );
                } else if (isFAQ(item)) {
                    return <FAQHomePageCard key={index} faq={item} />;
                }
            })}
        </View>
    );
};

export { HomepageList };
