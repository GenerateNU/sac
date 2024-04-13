import React from 'react';
import { View, Text } from 'react-native';

import { Link, Stack, useLocalSearchParams } from 'expo-router';
import { useClub } from '@/hooks/use-club';

const ClubPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();

    const { data: club, isLoading, error } = useClub(id);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!club) {
        return <Text>Club not found</Text>;
    }

    return (
        <>
            <Stack.Screen options={{ title: `${club.name}` }} />
            <View className='flex items-center justify-center flex-1'>
                <Text>{club.name}</Text>
                <Text>{club.recruitment_cycle}</Text>
                <Text>{club.application_link}</Text>
            </View>
        </>
    );
};

export default ClubPage;
