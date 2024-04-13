import { FlatList, Text } from 'react-native';

import { useEventHostsByEventIDQuery } from '@sac/lib/api/event';
import { uuid } from '@sac/lib/types/uuid';

import { ClubCard } from '@/components/club-card';

const AllHosts = ({ eventID }: { eventID: uuid }) => {
    const {
        data: hosts,
        error,
        isLoading
    } = useEventHostsByEventIDQuery(eventID);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        console.error(error);
        return <Text>Error</Text>;
    }

    if (!hosts || hosts.length === 0) {
        return <Text>No hosts found</Text>;
    }

    return (
        <FlatList
            horizontal
            showsHorizontalScrollIndicator={false}
            alwaysBounceHorizontal
            data={hosts}
            renderItem={({ item }) => <ClubCard club={item} />}
            keyExtractor={(item) => item.id}
        />
    );
};

export { AllHosts };
