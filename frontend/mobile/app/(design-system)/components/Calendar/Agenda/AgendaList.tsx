import React, { useCallback, useContext, useEffect, useRef, useState } from 'react';
import { Dimensions, ListRenderItem, ViewToken } from 'react-native';
import { CalendarContext } from 'react-native-calendars-sac';
import { UpdateSources } from 'react-native-calendars-sac/src/expandableCalendar/commons';
import { FlatList } from 'react-native-gesture-handler';

import { Box } from '../../Box/Box';
import { EventCardCalendarSkeleton } from '../../EventCard/Skeletons/EventCardCalendarSkeleton';
import { EventSection } from '../mockData';
import { AgendaDateSection } from './AgendaDateSection';
import { Spacing } from '@/app/(design-system)/shared/spacing';

const { width: screenWidth } = Dimensions.get('window');
const width = screenWidth - (Spacing.s * 2);

export type AgendaListProps = {
    sections: EventSection[];
    fetchData: (direction: 'up' | 'down') => void;
    isLoading: boolean;
    autoScroll: boolean;
    setAutoScroll: React.Dispatch<React.SetStateAction<boolean>>;
    setCurrentIndex: React.Dispatch<React.SetStateAction<number | null>>;
    ref: React.MutableRefObject<FlatList<EventSection> | null>;
};

export default function AgendaList({
    sections,
    fetchData,
    isLoading,
    autoScroll,
    setAutoScroll,
    setCurrentIndex,
    ref,
}: AgendaListProps) {
    const { date, setDate } = useContext(CalendarContext);
    const currentSection = useRef<string | null>(null);
    const [scrolling, setScrolling] = useState(false);

    const scrollToDate = useCallback(
        (dateSections: EventSection[]) => {
            const index = dateSections.findIndex(
                (section) => section.date === date
            );
            if (index !== -1) {
                setScrolling(true);
                ref.current?.scrollToIndex({
                    index,
                    animated: true,
                    viewOffset: 1,
                    viewPosition: 0
                });
                setTimeout(() => {
                    setScrolling(false);
                    setAutoScroll(false);
                }, 500);
                currentSection.current = dateSections[index].date;
            }
        },
        [date, setAutoScroll]
    );

    const onViewableItemsChanged = ({
        viewableItems
    }: {
        viewableItems: ViewToken<EventSection>[];
    }) => {
        const currentItem = viewableItems?.[0]?.item.date;
        if (currentItem && currentItem !== date && !scrolling) {
            currentSection.current = currentItem;
            setDate?.(viewableItems[0].item.date, UpdateSources.LIST_DRAG);
        }
        setCurrentIndex(viewableItems[0].index);
    };

    useEffect(() => {
        if (currentSection.current === date || !autoScroll) return;
        scrollToDate(sections);
    }, [date, autoScroll, scrollToDate, sections]);

    const renderItem: ListRenderItem<EventSection> = ({ item, index }) => {
        return (
            <Box width={width}>
                <AgendaDateSection 
                    key={index}
                    section={item}
                />
            </Box>
            );
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
            onEndReached={() => fetchData('down')}
            onStartReached={() => fetchData('up')}
            ref={ref}
            showsVerticalScrollIndicator={false}
            onViewableItemsChanged={onViewableItemsChanged}
            horizontal
            snapToAlignment="center"
            decelerationRate="fast"
            pagingEnabled
            getItemLayout={(_, index) => ({
                length: width,
                offset: width * index,
                index,
            })}
            initialScrollIndex={7}
            maintainVisibleContentPosition={{
                minIndexForVisible: 1,
                autoscrollToTopThreshold: 10,
            }}
        />
    );
}