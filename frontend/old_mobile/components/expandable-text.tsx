import React, { useState } from 'react';
import {
    Text,
    TouchableOpacity,
    TouchableOpacityProps,
    View
} from 'react-native';

export interface ExpandableTextProps extends TouchableOpacityProps {
    children: React.ReactNode;
    maxLines?: number;
}

const ExpandableText = ({ children, maxLines = 5 }: ExpandableTextProps) => {
    const [isExpanded, setIsExpanded] = useState(false);

    return (
        <TouchableOpacity onPress={() => setIsExpanded(!isExpanded)}>
            <View>
                <Text numberOfLines={isExpanded ? undefined : maxLines}>
                    {children}
                </Text>
                {!isExpanded && (
                    <Text className="pt-2 font-light">Read More...</Text>
                )}
            </View>
        </TouchableOpacity>
    );
};

export default ExpandableText;
