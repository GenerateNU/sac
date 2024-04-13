import React from 'react';
import { FlatList, Text } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { Link } from 'expo-router';

import { Club } from '@sac/lib/api/club';
import { TEvent, useEventsQuery } from '@sac/lib/api/event';

import { clubs } from '@/data/clubs';
import { useAuthStore } from '@/hooks/use-auth';

const Home = () => {
    const { user } = useAuthStore();
    const { data: events, isLoading, error } = useEventsQuery();

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        console.error(error);
        return <Text>Error</Text>;
    }

    const renderEvent = ({ item: event }: { item: TEvent }) => (
        <Link
            className="mt-10 mr-3"
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
            className="mt-10 mr-3"
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
        <SafeAreaView className="flex items-center justify-center flex-1">
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
