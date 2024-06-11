import { useCallback, useEffect, useRef, useState } from 'react';
import { SafeAreaView } from 'react-native';
import {
    CalendarProvider,
    ExpandableCalendar
} from 'react-native-calendars-sac';
import { Theme } from 'react-native-calendars-sac/src/types';

import { eventApi } from '@generatesac/lib';

import { Box, Colors, createStyles } from '@/src/app/(design-system)';
import Calendar, {
    DAY_EPOCH_TIME,
    FetchNewEventsPropsUnion
} from '@/src/app/(design-system)/components/Calendar/Calendar';
import { EventSection } from '@/src/app/(design-system)/components/Calendar/DayTimeSection';
import { parseData } from '@/src/app/(design-system)/components/Calendar/parser/calendarParser';

const TODAY = new Date();

const BATCH_SIZE = 15;
const BATCH_SIZE_HALF = Math.floor(BATCH_SIZE / 2);

const CalendarPage = () => {
    const [loadedEvents, setLoadedEvents] = useState<EventSection[]>([]);
    const [minDate, setMinDate] = useState(
        TODAY.getTime() - DAY_EPOCH_TIME * BATCH_SIZE_HALF
    );
    const [maxDate, setMaxDate] = useState(
        TODAY.getTime() + DAY_EPOCH_TIME * BATCH_SIZE_HALF
    );
    const [datePress, setDatePress] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(false);
    const oldEvents = useRef<EventSection[]>([]);
    const recentRefresh = useRef<string>();
    const [getEvents] = eventApi.useLazyEventsQuery({});

    const emptyDatePlaceholder = new Array(BATCH_SIZE_HALF).fill({});

    const calendarTheme: Theme = {
        selectedDayBackgroundColor: Colors.darkRed,
        todayTextColor: 'black',
        arrowColor: 'black',
        textMonthFontWeight: 'bold'
    };

    const fetchEvents = useCallback(
        async (startTime: number, endTime: number) => {
            const start = new Date(startTime).toISOString().split('T')[0];
            const end = new Date(endTime).toISOString().split('T')[0];

            return await getEvents({ start, end }).then(
                ({ data, error: fetchError }) => {
                    if (fetchError) {
                        setError(true);
                    } else {
                        setError(false);
                    }

                    return parseData(startTime, endTime, data);
                }
            );
        },
        [getEvents]
    );

    const fetchNewEvents = useCallback(
        async (props: FetchNewEventsPropsUnion) => {
            const { direction } = props;
            const refreshId = Math.random().toString(36).substring(7);
            recentRefresh.current = refreshId;

            if (direction === 'start') {
                // Store the old events:
                oldEvents.current = [...loadedEvents];
                setLoadedEvents([
                    ...emptyDatePlaceholder,
                    ...oldEvents.current
                ]);

                setIsLoading(true);
                setTimeout(async () => {
                    const newEvents = await fetchEvents(
                        minDate - DAY_EPOCH_TIME * BATCH_SIZE_HALF,
                        minDate
                    );
                    if (recentRefresh.current !== refreshId) return;
                    setMinDate(minDate - DAY_EPOCH_TIME * BATCH_SIZE_HALF);
                    setLoadedEvents([...newEvents, ...oldEvents.current]);
                    setIsLoading(false);
                }, 200);
            }
            if (direction === 'end') {
                // Store the old events:
                oldEvents.current = [...loadedEvents];
                setLoadedEvents([
                    ...oldEvents.current,
                    ...emptyDatePlaceholder
                ]);

                setIsLoading(true);
                setTimeout(async () => {
                    const newEvents = await fetchEvents(
                        maxDate,
                        maxDate + DAY_EPOCH_TIME * (BATCH_SIZE_HALF + 1)
                    );
                    if (recentRefresh.current !== refreshId) return;
                    setMaxDate(
                        maxDate + DAY_EPOCH_TIME * (BATCH_SIZE_HALF + 1)
                    );
                    setLoadedEvents([...oldEvents.current, ...newEvents]);
                    setIsLoading(false);
                }, 200);
            }
            if (direction === 'base') {
                const { date } = props;
                const minDateNew =
                    new Date(date).getTime() - DAY_EPOCH_TIME * BATCH_SIZE_HALF;
                const maxDateNew =
                    new Date(date).getTime() + DAY_EPOCH_TIME * BATCH_SIZE_HALF;
                setMinDate(minDateNew);
                setMaxDate(maxDateNew);

                setIsLoading(true);
                setTimeout(async () => {
                    const newEvents = await fetchEvents(minDateNew, maxDateNew);
                    if (recentRefresh.current !== refreshId) return;
                    setLoadedEvents(newEvents);
                    setIsLoading(false);
                }, 200);
            } else return;
        },
        [minDate, maxDate, loadedEvents, emptyDatePlaceholder, fetchEvents]
    );

    useEffect(() => {
        fetchNewEvents({
            direction: 'base',
            date: TODAY.toISOString().split('T')[0]
        });
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

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
                <Calendar
                    loadedEvents={loadedEvents}
                    loadBatchSize={BATCH_SIZE}
                    isLoading={isLoading}
                    refreshThreshhold={0.25}
                    refetchEvents={fetchNewEvents}
                    datePress={datePress}
                    setDatePress={setDatePress}
                    error={error}
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
