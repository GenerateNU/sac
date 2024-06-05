import { Dimensions, Animated, Easing } from "react-native";
import { useEffect, useRef, useState } from "react";
import { GestureEventPayload, PanGestureHandler, PanGestureHandlerEventPayload, ScrollView, State } from "react-native-gesture-handler";
import { Box } from "../Box/Box";
import Day from "./Day";
import { Spacing } from "../../shared/spacing";
import { EventSection } from "../Calendar/mockData";
import { EventCardCalendarSkeleton } from "../EventCard/Skeletons/EventCardCalendarSkeleton";

type SwipeDirection = 'left' | 'right';

type TestComponentProps = {
    // The events currently loaded waiting to be rendered / currently being rendered
    loadedEvents: EventSection[];
    // The number of events to load at a time
    loadBatchSize: number,
    // Whether the component is currently loading a full batch of new events
    isLoading: boolean,
    // Refresh threshhold for loading new events: a number between 0 and 1:
    refreshThreshhold: number,
    // Callback function to intiate on threshold reach:
    onRefreshThreasholdReach: (direction: 'start' | 'end') => void,
}

export default function TestComponent({ 
    loadedEvents,
    loadBatchSize,
    isLoading,
    refreshThreshhold,
    onRefreshThreasholdReach,
}: TestComponentProps) {
    // Ensure that the loaded events are at least 3:
    if (loadedEvents.length < 3 && !isLoading) {
        throw new Error('Loaded events must be at least 3');
    }
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
        new Animated.Value(width),
    ]);

    // Keep track of the translations in number form for use in logical operations:
    const record_translateX_array = useRef<number[]>([-width, 0, width])

    // Keep track of the old translation value:
    const old_translateX = useRef<number>(0)

    // Keep track of the number of events loaded on most recent batch:
    const old_events = useRef<EventSection[]>([]);

    const recent_refresh_direction = useRef<'start' | 'end'>('end');

    // The number of items away from the edge that will trigger a refetch:
    const refresh_threshold_item_count = Math.floor(Math.floor(loadBatchSize / 2) * Math.min(Math.abs(refreshThreshhold), 1));

    // Current events being rendered:
    const [events, setEvents] = useState<EventSection[]>([])
    
    // Keep track if the page has been initially loaded:
    const [initialLoadComplete, setInitialLoadComplete] = useState(false);

    useEffect(() => {
        if (!isLoading && !initialLoadComplete) {
            setEvents([
                loadedEvents[data_index_array.current[0]],
                loadedEvents[data_index_array.current[1]],
                loadedEvents[data_index_array.current[2]]
            ])
            setInitialLoadComplete(true);
            old_events.current = [...loadedEvents]
            console.log('event total now at: ', loadedEvents.length);
        }
    }, [isLoading]);

    useEffect(() => {
        if (loadedEvents && initialLoadComplete) {
            const totalNewEvents = loadedEvents.length - old_events.current.length;
            if (recent_refresh_direction.current === 'start') {
                data_index_array.current = data_index_array.current.map((index) => index + totalNewEvents)
            }
            old_events.current = [...loadedEvents]
            console.log('event total now at: ', loadedEvents.length);

        }
    }, [loadedEvents])

    // Handle swipe gestures: 
    const handleGestureEvent = Animated.event(
        [{}],
        {
            listener: ({ nativeEvent }: { nativeEvent: Readonly<GestureEventPayload & PanGestureHandlerEventPayload> }) => {
                // Record the translation of each view:
                record_translateX_array.current.forEach((_, index, arr) => {
                    arr[index] += (nativeEvent.translationX - old_translateX.current)
                })

                // Update the animated values:
                translateX_array.current.forEach((translateX, index) => {
                    translateX.setValue(record_translateX_array.current[index])
                })

                old_translateX.current = nativeEvent.translationX
            },
            useNativeDriver: false,
        }
    );

    const handleEventChange = (swipeDirection: SwipeDirection, mainIndex: number) => {
        let targetIndex: number;
        if (swipeDirection === 'left') {
            targetIndex = mainIndex === 0 ? 2 : mainIndex - 1
            data_index_array.current[targetIndex] += 3;
            setEvents((prevEvents) => {
                prevEvents[targetIndex] = loadedEvents[data_index_array.current[targetIndex]]
                return [...prevEvents]
            })
            const distanceFromEnd = loadedEvents.length - data_index_array.current[targetIndex]
            if (distanceFromEnd === refresh_threshold_item_count) {
                recent_refresh_direction.current = 'end';
                onRefreshThreasholdReach('end')
            }
        }
        else {
            targetIndex = (mainIndex + 1) % 3
            data_index_array.current[targetIndex] -= 3;
            setEvents((prevEvents) => {
                prevEvents[targetIndex] = loadedEvents[data_index_array.current[targetIndex]]
                return [...prevEvents]
            })

            const distanceFromStart = data_index_array.current[targetIndex]
            if (distanceFromStart === refresh_threshold_item_count) {
                recent_refresh_direction.current = 'start';
                onRefreshThreasholdReach('start')
            }
        }
    }

    // Handle the pan animation:
    const handlePanAnimation = (mainIndex: number, panVelocity: number) => {
        // Card not been fully swiped:
        if (record_translateX_array.current[mainIndex] < width / 2 && record_translateX_array.current[mainIndex] > -width / 2) {
            // Set the current translateX value of main index to 0:
            record_translateX_array.current[mainIndex] = 0;

            // Set the translate value of the next index to the width of the view:
            record_translateX_array.current[(mainIndex + 1) % 3] = width;

            // Set the translate value of the previous index to the negative width of the view:
            record_translateX_array.current[mainIndex === 0 ? 2 : (mainIndex - 1) % 3] = -width;

            // Animate the views back to their original positions:
            let animation_array = translateX_array.current.map((_, mainIndex) =>
                Animated.timing(translateX_array.current[mainIndex], {
                    toValue: record_translateX_array.current[mainIndex],
                    useNativeDriver: false,
                    duration: 200,
                    easing: Easing.out(Easing.ease)
                })
            )
            Animated.parallel(animation_array, { stopTogether: true }).start()
            return;
        }
        // Card full swipe left:
        if (record_translateX_array.current[mainIndex] < -width / 2) {
            // Set the current index translateX value to the negative width of the view:
            record_translateX_array.current[mainIndex] = -width;

            // Set the next index translateX value to 0:
            record_translateX_array.current[(mainIndex + 1) % 3] = 0;

            // Set the previous index translateX value to be the width of the view:
            record_translateX_array.current[mainIndex === 0 ? 2 : mainIndex - 1] = width;

            // Animate the views to their new positions:
            let animation_array = translateX_array.current.map((_, index) => {
                if (index !== (mainIndex === 0 ? 2 : mainIndex - 1)) {
                    return Animated.timing(translateX_array.current[index], {
                        toValue: record_translateX_array.current[index],
                        useNativeDriver: false,
                        duration: 150,
                        easing: Easing.out(Easing.ease)
                    })
                }
            }) as Animated.CompositeAnimation[];
            Animated.parallel(animation_array, { stopTogether: true }).start(() => {
                handleEventChange('left', mainIndex)
                // Update the main index:
                main_index.current = (mainIndex + 1) % 3;
            })

            return;
        }
        // Card full swipe right:
        else {
            // Set the current index translateX value to the positive width of the view:
            record_translateX_array.current[mainIndex] = width;

            // Set the previous index translateX value to 0:
            record_translateX_array.current[mainIndex === 0 ? 2 : mainIndex - 1] = 0;

            // Set the next index translateX value to be the negative width of the view:
            record_translateX_array.current[(mainIndex + 1) % 3] = -width;

            // Animate the views to their new positions:
            let animation_array = translateX_array.current.map((_, index) => {
                if (index !== ((mainIndex + 1) % 3)) {
                    return Animated.timing(translateX_array.current[index], {
                        toValue: record_translateX_array.current[index],
                        useNativeDriver: false,
                        duration: 150,
                        easing: Easing.out(Easing.ease)
                    })
                }
            }) as Animated.CompositeAnimation[];
            Animated.parallel(animation_array, { stopTogether: true }).start(() => {
                handleEventChange('right', mainIndex)
                // Update the main index:
                main_index.current = mainIndex === 0 ? 2 : mainIndex - 1;
            })
            return;
        }
    }

    const handleStateChanged = ({ nativeEvent }: { nativeEvent: Readonly<GestureEventPayload & PanGestureHandlerEventPayload> }) => {
        // If the finger has been lifted
        if (nativeEvent.state === State.END) {
            old_translateX.current = 0;
            handlePanAnimation(
                main_index.current,
                nativeEvent.velocityX,
            );
        }
    }

    const renderItem = (index: number) => {
        return (
            <ScrollView 
                style={{ marginHorizontal: Spacing.s }}
                showsVerticalScrollIndicator={false}
            >
                <Day section={events[index]} />
            </ScrollView>
        )
    }

    if (!initialLoadComplete) {
        return (
            <Box flexDirection="column" gap="s" marginVertical="s">
                <EventCardCalendarSkeleton />
                <EventCardCalendarSkeleton />
                <EventCardCalendarSkeleton />
            </Box>
        )
    }

    return (
        <Box 
            flex={1}
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
            position='relative'
        >
            <PanGestureHandler
                onGestureEvent={handleGestureEvent}
                onHandlerStateChange={handleStateChanged}
            >
                <Animated.View 
                    style={{
                        width: width,
                        height: '100%',
                        justifyContent: "center",
                        alignItems: "center",
                        position: 'absolute',
                        transform: [{ translateX: translateX_array.current[0] }],
                        marginHorizontal: Spacing.s,
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
                        justifyContent: "center",
                        alignItems: "center",
                        position: 'absolute',
                        transform: [{ translateX: translateX_array.current[1] }],
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
                        justifyContent: "center",
                        alignItems: "center",
                        position: 'absolute',
                        transform: [{ translateX: translateX_array.current[2] }],
                        marginHorizontal: Spacing.s,
                    }}
                >
                     {renderItem(2)}
                </Animated.View>
            </PanGestureHandler>
        </Box>
    )
}