import React, { forwardRef, useCallback } from 'react';
import { TouchableOpacity } from 'react-native-gesture-handler';

import { router } from 'expo-router';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';
import { SerializedError } from '@reduxjs/toolkit';
import { FetchBaseQueryError } from '@reduxjs/toolkit/query';
import { Avatar } from '@rneui/base';

import { setEventShouldPreview } from '@/src/store/slices/eventSlice';
import { useAppDispatch, useAppSelector } from '@/src/store/store';

import { Box } from '../../Box/Box';
import { Button } from '../../Button/Button';
import { Text } from '../../Text/Text';
import PreviewError from '../PreviewError';
import EventPreviewSkeleton from './EventPreviewSkeleton';

interface EventPreviewProps {
    eventId: string;
    isLoading: boolean;
    error: FetchBaseQueryError | SerializedError | undefined;
}

type Ref = BottomSheet;

export const EventPreview = forwardRef<Ref, EventPreviewProps>(
    ({ eventId, isLoading, error }, ref) => {
        const event = useAppSelector((state) => state.event);
        const { name: clubName, logo: clubLogo } = useAppSelector(
            (state) => state.club
        );

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
                <Box
                    flexDirection="column"
                    marginHorizontal="xxl"
                    overflow="hidden"
                >
                    {isLoading || event.id === '' ? (
                        <EventPreviewSkeleton />
                    ) : error ? (
                        <PreviewError />
                    ) : (
                        <>
                            <Text paddingTop="s" variant="header-1">
                                {`${event.name.slice(0, 32)}${event.name.length > 32 ? '...' : ''}`}
                            </Text>
                            <Box
                                paddingVertical="s"
                                flexDirection="row"
                                gap="xxs"
                            >
                                <TouchableOpacity
                                    onPress={() => {
                                        dispatch(setEventShouldPreview(false));
                                        router.push(`/app/club/${event.host}`);
                                    }}
                                >
                                    <Box
                                        marginTop="xs"
                                        flexDirection="row"
                                        gap="xxs"
                                        alignItems="center"
                                    >
                                        <Text color="darkGray">
                                            Hosted by {clubName}
                                        </Text>
                                        <Avatar
                                            rounded
                                            source={{ uri: clubLogo }}
                                            size={20}
                                        />
                                    </Box>
                                </TouchableOpacity>
                            </Box>
                            <Box paddingTop="s" flexDirection="column" gap="s">
                                <Text fontWeight="bold">About Event</Text>
                                <Text variant="body-2">{`${event.description.slice(0, 220).trim()}${event.description.length > 220 ? '...' : ''}`}</Text>
                            </Box>
                            <Button
                                marginTop="l"
                                onPress={() => {
                                    dispatch(setEventShouldPreview(false));
                                    router.push(`/app/event/${eventId}`);
                                }}
                                width="100%"
                                size="md"
                                variant="standardButton"
                                color={'darkRed'}
                                alignContent="flex-end"
                            >
                                Learn More
                            </Button>
                        </>
                    )}
                </Box>
            </BottomSheet>
        );
    }
);
