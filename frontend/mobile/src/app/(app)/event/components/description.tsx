import React, { forwardRef, useCallback } from 'react';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';

import { Box, Text } from '@/src/app/(design-system)';

import { Divider } from './divider';

type Ref = BottomSheet;

interface DescriptionProps {
    description: string;
    type: 'club' | 'event';
}

export const Description = forwardRef<Ref, DescriptionProps>(
    ({ description, type }, ref) => {
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
                enablePanDownToClose
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
            >
                <Box
                    justifyContent="center"
                    marginHorizontal="xl"
                    flexDirection="column"
                    gap="s"
                >
                    <Text
                        paddingTop="s"
                        textAlign="center"
                        variant="subheader-2"
                        fontWeight="500"
                    >
                        About {type === 'club' ? 'Us': 'Event'}
                    </Text>
                    <Divider />
                    <Text paddingTop="m">{description}</Text>
                </Box>
            </BottomSheet>
        );
    }
);
