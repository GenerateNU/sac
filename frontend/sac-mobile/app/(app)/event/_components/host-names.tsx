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
        <View className="flex-row flex-wrap w-full">
            <Text className="font-light tracking-tight text-black text-sm">
                Hosted by{' '}
            </Text>
            {hosts.map((host, index) => (
                <Fragment key={index}>
                    <Pressable
                        onPress={() => router.push(`/club/${host.id}`)}
                        className={'font-bold'}
                    >
                        <Text className="text-sm font-medium tracking-tight text-black">
                            {host.name}
                        </Text>
                    </Pressable>
                    <Text className="text-sm font-light tracking-tight text-gray-500">
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
