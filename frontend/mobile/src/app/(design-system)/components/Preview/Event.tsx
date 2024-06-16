import React, { forwardRef, useCallback } from 'react';
import { Pressable } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import { Tag } from '@generatesac/lib';
import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';
import { Avatar } from '@rneui/base';

import { SACColors } from '../../shared/colors';
import { Box } from '../Box/Box';
import { Button } from '../Button/Button';
import { EventPageTags } from '../Tag/Tags';
import { Text } from '../Text/Text';

interface EventPreviewProps {
    eventId: string;
}

const tags: Tag[] = [
    {
        id: 'abc123',
        created_at: '2024-06-16T09:30:00Z',
        updated_at: '2024-06-16T10:45:00Z',
        name: 'Product A',
        category_id: 'cat001'
    },
    {
        id: 'def456',
        created_at: '2024-06-15T14:00:00Z',
        updated_at: '2024-06-16T08:20:00Z',
        name: 'Service B',
        category_id: 'cat002'
    },
    {
        id: 'ghi789',
        created_at: '2024-06-14T11:45:00Z',
        updated_at: '2024-06-16T11:00:00Z',
        name: 'Item C',
        category_id: 'cat003'
    },
    {
        id: 'jkl012',
        created_at: '2024-06-15T17:20:00Z',
        updated_at: '2024-06-16T09:05:00Z',
        name: 'Project D',
        category_id: 'cat004'
    },
    {
        id: 'mno345',
        created_at: '2024-06-13T13:00:00Z',
        updated_at: '2024-06-16T10:30:00Z',
        name: 'Task E',
        category_id: 'cat005'
    }
];

const mockEvent = {
    name: 'Spring Showcase 2024',
    clubId: '1',
    hostedClub: 'Generate Product Development Studio',
    preview: `Sandbox is Northeastern's student-led software consultancy building software for Northeastern clients. We work closely with clients across Northeastern to help them best leverage computation.`,
    color: 'darkBlue' as SACColors,
    clubLogo: 'https://avatars.githubusercontent.com/u/45272992?s=200&v=4',
    tags: tags as Tag[]
};

type Ref = BottomSheet;

export const EventPreview = forwardRef<Ref, EventPreviewProps>(
    ({ eventId }, ref) => {
        const renderBackdrop = useCallback(
            (props: any) => (
                <BottomSheetBackdrop
                    appearsOnIndex={0}
                    disappearsOnIndex={-1}
                    {...props}
                />
            ),
            []
        );

        return (
            <BottomSheet
                ref={ref}
                index={-1}
                snapPoints={['50%']}
                enablePanDownToClose={true}
                handleIndicatorStyle={{ backgroundColor: 'white' }}
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
            >
                <Box marginHorizontal="xl" flexDirection="column">
                    <Text paddingTop="s" variant="header-1">
                        {mockEvent.name}
                    </Text>
                    <Box paddingVertical="s" flexDirection="row" gap="xxs">
                        <Text variant="caption-1">Hosted by</Text>
                        <TouchableOpacity
                            onPress={() =>
                                router.push(`/club/${mockEvent.clubId}`)
                            }
                        >
                            <Box flexDirection="row" gap="xxs">
                                <Avatar
                                    rounded
                                    size={15}
                                    source={{ uri: mockEvent.clubLogo }}
                                />
                                <Text variant="caption-1">
                                    {mockEvent.hostedClub}
                                </Text>
                            </Box>
                        </TouchableOpacity>
                    </Box>
                    <Box
                        paddingTop="s"
                        paddingBottom="l"
                        flexDirection="column"
                        gap="s"
                    >
                        <Text>About Event</Text>
                        <Text variant="caption-1">{mockEvent.preview}</Text>
                    </Box>
                    <EventPageTags
                        tags={mockEvent.tags}
                        color={mockEvent.color}
                    />
                    <Button
                        marginTop="m"
                        onPress={() => router.push(`/event/${eventId}`)}
                        width="50%"
                        size="md"
                        variant="standardButton"
                        color={mockEvent.color}
                        alignContent="flex-end"
                    >
                        Learn More
                    </Button>
                </Box>
            </BottomSheet>
        );
    }
);
