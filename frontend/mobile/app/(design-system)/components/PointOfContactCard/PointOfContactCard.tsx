import React from 'react';

import { Avatar } from '@rneui/themed';

import { Colors, SACColors } from '../../shared/colors';
import { Box } from '../Box/Box';
import { Text } from '../Text/Text';

interface PointOfContactProps {
    imageUrl?: string;
    name: string;
    position: string;
    color: SACColors;
}

export const PointOfContact: React.FC<PointOfContactProps> = ({
    name,
    position,
    color
}) => {
    return (
        <Box flexDirection="row" alignItems="center" marginRight="s">
            <Avatar
                size={32}
                icon={{ name: 'user', type: 'font-awesome' }}
                containerStyle={{
                    backgroundColor: Colors[color],
                    borderRadius: 4
                }}
                // source={imageUrl ? { uri: imageUrl } : undefined}
            />
            <Box marginLeft="xs">
                <Text variant="caption-1" style={{ fontWeight: '600' }}>
                    {name}
                </Text>
                <Text variant="caption-1">{position}</Text>
            </Box>
        </Box>
    );
};
