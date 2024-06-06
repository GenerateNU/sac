import React from 'react';

import { faXmark } from '@fortawesome/free-solid-svg-icons';

import { SACColors, defaultColor } from '../../shared/colors';
import { Button } from '../Button/Button';

type TagProps = {
    children: React.ReactNode;
    onPress?: () => void;
    borderColor?: SACColors;
    color?: SACColors;
    state?: 'selected' | 'unselected' | 'remove';
};

export const Tag: React.FC<TagProps> = ({
    children,
    color = defaultColor,
    state = 'selected',
    borderColor,
    onPress
}) => {
    if (state === 'selected' || state === 'unselected') {
        return (
            <Button
                variant="standardButton"
                color={state === 'selected' ? color : 'white'}
                size={'small'}
                borderRadius={'full'}
                onPress={onPress}
                borderWidth={borderColor ? 1 : 0}
                borderColor={borderColor ? borderColor : undefined}
            >
                {children}
            </Button>
        );
    }
    if (state === 'remove') {
        return (
            <Button
                variant="iconButton"
                color={color}
                icon={faXmark}
                iconColor={'white'}
                iconPosition={'right'}
                size={'small'}
                borderRadius={'full'}
                onPress={onPress}
            >
                {children}
            </Button>
        );
    }
    return null;
};
