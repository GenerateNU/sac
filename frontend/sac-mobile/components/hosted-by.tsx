import { ScrollView, Text, View } from 'react-native';

import { useHosts } from '@/hooks/use-hosts';
import { uuid } from '@/types/uuid';

import ClubCard from './club-card';

const HostedBy = ({ eventID }: { eventID: uuid }) => {
    const { data: hosts, error, isLoading } = useHosts(eventID);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!hosts || hosts.length === 0) {
        return <Text>No hosts found</Text>;
    }
    console.log(hosts);

    return (
        <>
            <Text className="py-2 text-2xl font-bold">Hosted by</Text>
            <ScrollView
                className="w-full"
                horizontal={true}
                showsHorizontalScrollIndicator={false}
            >
                <View className="flex-row justify-between space-x-5">
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

export default HostedBy;
