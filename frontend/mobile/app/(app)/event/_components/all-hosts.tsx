import { FlatList, Text } from 'react-native';

import { ClubCard } from '@/components/club-card';
import { useEventHosts } from '@/hooks/use-event';
import { uuid } from '@/types/uuid';

const AllHosts = ({ eventID }: { eventID: uuid }) => {
    const { data: hosts, error, isLoading } = useEventHosts(eventID);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
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
