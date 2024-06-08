import { SafeAreaView } from 'react-native';

import InfiniteScrollCarousel, { DAY_EPOCH_TIME, FetchNewEventsPropsUnion, RefreshDirection } from '@/app/(design-system)/components/TestComponent/TestComponent';
import { useCallback, useEffect, useRef, useState } from 'react';
import { EventSection } from '@/app/(design-system)/components/Calendar/mockData';
import { getEvents } from '@/app/(design-system)/components/TestComponent/mock/mockData';
import { CalendarProvider, ExpandableCalendar } from 'react-native-calendars-sac';
import { Box, createStyles } from '@/app/(design-system)';
import { Theme } from 'react-native-calendars-sac/src/types';

const TODAY = new Date(); 

const calendarTheme: Theme = {
    selectedDayBackgroundColor: 'black',
    todayTextColor: 'black',
    arrowColor: 'black',
    textMonthFontWeight: 'bold'
};

const BATCH_SIZE = 15
const BATCH_SIZE_HALF = Math.floor(BATCH_SIZE / 2);

const CalendarPage = () => {
    const [loadedEvents, setLoadedEvents] = useState<EventSection[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [minDate, setMinDate] = useState(TODAY.getTime() - DAY_EPOCH_TIME * BATCH_SIZE_HALF);
    const [maxDate, setMaxDate] = useState(TODAY.getTime() + DAY_EPOCH_TIME * BATCH_SIZE_HALF);
    const [datePress, setDatePress] = useState(false);
    const oldEvents = useRef<EventSection[]>([]);
    const recentRefresh = useRef<string>();

    const emptyDatePlaceholder = new Array(BATCH_SIZE_HALF).fill({});

    useEffect(() => {
        fetchNewEvents({ 
            direction: 'base', 
            date: TODAY.toISOString().split('T')[0] 
        });
    }, []);

    const fetchEvents = async (startTime: number, endTime: number) => {
        const events = getEvents(startTime, endTime);
        return events;
    };

    const fetchNewEvents = useCallback(async (props: FetchNewEventsPropsUnion) => {
        // Create a random string identifier for the refresh:
        const refreshId = Math.random().toString(36).substring(7);
        recentRefresh.current = refreshId;
        console.log('refreshID:', refreshId);

        const { direction } = props;
        if (direction === 'start') {
            // Store the old events:
            oldEvents.current = [...loadedEvents];
            setLoadedEvents([...emptyDatePlaceholder, ...oldEvents.current]);

            setIsLoading(true);
            setTimeout(async () => {
                const newEvents = await fetchEvents(minDate - DAY_EPOCH_TIME * BATCH_SIZE_HALF, minDate - DAY_EPOCH_TIME);
                if (recentRefresh.current !== refreshId) return;
                setMinDate(minDate - DAY_EPOCH_TIME * BATCH_SIZE_HALF);
                setLoadedEvents([...newEvents, ...oldEvents.current]);
                setIsLoading(false)
            }, 2000)
        } 
        if (direction === 'end') {
            // Store the old events:
            oldEvents.current = [...loadedEvents];
            setLoadedEvents([...oldEvents.current, ...emptyDatePlaceholder]);

            setIsLoading(true);
            setTimeout(async () => {
                const newEvents = await fetchEvents(maxDate + DAY_EPOCH_TIME, maxDate + DAY_EPOCH_TIME * (BATCH_SIZE_HALF + 1));
                if (recentRefresh.current !== refreshId) return;
                setMaxDate(maxDate + DAY_EPOCH_TIME * (BATCH_SIZE_HALF + 1));
                setLoadedEvents([...oldEvents.current, ...newEvents]);
                setIsLoading(false);
            }, 2000);
        } 
        if (direction === 'base') {
            const { date } = props;
            const minDateNew = new Date(date).getTime() - DAY_EPOCH_TIME * BATCH_SIZE_HALF;
            const maxDateNew = new Date(date).getTime() + DAY_EPOCH_TIME * BATCH_SIZE_HALF;
            setMinDate(minDateNew);
            setMaxDate(maxDateNew);

            setIsLoading(true);
            setTimeout(async () => {
                const newEvents = await fetchEvents(minDateNew, maxDateNew);
                console.log(new Date(minDateNew), new Date(maxDateNew));
                if (recentRefresh.current !== refreshId) return;
                console.log('request not cancelled!');
                setLoadedEvents(newEvents);
                setIsLoading(false);
            }, 3000)
        }
        else return;
    }, [minDate, maxDate, loadedEvents]);

    return (
        <SafeAreaView style={{ flex: 1 }}>
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
                            closeOnDayPress={false}
                            disableWeekScroll={true}
                            onDayPress={() => setDatePress(true)}
                        />
                    </Box>
                </Box>
                <InfiniteScrollCarousel
                    loadedEvents={loadedEvents}
                    loadBatchSize={BATCH_SIZE}
                    isLoading={isLoading}
                    refreshThreshhold={0.25}
                    refetchEvents={fetchNewEvents}
                    datePress={datePress}
                    setDatePress={setDatePress}
                />
            </CalendarProvider>
        </SafeAreaView>
    );
};

export default CalendarPage;

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
