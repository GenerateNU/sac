import React from 'react';
import { Text, View } from 'react-native';

import { useEvent } from '@/hooks/use-event';

const Event = () => {
    const { data: event, isLoading, error } = useEvent('tester');

    // TODO: Handle error once we have error components
    // TODO: Handle loading state once we have loading components

    return (
        <>
            {event && (
                <View className="flex-1 justify-center items-center">
                    <Text>{event.name}</Text>
                    <Text>{event.content}</Text>
                </View>
            )}
        </>
    );
};

export default Event;
