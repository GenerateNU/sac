import React, { forwardRef, useCallback, useEffect } from 'react';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import { Avatar } from '@rneui/base';

import { Box } from '../Box/Box';
import { Button } from '../Button/Button';
import { Text } from '../Text/Text';
import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';
import { useAppDispatch, useAppSelector } from '@/src/store/store';
import { setEventShouldPreview } from '@/src/store/slices/eventSlice';
import useEvent from '@/src/hooks/useEvent';

interface EventPreviewProps {
    eventId: string;
}

type Ref = BottomSheet;

export const EventPreview = forwardRef<Ref, EventPreviewProps>(
    ({ eventId }, ref) => {
        const { apiLoading, apiError } = useEvent(eventId);
        const event = useAppSelector((state) => state.event)
        const { name: clubName, logo: clubLogo } = useAppSelector((state) => state.club);

        const dispatch = useAppDispatch();

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
                snapPoints={['52.5%']}
                enablePanDownToClose={true}
                handleIndicatorStyle={{ backgroundColor: 'white' }}
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
                onClose={() => dispatch(setEventShouldPreview(false))}
            >
                {apiLoading ? (
                    <Text> Hello </Text>
                ) : apiError ? (
                    <Text> Error </Text>
                ) : (
                    <Box flexDirection="column" marginHorizontal='xxl' overflow='hidden'>
                        <Text paddingTop="s" variant="header-1">
                            {`${event.name.slice(0, 32)}${event.name.length > 32 ? '...' : ''}`}
                        </Text>
                        <Box paddingVertical="s" flexDirection="row" gap="xxs">
                            <TouchableOpacity
                                onPress={() =>
                                    router.push(`/club/${event.host}`)
                                }
                            >
                                <Box flexDirection="row" gap="xxs" alignItems='center'>
                                    <Avatar
                                        rounded
                                        size={20}
                                        source={{ uri: clubLogo }}
                                    />
                                    <Text variant="body-2" fontWeight='300'>
                                        {`${clubName.slice(0, 36).trim()}${clubName.length > 37 ? '...' : ''}`}
                                    </Text>
                                </Box>
                            </TouchableOpacity>
                        </Box>
                        <Box
                            paddingTop="s"
                            flexDirection="column"
                            gap="s"
                        >
                            <Text fontWeight='bold'>About Event</Text>
                            <Text variant="body-2">{`${event.description.slice(0, 220).trim()}${event.description.length > 220 ? '...' : ''}`}</Text>
                        </Box>
                        <Button
                            marginTop="l"
                            onPress={() => router.push(`/event/${eventId}`)}
                            width="100%"
                            size="md"
                            variant="standardButton"
                            color={'darkRed'}
                            alignContent="flex-end"
                        >
                            Learn More
                        </Button>
                    </Box>
                )}
            </BottomSheet>
        );
    }
);
