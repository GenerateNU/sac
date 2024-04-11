import { Fragment } from 'react';
import { Pressable, Text, View } from 'react-native';

import { useEventHosts } from '@/hooks/use-event';
import { uuid } from '@/types/uuid';
import { router } from 'expo-router';

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
        <View className="flex-row flex-wrap w-full">
            <Text className="text-base font-light tracking-tight text-zinc-900">
                Hosted by{' '}
            </Text>
            {hosts.map((host, index) => (
                <Fragment key={index}>
                    <Pressable
                        onPress={() => router.push(`/club/${host.id}`)}
                        className={'font-bold'}
                    >
                        <Text className="text-base font-bold tracking-tight text-zinc-900">
                            {host.name}
                        </Text>
                    </Pressable>
                    <Text className="text-base font-light tracking-tight text-zinc-900">
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
