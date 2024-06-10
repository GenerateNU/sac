import { useCallback, useContext, useEffect, useRef, useState } from 'react';
import { Animated, Dimensions, Easing } from 'react-native';
import { CalendarContext } from 'react-native-calendars-sac';
import { UpdateSources } from 'react-native-calendars-sac/src/expandableCalendar/commons';
import {
    GestureEventPayload,
    PanGestureHandler,
    PanGestureHandlerEventPayload,
    ScrollView,
    State
} from 'react-native-gesture-handler';

import { Box } from '@/app/(design-system)/components/Box/Box';
import Day from '@/app/(design-system)/components/Calendar/Day';
import { EventSection } from '@/app/(design-system)/components/Calendar/DayTimeSection';
import { EventCardCalendarSkeleton } from '@/app/(design-system)/components/EventCard/Skeletons/EventCardCalendarSkeleton';
import { Spacing } from '@/app/(design-system)/shared/spacing';

type SwipeDirection = 'left' | 'right';

type FetchNewEventsProps = {
    direction: 'base';
    date: string;
};
type FetchNewEventsPropsWithoutDate = {
    direction: 'start' | 'end';
};

export type FetchNewEventsPropsUnion =
    | FetchNewEventsProps
    | FetchNewEventsPropsWithoutDate;
export type RefreshDirection = 'start' | 'end' | 'base';

export const DAY_EPOCH_TIME = 1000 * 60 * 60 * 24;

type InfiniteScrollCarouselProps = {
    // The events currently loaded waiting to be rendered / currently being rendered
    loadedEvents: EventSection[];
    // The number of events to load at a time
    loadBatchSize: number;
    // Whether the component is currently loading a full batch of new events
    isLoading: boolean;
    // Refresh threshhold for loading new events: a number between 0 and 1:
    refreshThreshhold: number;
    // Callback function to intiate on threshold reach:
    refetchEvents: (props: FetchNewEventsPropsUnion) => Promise<void>;
    // Check to see if the API call has errored:
    error?: boolean;
    // State variable to track whether a date was pressed or not:
    datePress: boolean;
    // Function to set the date press state:
    setDatePress: React.Dispatch<React.SetStateAction<boolean>>;
};

export default function Calendar({
    loadedEvents,
    loadBatchSize,
    isLoading,
    refreshThreshhold,
    refetchEvents,
    error,
    datePress,
    setDatePress
}: InfiniteScrollCarouselProps) {
    // Ensure that the load batch size is at least 3:
    if (loadBatchSize < 3) {
        throw new Error('Load batch size must be at least 3');
    }

    // Extract the width of the view:
    const { width } = Dimensions.get('window');

    // Keep track of the index being viewed:
    const main_index = useRef<number>(1);

    // Keep track of the index of the data being viewed:
    const data_index_array = useRef<number[]>([
        Math.floor(loadBatchSize / 2) - 1,
        Math.floor(loadBatchSize / 2),
        Math.floor(loadBatchSize / 2) + 1
    ]);

    // Keep track of the animated translation of each of the views:
    const translateX_array = useRef<Animated.Value[]>([
        new Animated.Value(-width),
        new Animated.Value(0),
        new Animated.Value(width)
    ]);

    // Keep track of the translations in number form for use in logical operations:
    const record_translateX_array = useRef<number[]>([-width, 0, width]);

    // Keep track of the old translation value:
    const old_translateX = useRef<number>(0);

    // Keep track of the number of events loaded on most recent batch:
    const old_events = useRef<EventSection[]>([]);

    const recent_refresh_direction = useRef<RefreshDirection>('end');

    // The number of items away from the edge that will trigger a refetch:
    const refresh_threshold_item_count = Math.floor(
        Math.floor(loadBatchSize / 2) * Math.min(Math.abs(refreshThreshhold), 1)
    );

    // Current events being rendered:
    const [events, setEvents] = useState<EventSection[]>([]);

    // Keep track if the page has been initially loaded:
    const [initialLoadComplete, setInitialLoadComplete] = useState(false);

    /**
     * CODE FOR CONNECTING COMPONENT TO CALENDAR
     * - Is entirely dependent on the react-native-calendars dep
     **/
    const { date, setDate } = useContext(CalendarContext);

    // Trigger a refresh of events:
    const triggerRefresh = useCallback(
        async (refreshDirection: RefreshDirection) => {
            recent_refresh_direction.current = refreshDirection;
            if (refreshDirection === 'base') {
                await refetchEvents({
                    direction: refreshDirection,
                    date: date
                });
            } else {
                await refetchEvents({ direction: refreshDirection });
            }
        },
        [refetchEvents, date]
    );

    // Handle swipe gestures:
    const handleGestureEvent = Animated.event([{}], {
        listener: ({
            nativeEvent
        }: {
            nativeEvent: Readonly<
                GestureEventPayload & PanGestureHandlerEventPayload
            >;
        }) => {
            // Record the translation of each view:
            record_translateX_array.current.forEach((_, index, arr) => {
                arr[index] += nativeEvent.translationX - old_translateX.current;
            });

            // Update the animated values:
            translateX_array.current.forEach((translateX, index) => {
                translateX.setValue(record_translateX_array.current[index]);
            });

            old_translateX.current = nativeEvent.translationX;
        },
        useNativeDriver: false
    });

    const handleCalendarDateSelect = useCallback(
        async (
            allLoadedEvents: EventSection[],
            newSelectedDateString: string,
            dataIndexArray: number[],
            mainIndex: number
        ) => {
            // If there are no events loaded:
            if (allLoadedEvents.length === 0) {
                return;
            }

            const earliestDate = new Date(allLoadedEvents[0].date);
            const latestDate = new Date(
                allLoadedEvents[loadedEvents.length - 1].date
            );
            const currentDate = new Date(
                allLoadedEvents[dataIndexArray[mainIndex]].date
            );
            const newSelectedDate = new Date(newSelectedDateString);

            // If the calendar is currently fetching and the last fetch was an initial load:
            if (isLoading && recent_refresh_direction.current === 'base') {
                await triggerRefresh('base');
                return;
            }

            // Determine the direction of the click:
            const direction =
                newSelectedDate.getTime() > currentDate.getTime()
                    ? 'end'
                    : 'start';

            // Compute how many indices the select jump is - the number of days between the current day and the selected day:
            const daysBetween = Math.floor(
                (newSelectedDate.getTime() - currentDate.getTime()) /
                    DAY_EPOCH_TIME
            );

            // If date is in the range of the loaded dates:
            if (
                newSelectedDate.getTime() >= earliestDate.getTime() &&
                newSelectedDate.getTime() <= latestDate.getTime()
            ) {
                // Update indices:
                data_index_array.current = dataIndexArray.map(
                    (index) => index + daysBetween
                );

                // Compute how far away the new index is from the end of the list:
                const distanceFromListEnd =
                    direction === 'end'
                        ? allLoadedEvents.length -
                          data_index_array.current[mainIndex]
                        : data_index_array.current[mainIndex];

                // Update the events in view:
                updateEventsToCurrent(allLoadedEvents, 'base');

                // Fetch more dates in the direction of click:
                if (distanceFromListEnd <= refresh_threshold_item_count) {
                    await triggerRefresh(direction);
                }
            }
            // If the date is outside the range of the loaded dates:
            else {
                // Reset initial load:
                setInitialLoadComplete(false);

                const indexes = [
                    Math.floor(loadBatchSize / 2) - 1,
                    Math.floor(loadBatchSize / 2),
                    Math.floor(loadBatchSize / 2) + 1
                ];

                // Reset the data index array:
                data_index_array.current = [
                    main_index.current === 1
                        ? indexes[0]
                        : main_index.current === 0
                          ? indexes[1]
                          : indexes[2],
                    main_index.current === 1
                        ? indexes[1]
                        : main_index.current === 0
                          ? indexes[2]
                          : indexes[0],
                    main_index.current === 1
                        ? indexes[2]
                        : main_index.current === 0
                          ? indexes[0]
                          : indexes[1]
                ];

                // Trigger date refresh:
                await triggerRefresh('base');
            }
        },
        [
            isLoading,
            loadBatchSize,
            loadedEvents.length,
            refresh_threshold_item_count,
            triggerRefresh
        ]
    );

    // Update the events to what is currently been viewed:
    const updateEventsToCurrent = (
        allLoadedEvents: EventSection[],
        type: SwipeDirection | 'base',
        targetIndex?: number
    ) => {
        if (type === 'base') {
            setEvents([
                allLoadedEvents[data_index_array.current[0]],
                allLoadedEvents[data_index_array.current[1]],
                allLoadedEvents[data_index_array.current[2]]
            ]);
            old_events.current = [...allLoadedEvents];
            return;
        }

        if (targetIndex === undefined) {
            throw new Error(
                'Target index must be defined for left / right swipe direction'
            );
        }

        setEvents((prevEvents) => {
            prevEvents[targetIndex] =
                allLoadedEvents[data_index_array.current[targetIndex]];
            return [...prevEvents];
        });
    };

    // Used for triggering a date change on the calendar when the user scrolls:
    const handleDateChangeOnPan = (direction: SwipeDirection) => {
        let directionMultiplier = direction === 'left' ? 1 : -1;
        const newDateTimeStamp =
            new Date(date).getTime() + DAY_EPOCH_TIME * directionMultiplier;
        setDate(
            new Date(newDateTimeStamp).toISOString().split('T')[0],
            UpdateSources.PAGE_SCROLL
        );
    };

    // Used to handle re-render on pan of react components and trigger data refetch:
    const handleCalendarDatePan = async (
        swipeDirection: SwipeDirection,
        mainIndex: number
    ) => {
        // Update the calendar to reflect the pan:
        handleDateChangeOnPan(swipeDirection);

        let targetIndex: number;
        if (swipeDirection === 'left') {
            targetIndex = mainIndex === 0 ? 2 : mainIndex - 1;
            data_index_array.current[targetIndex] += 3;

            // Update the current events in view:
            updateEventsToCurrent(loadedEvents, 'left', targetIndex);

            const distanceFromEnd =
                loadedEvents.length - data_index_array.current[targetIndex];
            if (distanceFromEnd === refresh_threshold_item_count) {
                await triggerRefresh('end');
            }
        } else {
            targetIndex = (mainIndex + 1) % 3;
            data_index_array.current[targetIndex] -= 3;

            // Update the current events in view:
            updateEventsToCurrent(loadedEvents, 'right', targetIndex);

            const distanceFromStart = data_index_array.current[targetIndex];

            if (distanceFromStart === refresh_threshold_item_count) {
                await triggerRefresh('start');
            }
        }
    };

    // Handle the pan animation:
    const handlePanAnimation = async (mainIndex: number) => {
        const currIndex = data_index_array.current[mainIndex];
        const beforeThreshold =
            currIndex === refresh_threshold_item_count + 2 ||
            currIndex ===
                loadedEvents.length - refresh_threshold_item_count - 1;

        // Card not been fully swiped:
        if (
            (record_translateX_array.current[mainIndex] < width / 3 &&
                record_translateX_array.current[mainIndex] > -width / 3) ||
            (isLoading && beforeThreshold)
        ) {
            // Set the current translateX value of main index to 0:
            record_translateX_array.current[mainIndex] = 0;

            // Set the translate value of the next index to the width of the view:
            record_translateX_array.current[(mainIndex + 1) % 3] = width;

            // Set the translate value of the previous index to the negative width of the view:
            record_translateX_array.current[
                mainIndex === 0 ? 2 : (mainIndex - 1) % 3
            ] = -width;

            // Animate the views back to their original positions:
            let animation_array = translateX_array.current.map((_, index) =>
                Animated.timing(translateX_array.current[index], {
                    toValue: record_translateX_array.current[index],
                    useNativeDriver: false,
                    duration: 200,
                    easing: Easing.out(Easing.ease)
                })
            );
            Animated.parallel(animation_array, { stopTogether: true }).start();
            return;
        }
        // Card full swipe left:
        if (record_translateX_array.current[mainIndex] < -width / 3) {
            // Set the current index translateX value to the negative width of the view:
            record_translateX_array.current[mainIndex] = -width;

            // Set the next index translateX value to 0:
            record_translateX_array.current[(mainIndex + 1) % 3] = 0;

            // Set the previous index translateX value to be the width of the view:
            record_translateX_array.current[
                mainIndex === 0 ? 2 : mainIndex - 1
            ] = width;

            // Animate the views to their new positions:
            let animation_array = translateX_array.current.map((_, index) => {
                if (index !== (mainIndex === 0 ? 2 : mainIndex - 1)) {
                    return Animated.timing(translateX_array.current[index], {
                        toValue: record_translateX_array.current[index],
                        useNativeDriver: false,
                        duration: 150,
                        easing: Easing.out(Easing.ease)
                    });
                }
            }) as Animated.CompositeAnimation[];
            Animated.parallel(animation_array, { stopTogether: true }).start(
                async () => {
                    // Update the main index:
                    main_index.current = (mainIndex + 1) % 3;

                    await handleCalendarDatePan('left', mainIndex);
                }
            );

            return;
        }
        // Card full swipe right:
        else {
            // Set the current index translateX value to the positive width of the view:
            record_translateX_array.current[mainIndex] = width;

            // Set the previous index translateX value to 0:
            record_translateX_array.current[
                mainIndex === 0 ? 2 : mainIndex - 1
            ] = 0;

            // Set the next index translateX value to be the negative width of the view:
            record_translateX_array.current[(mainIndex + 1) % 3] = -width;

            // Animate the views to their new positions:
            let animation_array = translateX_array.current.map((_, index) => {
                if (index !== (mainIndex + 1) % 3) {
                    return Animated.timing(translateX_array.current[index], {
                        toValue: record_translateX_array.current[index],
                        useNativeDriver: false,
                        duration: 150,
                        easing: Easing.out(Easing.ease)
                    });
                }
            }) as Animated.CompositeAnimation[];
            Animated.parallel(animation_array, { stopTogether: true }).start(
                async () => {
                    await handleCalendarDatePan('right', mainIndex);
                    // Update the main index:
                    main_index.current = mainIndex === 0 ? 2 : mainIndex - 1;
                }
            );
            return;
        }
    };

    const handleStateChanged = ({
        nativeEvent
    }: {
        nativeEvent: Readonly<
            GestureEventPayload & PanGestureHandlerEventPayload
        >;
    }) => {
        // If the finger has been lifted
        if (nativeEvent.state === State.END) {
            old_translateX.current = 0;
            handlePanAnimation(main_index.current);
        }
    };

    const renderLoading = () => {
        return (
            <Box flexDirection="column" gap="s" marginVertical="s">
                <EventCardCalendarSkeleton />
                <EventCardCalendarSkeleton />
                <EventCardCalendarSkeleton />
            </Box>
        );
    };

    const renderItem = (index: number) => {
        const item = events[index];
        const scrollEnabled = !item?.data || Object.keys(item.data).length > 0;

        if (item?.date) {
            return (
                <ScrollView
                    style={{ paddingHorizontal: Spacing.s, width: '100%' }}
                    showsVerticalScrollIndicator={false}
                    scrollEnabled={scrollEnabled}
                >
                    <Day section={item} error={error} />
                </ScrollView>
            );
        }

        // If we are waiting to return data:
        if (isLoading) {
            return renderLoading();
        } else {
            return (
                <ScrollView
                    style={{ paddingHorizontal: Spacing.s, width: '100%' }}
                    showsVerticalScrollIndicator={false}
                >
                    <Day section={{ date: date, data: {} }} />
                </ScrollView>
            );
        }
    };

    // useEffect for handling the initial load:
    useEffect(() => {
        if (!isLoading && !initialLoadComplete) {
            updateEventsToCurrent(loadedEvents, 'base');
            setInitialLoadComplete(true);
            old_events.current = [...loadedEvents];
        }
    }, [isLoading, initialLoadComplete, loadedEvents]);

    // useEffect for handling events added after initial load is complete:
    useEffect(() => {
        if (loadedEvents && initialLoadComplete) {
            const totalNewEvents =
                loadedEvents.length - old_events.current.length;
            if (recent_refresh_direction.current === 'start') {
                data_index_array.current = data_index_array.current.map(
                    (index) => index + totalNewEvents
                );
            }

            // Handle when the user is looking at a loading screen:
            if (
                !old_events.current[
                    data_index_array.current[main_index.current]
                ]?.date
            ) {
                updateEventsToCurrent(loadedEvents, 'base');
            }
            old_events.current = [...loadedEvents];
        }
    }, [loadedEvents, initialLoadComplete]);

    useEffect(() => {
        if (datePress) {
            handleCalendarDateSelect(
                loadedEvents,
                date,
                data_index_array.current,
                main_index.current
            );
            setDatePress(false);
        }
    }, [datePress, date, loadedEvents, setDatePress, handleCalendarDateSelect]);

    if (!initialLoadComplete) {
        return renderLoading();
    }

    return (
        <Box
            flex={1}
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
            position="relative"
        >
            <PanGestureHandler
                onGestureEvent={handleGestureEvent}
                onHandlerStateChange={handleStateChanged}
            >
                <Animated.View
                    style={{
                        width: width,
                        height: '100%',
                        justifyContent: 'center',
                        alignItems: 'center',
                        position: 'absolute',
                        transform: [
                            { translateX: translateX_array.current[0] }
                        ],
                        marginHorizontal: Spacing.s
                    }}
                >
                    {renderItem(0)}
                </Animated.View>
            </PanGestureHandler>
            <PanGestureHandler
                onGestureEvent={handleGestureEvent}
                onHandlerStateChange={handleStateChanged}
            >
                <Animated.View
                    style={{
                        width: width,
                        height: '100%',
                        justifyContent: 'center',
                        alignItems: 'center',
                        position: 'absolute',
                        transform: [
                            { translateX: translateX_array.current[1] }
                        ],
                        marginHorizontal: Spacing.s
                    }}
                >
                    {renderItem(1)}
                </Animated.View>
            </PanGestureHandler>
            <PanGestureHandler
                onGestureEvent={handleGestureEvent}
                onHandlerStateChange={handleStateChanged}
            >
                <Animated.View
                    style={{
                        width: width,
                        height: '100%',
                        justifyContent: 'center',
                        alignItems: 'center',
                        position: 'absolute',
                        transform: [
                            { translateX: translateX_array.current[2] }
                        ],
                        marginHorizontal: Spacing.s
                    }}
                >
                    {renderItem(2)}
                </Animated.View>
            </PanGestureHandler>
        </Box>
    );
}
