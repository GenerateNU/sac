import React, { useCallback, useEffect, useRef, useState } from 'react';
import {
    CalendarProvider,
    ExpandableCalendar
} from 'react-native-calendars-sac';
import { Theme } from 'react-native-calendars-sac/src/types';

import { getMarkedDates } from '@/utils/mock';

import { createStyles } from '../../theme';
import { Box } from '../Box/Box';
import { GlobalLayout } from '../GlobalLayout/GlobalLayout';
import AgendaList from './Agenda/AgendaList';
import { EventSection, fetchEvents } from './mockData';
import { Dimensions } from 'react-native';
import { FlatList } from 'react-native-gesture-handler';

const FETCH_RANGE = 1000 * 60 * 60 * 24 * 13;
const NEXT_DAY = 1000 * 60 * 60 * 24;

const calendarTheme: Theme = {
    selectedDayBackgroundColor: 'black',
    todayTextColor: 'black',
    arrowColor: 'black',
    textMonthFontWeight: 'bold'
};

export default function Calendar() {
    const [isLoading, setIsLoading] = useState(true);
    const [events, setEvents] = useState<EventSection[]>([]);
    const [autoScroll, setAutoScroll] = useState<boolean>(false);
    const [minDate, _] = useState(new Date().getTime());
    const [maxDate, setMaxDate] = useState(new Date().getTime());
    const [currentIndex, setCurrentIndex] = useState<number | null>(null);
    const eventList = useRef<FlatList<EventSection> | null>(null);

    const fetchData = useCallback(
        async (direction = 'up') => {
            console.log('fetching data!')
            setIsLoading(true);
            setTimeout(() => {
                const time =
                    direction === 'up' ? minDate - FETCH_RANGE : maxDate;
                fetchEvents(time, FETCH_RANGE).then((data) => {
                    if (direction === 'down') {
                        console.log('setting date end')
                        setEvents((oldData) => [...oldData, ...data]);
                    } else {
                        console.log('setting date beginning')
                        prependData(data)
                    }
                    setIsLoading(false);
                    setMaxDate((oldDate) => oldDate + FETCH_RANGE + NEXT_DAY);
                });
            }, 0);
        },
        [minDate, maxDate]
    );

    const fetchNextPage = async (direction: 'up' | 'down') => {
        if (isLoading) return;
        fetchData(direction);
    };

    useEffect(() => {
        fetchData('down');
    }, []);

    const prependData = (newSections: EventSection[]) => {
        console.log(currentIndex);
        if (currentIndex !== null) {
            const newScrollOffset = (currentIndex + newSections.length) * Dimensions.get('window').width;
            console.log(newScrollOffset);
            setEvents((prevSections) => [...newSections, ...prevSections]);
            console.log("scrolling");
            eventList.current?.scrollToOffset({ offset: newScrollOffset, animated: false });
        }
    };

    return (
        <CalendarProvider
            date={new Date().toISOString().split('T')[0]}
            theme={calendarTheme}
            disabledOpacity={0.5}
        >
            <Box {...styles.expandableCalendarContainer}>
                <Box {...styles.expandableCalendarSubContainer}>
                    <ExpandableCalendar
                        horizontal
                        pagingEnabled={true}
                        hideArrows
                        theme={calendarTheme}
                        markedDates={getMarkedDates()} // TODO: Add marked dates
                        onDayPress={() => setAutoScroll(true)}
                        closeOnDayPress={false}
                        disableWeekScroll={true}
                    />
                </Box>
            </Box>
            <GlobalLayout>
                <AgendaList
                    sections={events}
                    fetchData={fetchNextPage}
                    isLoading={isLoading}
                    autoScroll={autoScroll}
                    setAutoScroll={setAutoScroll}
                    setCurrentIndex={setCurrentIndex}
                    ref={eventList}
                />
            </GlobalLayout>
        </CalendarProvider>
    );
}

const styles = createStyles({
    expandableCalendarContainer: {
        backgroundColor: 'white',
        borderBottomStartRadius: 'lg',
        borderBottomEndRadius: 'lg',
        shadowColor: 'black',
        shadowOffset: { width: 0, height: 5 },
        shadowOpacity: 0.2
    },
    expandableCalendarSubContainer: {
        borderBottomStartRadius: 'lg',
        borderBottomEndRadius: 'lg',
        backgroundColor: 'white',
        overflow: 'hidden'
    }
});
