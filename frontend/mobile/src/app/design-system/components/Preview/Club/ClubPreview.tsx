import React, { forwardRef, useCallback, useState } from 'react';

import { router } from 'expo-router';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';
import { SerializedError } from '@reduxjs/toolkit';
import { FetchBaseQueryError } from '@reduxjs/toolkit/query';
import { Avatar } from '@rneui/base';

import { setClubShouldPreview } from '@/src/store/slices/clubSlice';
import { useAppDispatch, useAppSelector } from '@/src/store/store';

import { Box } from '../../Box/Box';
import { Button } from '../../Button/Button';
import { Text } from '../../Text/Text';
import PreviewError from '../PreviewError';
import ClubPreviewSkeleton from './ClubPreviewSkeleton';

interface ClubPreviewProps {
    clubId: string;
    isLoading: boolean;
    error: FetchBaseQueryError | SerializedError | undefined;
}

type Ref = BottomSheet;
const LOGO_HEIGHT = 77;

export const ClubPreview = forwardRef<Ref, ClubPreviewProps>(
    ({ clubId, isLoading, error }, ref) => {
        const club = useAppSelector((state) => state.club);
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

        const ClubLogo = () => {
            return () => (
                <Box marginLeft="xl">
                    <Avatar
                        containerStyle={{
                            top: -LOGO_HEIGHT / 2,
                            position: 'absolute',
                            borderWidth: 2,
                            borderColor: 'white',
                            backgroundColor: 'white'
                        }}
                        rounded
                        size={LOGO_HEIGHT}
                        source={{
                            uri: club.logo
                        }}
                    />
                </Box>
            );
        };

        const NoLogo = () => {
            return () => <></>;
        };

        const [logo, setLogo] = useState(NoLogo);

        return (
            <BottomSheet
                ref={ref}
                index={-1}
                snapPoints={['45%']}
                enablePanDownToClose={true}
                handleComponent={logo}
                onAnimate={(_, toIndex: number) => {
                    toIndex === 0 && !error
                        ? setLogo(ClubLogo)
                        : setLogo(NoLogo);
                }}
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
                onClose={() => dispatch(setClubShouldPreview(false))}
            >
                <Box overflow="visible" marginLeft="xl" zIndex={1} />
                <Box
                    gap="l"
                    paddingTop="xxl"
                    marginHorizontal="xl"
                    flexDirection="column"
                >
                    {isLoading || club.id === '' ? (
                        <ClubPreviewSkeleton />
                    ) : error ? (
                        <PreviewError />
                    ) : (
                        <>
                            <Text paddingTop="s" variant="header-1">
                                {club.name}
                            </Text>
                            <Box flexDirection="column" gap="xs">
                                <Text fontWeight="bold">Who are we?</Text>
                                <Text variant="body-2">{`${club.preview}...`}</Text>
                            </Box>
                            <Button
                                onPress={() => {
                                    dispatch(setClubShouldPreview(false));
                                    router.push(`/app/club/${clubId}`);
                                }}
                                width="100%"
                                size="md"
                                variant="standardButton"
                                color={'darkRed'}
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
