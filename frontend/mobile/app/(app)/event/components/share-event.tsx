import React, { forwardRef, useCallback, useState } from 'react';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';
import Clipboard from '@react-native-community/clipboard';

import { Box, Text } from '@/app/(design-system)';
import { Button } from '@/app/(design-system)/components/Button/Button';

import { Divider } from './divider';

const ShareEventText = `Spread the word and let others join in on the conversation with just a single click. Sharing has never been easier!`;

interface ShareEventProps {
    link: string;
}

type Ref = BottomSheet;

const ShareEventBottomSheet = forwardRef<Ref, ShareEventProps>(
    ({ link }, ref) => {
        const snapPoints = ['40%'];
        const [isCopied, setIsCopied] = useState(false);

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

        const handleCopyPress = () => {
            Clipboard.setString(link);
            setIsCopied(true);
        };

        return (
            <BottomSheet
                ref={ref}
                index={-1}
                snapPoints={snapPoints}
                enablePanDownToClose={true}
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
            >
                <Box marginHorizontal="xl" flexDirection="column" gap="s">
                    <Text
                        paddingTop="s"
                        textAlign="center"
                        variant="subheader-2"
                        fontWeight="500"
                    >
                        Share Event
                    </Text>
                    <Divider />
                    <Text marginVertical="s" variant="body-1">
                        {ShareEventText}
                    </Text>
                    <Button
                        marginBottom="xxl"
                        variant="standardButton"
                        onPress={handleCopyPress}
                    >
                        <Text fontWeight="500">
                            {isCopied ? 'Copied!' : 'Click to copy link'}
                        </Text>
                    </Button>
                </Box>
            </BottomSheet>
        );
    }
);

export { ShareEventBottomSheet };
