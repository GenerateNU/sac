import { Text, View } from 'react-native';

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
        <View>
            <View className="inline-flex flex-col justify-between">
                <Text className="text-xl font-light leading-relaxed text-zinc-900">
                    Hosted by
                </Text>
            </View>
            <View className="mt-2">
                {hosts.map((host) => {
                    return <ClubCard key={host.id} club={host} />;
                })}
            </View>
        </View>
    );
};

export { AllHosts };
