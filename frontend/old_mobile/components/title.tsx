import React from 'react';
import { Text } from 'react-native';

const Title = ({ title }: { title: string }) => {
    return (
        <Text className="text-zinc-900 text-[32px] leading-9 font-bold">
            {title}
        </Text>
    );
};

export { Title };
