import React from 'react';
import { Pressable, StyleSheet } from 'react-native';

import { router } from 'expo-router';

import { Box, Text } from '@/src/app/(design-system)';
import { EventCard } from '@/src/app/(design-system)/components/EventCard';

const HomePage = () => {
    const item = {
        name: 'Your Event Name',
        host: 'Your Club Name',
        start_time: new Date(),
        end_time: new Date()
    };

    return (
        <Box style={styles.container}>
            <Text variant="body-1">Home</Text>
            <Pressable onPress={() => router.push(`/event/1`)}>
                <EventCard
                    variant="small"
                    event={item.name}
                    club={item.host}
                    clubId="1"
                    startTime={item.start_time}
                    endTime={item.end_time}
                    image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLF3ord7lnV_5Je-pC2AUgUiesHNPcZlpI7A&s"
                />
            </Pressable>
        </Box>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center'
    },
    contentContainer: {
        flex: 1,
        alignItems: 'center'
    }
});

export default HomePage;
