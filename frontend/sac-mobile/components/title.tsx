import React from 'react';
import { Text } from 'react-native';

const Title = ({ title }: { title: string }) => {
    return (
        <Text className="text-zinc-900 text-[32px] font-normal leading-9">
            {title}
        </Text>
    );
};

export { Title };
