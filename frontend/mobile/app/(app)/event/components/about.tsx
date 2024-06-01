import React from 'react';
import { Linking } from 'react-native';

import { faVideo } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';
import { Tag } from '@generatesac/lib';

import {
    Box,
    Colors,
    ExpandableText,
    SACColors,
    Text,
    textColorVariants
} from '@/app/(design-system)';
import { Tag as TagComponent } from '@/app/(design-system)';
import { Button } from '@/app/(design-system)/components/Button/Button';

interface AboutEventProps {
    tags: Tag[];
    description: string;
    zoomLink?: string;
    color: SACColors;
}

export const AboutEvent = ({
    tags,
    color,
    description,
    zoomLink
}: AboutEventProps) => {
    const handlePress = () => {
        if (zoomLink) {
            Linking.openURL(zoomLink);
        }
    };

    const renderTag = (item: Tag) => (
        <TagComponent key={item.id} state="selected" color={color}>
            <Text variant="body-2">{item.name}</Text>
        </TagComponent>
    );

    return (
        <Box flexDirection="column" alignItems="flex-start">
            <Text marginBottom="m" variant="subheader-1">
                About Event
            </Text>
            <ExpandableText>{description}</ExpandableText>
            <Box paddingTop="xl" flexDirection="row" flexWrap="wrap" gap="xs">
                {tags.length >= 5
                    ? tags.slice(0, 5).map((item) => renderTag(item))
                    : tags.map((item) => renderTag(item))}
            </Box>
            {zoomLink && (
                <Button
                    marginTop="xl"
                    color={color}
                    variant="standardButton"
                    onPress={handlePress}
                    alignItems="flex-start"
                    height={56}
                >
                    <Box
                        gap="s"
                        flexDirection="row"
                        alignItems="center"
                        justifyContent="center"
                    >
                        <Box
                            marginTop="xxs"
                            flexDirection="row"
                            alignItems="center"
                            justifyContent="center"
                            aspectRatio={1}
                            width={25}
                            backgroundColor="white"
                            borderRadius="full"
                        >
                            <FontAwesomeIcon
                                color={Colors[color]}
                                icon={faVideo}
                            />
                        </Box>
                        <Text
                            paddingTop="xxs"
                            color={
                                textColorVariants[
                                    color as keyof typeof textColorVariants
                                ]
                            }
                        >
                            Click here for Zoom Link
                        </Text>
                    </Box>
                </Button>
            )}
        </Box>
    );
};
