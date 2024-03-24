import React, { useState } from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View } from 'react-native';


export interface ExpandableTextProps extends TouchableOpacityProps {
    maxLines?: number;
}

const ExpandableText = ({ children, maxLines = 5, }: { children: React.ReactNode, maxLines?: number }) => {
    const [isExpanded, setIsExpanded] = useState(false);

    return (
        <TouchableOpacity onPress={() => setIsExpanded(!isExpanded)}>
            <View>
                <Text numberOfLines={isExpanded ? undefined : maxLines}>
                    {children}
                </Text>
                {!isExpanded && (
                    <Text className='pt-2 italic font-semibold'>Show More</Text>
                )}
            </View>
        </TouchableOpacity>
    );
};

export default ExpandableText;