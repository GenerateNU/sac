import React from 'react';

import { faXmark } from '@fortawesome/free-solid-svg-icons';

import { SACColors, defaultColor } from '../../shared/colors';
import { Button } from '../Button/Button';

type TagProps = {
    children: React.ReactNode;
    onPress?: () => void;
    color?: SACColors;
    variant?: 'selected' | 'unselected' | 'remove' | 'eventCard';
};

export const Tag: React.FC<TagProps> = ({
    children,
    color = defaultColor,
    variant = 'selected',
    onPress
}) => {
    if (variant === 'selected' || variant === 'unselected') {
        return (
            <Button
                variant="standardButton"
                color={variant === 'selected' ? color : 'white'}
                size={'sm'}
                borderRadius={'full'}
                onPress={onPress}
                outline={variant === 'unselected' ? true : false}
            >
                {children}
            </Button>
        );
    }
    if (variant === 'remove') {
        return (
            <Button
                variant="iconButton"
                color={color}
                icon={faXmark}
                iconColor={'white'}
                iconPosition={'right'}
                size={'sm'}
                borderRadius={'full'}
                onPress={onPress}
            >
                {children}
            </Button>
        );
    }
    if (variant === 'eventCard') {
        return (
            <Button
                variant="standardButton"
                color={'white'}
                size={'xs'}
                borderRadius={'full'}
                outline
            >
                {children}
            </Button>
        );
    }
    return null;
};
