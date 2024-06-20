import { SafeAreaView } from 'react-native';

import { Stack, router } from 'expo-router';

import { useAppSelector } from '@/src/store/store';

import { Arrow, Box, Button, Text } from '../../design-system';
import { EventCard } from '../../design-system/components/EventCard';
import { GlobalLayout } from '../../design-system/components/GlobalLayout/GlobalLayout';

const UserEvents = () => {
    const { rsvps } = useAppSelector((state) => state.user);

    return (
        <>
            <Stack.Screen
                options={{
                    headerTitle: () => (
                        <Text variant="subheader-2">Upcoming Events</Text>
                    ),
                    headerTransparent: true,
                    headerShown: true,
                    headerLeft: () => <Arrow color="black" />
                }}
            />
            <SafeAreaView>
                <GlobalLayout>
                    <Box
                        alignItems="center"
                        flexDirection="column"
                        flex={0}
                        marginTop="xl"
                    >
                        <Box gap="s" width="100%">
                            {rsvps.map((event) => {
                                return (
                                    <EventCard
                                        key={event.id}
                                        variant="club"
                                        event={event.name}
                                        eventId={event.id}
                                        club={event.host_name}
                                        logo={event.host_logo}
                                        startTime={event.start_time}
                                        endTime={event.end_time}
                                        image={event.host_logo ?? ''}
                                    />
                                );
                            })}
                            <Box marginTop="m">
                                <Button
                                    variant="standardButton"
                                    color="darkRed"
                                    onPress={() => router.push('/app/calendar')}
                                >
                                    Explore New Events
                                </Button>
                            </Box>
                        </Box>
                    </Box>
                </GlobalLayout>
            </SafeAreaView>
        </>
    );
};

export default UserEvents;
