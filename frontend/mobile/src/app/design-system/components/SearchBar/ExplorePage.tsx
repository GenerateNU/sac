import React from 'react';
import {
    FlatList,
    Pressable,
    SafeAreaView,
    ScrollView,
    StyleSheet,
    View
} from 'react-native';

import { router } from 'expo-router';

import { eventApi } from '@generatesac/lib';
import { clubApi } from '@generatesac/lib';

import { Text } from '../../components/Text/Text';
import { Spacing } from '../../shared/spacing';
import { Box } from '../Box/Box';
import { EventCard } from '../EventCard';
import { GlobalLayout } from '../GlobalLayout/GlobalLayout';
import { ClubCard } from '../clubCard/ClubCard';
import SearchBar from './SearchBar';

const ExplorePage = () => {
    const numColumns = 2;
    const {
        data: eventsData,
        error: eventsError,
        isLoading: eventsLoading
    } = eventApi.useEventsQuery({});

    const {
        data: clubsData,
        error: clubsError,
        isLoading: clubsLoading
    } = clubApi.useClubsQuery({});

    const handleSearch = (query: string) => {
        console.log('Search query:', query);
    };

    const handleCancel = () => {
        console.log('Search canceled');
    };

    const handleBack = () => {
        console.log('Back pressed');
    };

    return (
        <GlobalLayout>
            <SafeAreaView style={styles.container}>
                <ScrollView showsVerticalScrollIndicator={false}>
                    <SearchBar
                        placeholder="Search"
                        onSearch={handleSearch}
                        onCancel={handleCancel}
                        onBack={handleBack}
                    />
                    <Box paddingBottom="s">
                        <Box
                            flexDirection="row"
                            justifyContent="space-between"
                            alignItems="center"
                        >
                            <Text variant="subheader-2">Recommended Clubs</Text>
                            <Pressable
                                onPress={() => router.push('/app/clubs')}
                            >
                                <Text variant="caption-1">View all</Text>
                            </Pressable>
                        </Box>
                        {clubsLoading && <Text>Loading...</Text>}
                        {clubsError && <Text>Error loading clubs.</Text>}
                        {!clubsLoading && !clubsError && (
                            <FlatList
                                data={clubsData?.slice(0, 2) || []}
                                keyExtractor={(item) => item.id}
                                horizontal={false}
                                showsHorizontalScrollIndicator={false}
                                scrollEnabled={false}
                                renderItem={({ item }) => (
                                    <Box padding="xs">
                                        <Pressable
                                            onPress={() =>
                                                router.push(
                                                    `/app/club/${item.id}`
                                                )
                                            }
                                        >
                                            <ClubCard
                                                club={item.name}
                                                // tags={item.tags}
                                                imageUrl={item.logo || ''}
                                            />
                                        </Pressable>
                                    </Box>
                                )}
                            />
                        )}
                    </Box>
                    <Box>
                        <Box
                            flexDirection="row"
                            justifyContent="space-between"
                            alignItems="center"
                        >
                            <Text variant="subheader-2">
                                Recommended Events
                            </Text>
                            <Pressable
                                onPress={() => router.push('/app/events')}
                            >
                                <Text variant="caption-1">View all</Text>
                            </Pressable>
                        </Box>
                        {eventsLoading && <Text>Loading...</Text>}
                        {eventsError && <Text>Error loading events.</Text>}
                        {!eventsLoading && !eventsError && (
                            <FlatList
                                data={eventsData?.slice(0, 4) || []}
                                keyExtractor={(item) => item.id}
                                horizontal={false}
                                key={numColumns}
                                numColumns={numColumns}
                                scrollEnabled={false}
                                renderItem={({ item }) => (
                                    <Box padding="xs">
                                        <Pressable
                                            onPress={() =>
                                                router.push(
                                                    `/app/event/${item.id}`
                                                )
                                            }
                                        >
                                            <EventCard
                                                variant="small"
                                                event={item.name}
                                                club={item.host}
                                                eventId={item.id}
                                                startTime={item.start_time}
                                                endTime={item.end_time}
                                                image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLF3ord7lnV_5Je-pC2AUgUiesHNPcZlpI7A&s"
                                                logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                                            />
                                        </Pressable>
                                    </Box>
                                )}
                            />
                        )}
                    </Box>
                </ScrollView>
            </SafeAreaView>
        </GlobalLayout>
    );
};

export default ExplorePage;

const styles = StyleSheet.create({
    container: {
        // flex: 1,
        // flexDirection: 'column',
        // alignItems: 'flex-start',
        // justifyContent: 'flex-start',
        // marginVertical: Spacing.m,
        // gap: Spacing.l
    }
});
