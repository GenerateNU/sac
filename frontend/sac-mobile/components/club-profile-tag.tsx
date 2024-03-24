import React from 'react';
import { Text, TouchableOpacity, TouchableOpacityProps, View, Image } from 'react-native';

export interface ClubProfileTagProps extends TouchableOpacityProps {
    name?: string;
}

const ClubProfileTag = ({ name, ...props }: ClubProfileTagProps) => {
    return (
        <TouchableOpacity {...props} className='px-1'>
            <View className='bg-card-bg rounded-lg items-center w-40' >
                {name && <Text className="text-sm text-black">{name}</Text>}
            </View>
        </TouchableOpacity>
    );
};

export { ClubProfileTag };