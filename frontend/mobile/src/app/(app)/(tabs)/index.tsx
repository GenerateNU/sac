import React, { useRef, useState } from 'react';
import { StyleSheet } from 'react-native';

import BottomSheet from '@gorhom/bottom-sheet';

import { ClubPreview, EventPreview } from '@/src/app/(design-system)';
import { Box, Button, Text } from '@/src/app/(design-system)';

const HomePage = () => {
    const club = useRef<BottomSheet>(null);
    const event = useRef<BottomSheet>(null);

    const [clubPreview, setClubPreview] = useState('');
    const [eventPreview, setEventPreview] = useState('');

    return (
        <Box style={styles.container}>
            <Text variant="body-1">Home</Text>
            <Button
                onPress={() => {
                    setClubPreview('1');
                    club.current?.snapToIndex(0);
                }}
                variant="standardButton"
            >
                Club Preview
            </Button>
            <Button
                onPress={() => {
                    setEventPreview('1');
                    event.current?.snapToIndex(0);
                }}
                variant="standardButton"
            >
                Event Preview
            </Button>
            <EventPreview eventId={eventPreview} ref={event} />
            <ClubPreview clubId={clubPreview} ref={club} />
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
