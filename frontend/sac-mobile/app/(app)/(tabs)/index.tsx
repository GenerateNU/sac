import React from 'react';
import { Text, View, FlatList } from 'react-native';

import { Link } from 'expo-router';

import { useAuthStore } from '@/hooks/use-auth';
import { useEvents } from '@/hooks/use-event';
import { Event } from '@/types/event';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Club } from '@/types/club';
import { clubs } from '@/data/clubs';

const Home = () => {
    const { user } = useAuthStore();
    const { data: events, isLoading, error } = useEvents();

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    const renderEvent = ({ item: event }: { item: Event }) => (
        <Link
            className='mt-10 mr-3'
            key={event.id}
            href={{
                pathname: '/(app)/event/[id]',
                params: { id: event.id }
            }}
        >
            <Text>{event.name}</Text>
        </Link>
    );

    const renderClub = ({ item: club }: { item: Club }) => (
        <Link
            className='mt-10 mr-3'
            key={club.id}
            href={{
                pathname: '/(app)/club/[id]',
                params: { id: club.id }
            }}
        >
            <Text>{club.name}</Text>
        </Link>
    );

    return (
        <SafeAreaView className='flex items-center justify-center flex-1'>
            <Text>Welcome {user?.first_name}</Text>
            <FlatList
                numColumns={5}
                data={events}
                renderItem={renderEvent}
                keyExtractor={(item) => item.id.toString()}
            />
            <FlatList
                numColumns={5}
                data={clubs}
                renderItem={renderClub}
                keyExtractor={(item) => item.id.toString()}
            />
        </SafeAreaView>
    );
};

export default Home;
