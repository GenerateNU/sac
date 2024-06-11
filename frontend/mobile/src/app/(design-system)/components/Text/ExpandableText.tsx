import React, { useState } from 'react';
import { TouchableOpacity, TouchableOpacityProps } from 'react-native';

import { Box } from '../Box/Box';
import { Text } from './Text';

export interface ExpandableTextProps extends TouchableOpacityProps {
    children: React.ReactNode;
    maxLines?: number;
}

export const ExpandableText = ({
    children,
    maxLines = 5
}: ExpandableTextProps) => {
    const [isExpanded, setIsExpanded] = useState(false);
    return (
        <TouchableOpacity onPress={() => setIsExpanded(!isExpanded)}>
            <Box>
                <Text
                    lineHeight={22}
                    numberOfLines={isExpanded ? undefined : maxLines}
                >
                    {children}
                </Text>
                {!isExpanded && (
                    <Text paddingTop="xs" variant="caption-1">
                        Read More...
                    </Text>
                )}
            </Box>
        </TouchableOpacity>
    );
};
