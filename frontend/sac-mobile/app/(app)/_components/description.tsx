import React, { useState } from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

type DescriptionProps = {
    title: string;
    description: string;
};

const Description = ({ title, description }: DescriptionProps) => {
    const [expanded, setExpanded] = useState(false);

    const onExpand = () => {
        setExpanded(!expanded);
    };

    return (
        <View>
            <Text className="text-2xl font-normal leading-relaxed text-zinc-900">
                {title}
            </Text>
            <Text className="text-base font-light tracking-tight text-zinc-900">
                {expanded ? description : `${description.slice(0, 100)}...`}
            </Text>
            <TouchableOpacity
                className="py-1 text-xs text-zinc-900 font-extralight"
                onPress={() => onExpand()}
            >
                <Text>Read more...</Text>
            </TouchableOpacity>
        </View>
    );
};

export { Description };
