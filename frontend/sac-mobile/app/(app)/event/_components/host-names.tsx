import { Fragment } from 'react';
import { Pressable, Text, View } from 'react-native';

import { router } from 'expo-router';

import { useEventHosts } from '@/hooks/use-event';
import { uuid } from '@/types/uuid';

const HostNames = ({ eventID }: { eventID: uuid }) => {
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
        <View style={{ flexDirection: 'row', flexWrap: 'wrap', width: '100%' }}>
            <Text style={{ fontSize: 14, fontWeight: '300', color: 'black' }}>
                Hosted by{' '}
            </Text>
            {hosts.map((host, index) => (
                <Fragment key={index}>
                    <Pressable
                        onPress={() => router.push(`/club/${host.id}`)}
                        style={{ flexDirection: 'row', flexWrap: 'wrap' }}
                    >
                        <Text style={{ fontSize: 14, fontWeight: 'bold', color: 'black' }}>
                            {host.name}
                        </Text>
                    </Pressable>
                    <Text style={{ fontSize: 14, fontWeight: '300', color: 'gray' }}>
                        {index < hosts.length - 2
                            ? ', '
                            : index === hosts.length - 2
                              ? ' & '
                              : ''}
                    </Text>
                </Fragment>
            ))}
        </View>
    );
};

export { HostNames };
