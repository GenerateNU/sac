import React from 'react';
import { Text, View } from 'react-native';
import { FlatList } from 'react-native-gesture-handler';

import { useClubs } from '@/hooks/use-club';
import { Club } from '@/types/club';

import { ClubHomePageCard } from './club-homepage-card';

const RecommendedClubs = () => {
    const { data: clubs, isLoading, error } = useClubs();

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!clubs) {
        return <Text>No clubs found</Text>;
    }

    const renderClubs = ({ item }: { item: Club }) => {
        return <ClubHomePageCard club={item} />;
    };

    return (
        <View style={{ paddingLeft: 20, flex: 1 }}>
            <FlatList
                data={clubs}
                renderItem={renderClubs}
                keyExtractor={(club) => club.id}
                horizontal
                showsHorizontalScrollIndicator={false}
                contentContainerStyle={{ paddingRight: 20, paddingBottom: 20 }}
            />
        </View>
    );
};

export { RecommendedClubs };
