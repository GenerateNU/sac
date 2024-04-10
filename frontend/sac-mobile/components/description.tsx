import React, { useState } from 'react';
import { Text, TouchableOpacity, View } from 'react-native';

type DescriptionProps = {
    title: string;
    description: string;
};

const Description = ({ title, description }: DescriptionProps) => {
    const [expanded, setExpanded] = useState(false);

    const displayedCharacters = description.length < 100 ? description.length : 100;

    return (
        <View>
            <Text className="text-2xl font-normal leading-relaxed text-zinc-900">
                {title}
            </Text>
            <Text className="text-base font-light tracking-tight text-zinc-900">
                {description.length <= 100 ? (
                    description
                ) : (
                    <>
                        {description.slice(0, displayedCharacters)}
                        <Text style={{ color: 'blue' }}>...</Text>
                    </>
                )}
            </Text>
            {description.length > 100 && (
                <TouchableOpacity
                    className="py-1 text-xs text-zinc-900 font-extralight"
                    onPress={() => setExpanded(!expanded)}
                >
                    <Text>Read more...</Text>
                </TouchableOpacity>
            )}
        </View>
    );
};

export { Description };
