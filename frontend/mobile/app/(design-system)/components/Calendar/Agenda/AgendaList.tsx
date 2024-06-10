import { useCallback, useContext, useEffect, useRef, useState } from 'react';
import { ListRenderItem, ViewToken } from 'react-native';
import { CalendarContext } from 'react-native-calendars-sac';
import { UpdateSources } from 'react-native-calendars-sac/src/expandableCalendar/commons';
import { FlatList } from 'react-native-gesture-handler';

import { Box } from '../../Box/Box';
import { EventCardCalendarSkeleton } from '../../EventCard/Skeletons/EventCardCalendarSkeleton';
import { EventSection } from '../mockData';
import { AgendaDateSection } from './AgendaDateSection';

// Add back my old code:
export type AgendaListProps = {
    sections: EventSection[];
    fetchData: () => void;
    isLoading: boolean;
    autoScroll: boolean;
    setAutoScroll: React.Dispatch<React.SetStateAction<boolean>>;
};

export default function AgendaList({
    sections,
    fetchData,
    isLoading,
    autoScroll,
    setAutoScroll
}: AgendaListProps) {
    const { date, setDate } = useContext(CalendarContext);
    const eventList = useRef<FlatList<EventSection> | null>(null);
    const topSection = useRef<string | null>(null);
    const [scrolling, setScrolling] = useState(false);

    const scrollToDate = useCallback(
        (dateSections: EventSection[]) => {
            const index = dateSections.findIndex(
                (section) => section.date === date
            );
            if (index !== -1) {
                setScrolling(true);
                eventList.current?.scrollToIndex({
                    index,
                    animated: true,
                    viewOffset: 1,
                    viewPosition: 0
                });
                setTimeout(() => {
                    setScrolling(false);
                    setAutoScroll(false);
                }, 500);
                topSection.current = dateSections[index].date;
            }
        },
        [date, setAutoScroll]
    );

    const onViewableItemsChanged = ({
        viewableItems
    }: {
        viewableItems: ViewToken<EventSection>[];
    }) => {
        const topItem = viewableItems?.[0]?.item.date;
        if (topItem && topItem !== date && !scrolling) {
            topSection.current = topItem;
            setDate?.(viewableItems[0].item.date, UpdateSources.LIST_DRAG);
        }
    };

    useEffect(() => {
        if (topSection.current === date || !autoScroll) return;
        scrollToDate(sections);
    }, [date, autoScroll, scrollToDate, sections]);

    const renderItem: ListRenderItem<EventSection> = ({ item, index }) => {
        return <AgendaDateSection key={index} section={item} />;
    };

    const loadingComponent = () => {
        if (isLoading) {
            return (
                <Box flexDirection="column" gap="s" marginVertical="s">
                    <EventCardCalendarSkeleton />
                    <EventCardCalendarSkeleton />
                    <EventCardCalendarSkeleton />
                </Box>
            );
        } else {
            return null;
        }
    };

    return (
        <FlatList
            data={sections}
            renderItem={renderItem}
            onEndReached={fetchData}
            onEndReachedThreshold={1}
            ref={eventList}
            showsVerticalScrollIndicator={false}
            onViewableItemsChanged={onViewableItemsChanged}
            ListFooterComponent={loadingComponent}
        />
    );
}
