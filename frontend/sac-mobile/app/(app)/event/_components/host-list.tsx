import { ScrollView, Text, View } from 'react-native';

import { ClubCard } from '@/components/club-card';
import { useEventHosts } from '@/hooks/use-event';
import { uuid } from '@/types/uuid';

const HostList = ({ eventID }: { eventID: uuid }) => {
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
        <>
            <ScrollView
                className="w-full"
                horizontal={true}
                showsHorizontalScrollIndicator={false}
            >
                <View className="flex-row justify-between py-2 space-x-5">
                    {hosts.map((host) => (
                        <View key={host.id}>
                            <ClubCard club={host} />
                        </View>
                    ))}
                </View>
            </ScrollView>
        </>
    );
};

export { HostList };
