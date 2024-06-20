import React from 'react';

import { Avatar } from '@rneui/themed';

interface ClubIconProps {
    imageUrl: string;
    size?: number;
}

export const ClubIcon: React.FC<ClubIconProps> = ({ imageUrl, size = 77 }) => {
    return (
        <Avatar
            size={size}
            rounded
            source={{ uri: imageUrl }}
            containerStyle={{
                borderColor: 'white',
                borderStyle: 'solid',
                borderWidth: 2,
                backgroundColor: 'white'
            }}
        />
    );
};
