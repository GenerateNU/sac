import React, { forwardRef, useCallback, useState } from 'react';
import { router } from 'expo-router';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';
import { Avatar } from '@rneui/base';

import { SACColors, textColorVariants } from '../../shared/colors';
import { Box } from '../Box/Box';
import { Button } from '../Button/Button';
import { Text } from '../Text/Text';

interface ClubPreviewProps {
    clubId: string;
}

const mockClub = {
    name: 'Sandbox',
    preview: `Sandbox is Northeastern's student-led software consultancy building software for Northeastern clients. We work closely with clients across Northeastern to help them best leverage computation.`,
    application_title: 'Apply to be a Member',
    application_link: 'https://google.com',
    color: 'darkBlue' as SACColors,
    is_recruiting: true,
    logo: 'https://avatars.githubusercontent.com/u/45272992?s=200&v=4'
};

type Ref = BottomSheet;
const LOGO_HEIGHT = 77;

export const ClubPreview = forwardRef<Ref, ClubPreviewProps>(
    ({ clubId }, ref) => {
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
                            borderColor: 'white'
                        }}
                        rounded
                        size={LOGO_HEIGHT}
                        source={{
                            uri: mockClub.logo
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
                    toIndex === 0 ? setLogo(ClubLogo) : setLogo(NoLogo);
                }}
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
            >
                <Box overflow="visible" marginLeft="xl" zIndex={1} />
                <Box
                    gap="l"
                    paddingTop="xxl"
                    marginHorizontal="xl"
                    flexDirection="column"
                >
                    <Text paddingTop="s" variant="header-1">
                        {mockClub.name}
                    </Text>
                    <Box flexDirection="column" gap="xs">
                        <Text fontWeight='bold'>Who are we?</Text>
                        <Text variant="body-2">{mockClub.preview}</Text>
                    </Box>
                    <Button
                        onPress={() => router.push(`/club/${clubId}`)}
                        width="100%"
                        size="md"
                        variant="standardButton"
                        color={mockClub.color}
                    >
                        Learn More
                    </Button>
                </Box>
            </BottomSheet>
        );
    }
);
