import React from 'react';
import { GestureResponderEvent, TouchableOpacity } from 'react-native';

import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { BoxProps, createBox } from '@shopify/restyle';

import { Text } from '@/app/(design-system)/components/Text/Text';
import {
    SACColors,
    defaultColor,
    textColorVariants
} from '@/app/(design-system)/shared/colors';
import { Theme, createStyles } from '@/app/(design-system)/theme';

import { ComponentSizes } from '../../shared/types';
import { Icon } from '../Icon/Icon';

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
    iconPosition?: 'left' | 'right';
    justify?: 'center' | 'space-between' | 'flex-end' | 'flex-start';
    size?: ComponentSizes;
}

type VariantButtonProps = StandardButtonProps | IconButtonProps;
type ButtonProps = BaseButtonProps & VariantButtonProps & BoxProps<Theme>;

const BaseButton = createBox<Theme, ButtonProps>(TouchableOpacity);

const StandardButton: React.FC<
    BaseButtonProps & StandardButtonProps & BoxProps<Theme>
> = ({
    children,
    onPress,
    disabled = false,
    color = defaultColor,
    size = 'full',
    ...rest
}) => {
    return (
        <BaseButton
            backgroundColor={disabled ? 'disabled' : color}
            disabled={disabled}
            onPress={disabled ? undefined : onPress}
            {...styles.standardButton}
            {...buttonSizeStyles[size]}
            {...rest}
        >
            <Text color={textColorVariants[color]}>{children}</Text>
        </BaseButton>
    );
};

const IconButton: React.FC<
    BaseButtonProps & IconButtonProps & BoxProps<Theme>
> = ({
    children,
    onPress,
    disabled = false,
    color,
    icon,
    iconColor = 'white',
    iconPosition = 'left',
    justify = 'center',
    size = 'full',
    ...rest
}) => {
    const buttonIcon = <Icon size={size} icon={icon} color={iconColor} />;

    return (
        <BaseButton
            backgroundColor={disabled ? 'disabled' : color}
            justifyContent={justify}
            disabled={disabled}
            onPress={disabled ? undefined : onPress}
            {...styles.iconButton}
            {...buttonSizeStyles[size]}
            {...rest}
        >
            {iconPosition === 'left' && buttonIcon}
            <Text color={textColorVariants[color ?? 'black']}>{children}</Text>
            {iconPosition === 'right' && buttonIcon}
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
    standardButton: {
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 'l',
        paddingVertical: 'm',
        borderRadius: 'base'
    },
    iconButton: {
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 'l',
        paddingVertical: 'm',
        borderRadius: 'base',
        flexDirection: 'row',
        gap: 's'
    }
});

const buttonSizeStyles = createStyles({
    small: {
        minWidth: 80,
        paddingVertical: 'xs',
        paddingHorizontal: 'm'
    },
    medium: {
        minWidth: 115
    },
    full: {
        minWidth: '100%'
    }
});
