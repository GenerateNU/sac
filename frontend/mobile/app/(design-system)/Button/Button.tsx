import React from 'react';
import { GestureResponderEvent, TouchableOpacity } from 'react-native';

import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { createBox } from '@shopify/restyle';

import {
    SACColors,
    defaultColor,
    textColorVariants
} from '@/app/(design-system)/Colors';
import { Text } from '@/app/(design-system)/Text/Text';
import { Theme, createStyles } from '@/app/(design-system)/theme';

import { Icon } from '../Icon/Icon';
import { ComponentSizes } from '../shared/types';

interface BaseButtonProps {
    onPress?: (event: GestureResponderEvent) => void;
    disabled?: boolean;
    children?: React.ReactNode;
    color?: SACColors;
}

interface StandardButtonProps {
    variant: 'standardButton';
    size?: ComponentSizes;
}

interface IconButtonProps {
    variant: 'iconButton';
    icon: IconDefinition;
    iconColor?: SACColors;
    position?: 'left' | 'right';
    justify?: 'center' | 'space-between' | 'flex-end' | 'flex-start';
    size?: ComponentSizes;
}

type VariantButtonProps = StandardButtonProps | IconButtonProps;
type ButtonProps = BaseButtonProps & VariantButtonProps;

const BaseButton = createBox<Theme, ButtonProps>(TouchableOpacity);

const StandardButton: React.FC<BaseButtonProps & StandardButtonProps> = ({
    children,
    onPress,
    disabled = false,
    color = defaultColor,
    size = 'full'
}) => {
    return (
        <BaseButton
            variant="standardButton"
            {...styles.baseButton}
            {...buttonSizeStyles[size]}
            backgroundColor={disabled ? 'disabled' : color}
            disabled={disabled}
            onPress={disabled ? undefined : onPress}
        >
            <Text color={textColorVariants[color]}>{children}</Text>
        </BaseButton>
    );
};

const IconButton: React.FC<BaseButtonProps & IconButtonProps> = ({
    children,
    onPress,
    disabled = false,
    color,
    icon,
    iconColor = 'white',
    position = 'left',
    justify = 'center',
    size = 'full'
}) => {
    const buttonIcon = <Icon size={size} icon={icon} color={iconColor} />;

    return (
        <BaseButton
            variant="iconButton"
            {...styles.baseButton}
            {...styles.iconButton}
            {...buttonSizeStyles[size]}
            backgroundColor={disabled ? 'disabled' : color}
            justifyContent={justify}
            disabled={disabled}
            onPress={disabled ? undefined : onPress}
        >
            {position === 'left' && buttonIcon}
            <Text color={textColorVariants[color ?? 'black']}>{children}</Text>
            {position === 'right' && buttonIcon}
        </BaseButton>
    );
};

export const Button: React.FC<ButtonProps> = (props) => {
    if (props.variant === 'standardButton') {
        return <StandardButton {...props} />;
    }
    if (props.variant === 'iconButton') {
        return <IconButton {...props} />;
    }
    return null;
};

const styles = createStyles({
    baseButton: {
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 'l',
        paddingVertical: 'm',
        borderRadius: 8
    },
    iconButton: {
        flexDirection: 'row',
        gap: 's'
    }
});

const buttonSizeStyles = createStyles({
    small: {
        minWidth: 80,
        paddingVertical: 's',
        paddingHorizontal: 'm'
    },
    medium: {
        minWidth: 115
    },
    full: {
        minWidth: '100%'
    }
});
