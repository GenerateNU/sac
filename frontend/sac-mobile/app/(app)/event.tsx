import React from 'react';
import { Text, View } from 'react-native';

import { useEvent } from '@/hooks/use-event';

const Event = () => {
    const {
        data: event,
        isLoading,
        error
    } = useEvent('8a3b3e3e-3e3e-3e3e-3e3e-3e3e3e3e3e3e');

    if (error)
        return (
            <View className="flex-1 justify-center items-center">
                <Text>Error: {error.message}</Text>
            </View>
        );

    return (
        <>
            {isLoading && <Text>Loading...</Text>}
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
