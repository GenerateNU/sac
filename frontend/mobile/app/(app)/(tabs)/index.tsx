import { SafeAreaView, StyleSheet, Text } from 'react-native';

import { GlobalLayout } from '@/app/(design-system)/components/GlobalLayout/GlobalLayout';
import TestComponent from '@/app/(design-system)/components/TestComponent/TestComponent';
import { getEvents } from '@/app/(design-system)/components/TestComponent/mock/mockData';
import { EventSection } from '@/app/(design-system)/components/Calendar/mockData';
import { useCallback, useEffect, useState } from 'react';

const TODAY = new Date(); 
const DAY_EPOCH_TIME = 1000 * 60 * 60 * 24;

const HomePage = () => {
    const [loadedEvents, setLoadedEvents] = useState<EventSection[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [minDate, setMinDate] = useState(TODAY.getTime() - DAY_EPOCH_TIME * 7);
    const [maxDate, setMaxDate] = useState(TODAY.getTime() + DAY_EPOCH_TIME * 7);

    useEffect(() => {
        setTimeout(() => {
            fetchEvents(minDate, maxDate)
                .then((events) => {
                    setLoadedEvents(events);
                    setIsLoading(false);
                });
        }, 2000)
    }, []);

    const fetchEvents = async (startTime: number, endTime: number) => {
        const events = getEvents(startTime, endTime);
        console.log('new events fetched!')
        return events;
    };

    const fetchNewEvents = useCallback(async (direction: 'start' | 'end') => {
        if (direction === 'start') {
            console.log('fetching new events start!')
            const newEvents = await fetchEvents(minDate - DAY_EPOCH_TIME * 7, minDate);
            setMinDate(minDate - DAY_EPOCH_TIME * 7);
            setLoadedEvents((oldEvents) => [...newEvents, ...oldEvents]);
        } else {
            console.log('fetching new events end!')
            const newEvents = await fetchEvents(maxDate + DAY_EPOCH_TIME, maxDate + DAY_EPOCH_TIME * 8);
            setMaxDate(maxDate + DAY_EPOCH_TIME * 8);
            setLoadedEvents((oldEvents) => [...oldEvents, ...newEvents]);
        }
    }, [minDate, maxDate, loadedEvents]);

    return (
        <GlobalLayout>
            <SafeAreaView style={styles.container}>
                <Text style={styles.title}>Home</Text>
                <TestComponent
                    loadedEvents={loadedEvents}
                    loadBatchSize={15}
                    isLoading={isLoading}
                    refreshThreshhold={0.25}
                    onRefreshThreasholdReach={fetchNewEvents}
                />
                {/* <ScrollView showsVerticalScrollIndicator={false}>
                    <EventCard
                        event="Spring Showcase"
                        variant="small"
                        tags={tags}
                        club="Generate Product Development Studio"
                        startTime={new Date()}
                        endTime={new Date()}
                        logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                    />
                    <EventCard
                        event="Spring Showcase"
                        variant="big"
                        tags={tags}
                        club="Generate Product Development Studio"
                        startTime={new Date()}
                        endTime={new Date()}
                        logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                    />
                    <EventCard
                        event="Spring Showcase"
                        variant="club"
                        tags={tags}
                        club="Generate"
                        startTime={new Date()}
                        endTime={new Date()}
                        logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                    />
                    <EventCard
                        event="Spring Showcase"
                        variant="calendar"
                        tags={tags}
                        club="Generate"
                        startTime={new Date()}
                        endTime={date}
                        logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                    />
                </ScrollView> */}
            </SafeAreaView>
        </GlobalLayout>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center'
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold'
    },
    separator: {
        marginVertical: 30,
        height: 1,
        width: '80%'
    }
});

export default HomePage;
