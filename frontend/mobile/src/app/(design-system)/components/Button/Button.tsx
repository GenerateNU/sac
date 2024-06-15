import React from 'react';
import { GestureResponderEvent, TouchableOpacity } from 'react-native';

import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { BoxProps, createBox } from '@shopify/restyle';

import { Text } from '@/src/app/(design-system)/components/Text/Text';
import {
    SACColors,
    defaultColor,
    textColorVariants
} from '@/src/app/(design-system)/shared/colors';
import { Theme, createStyles } from '@/src/app/(design-system)/theme';

import { ComponentSizes } from '../../shared/types';
import { Icon } from '../Icon/Icon';

interface BaseButtonProps {
    onPress?: (event: GestureResponderEvent) => void;
    disabled?: boolean;
    children?: React.ReactNode;
    color?: SACColors;
    outline?: boolean;
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
    outline = false,
    ...rest
}) => {
    const localStyles = computeButtonStyles({ disabled, color, size, outline });

    return (
        <BaseButton
            disabled={disabled || !onPress}
            onPress={disabled ? undefined : onPress}
            {...localStyles}
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
    outline = false,
    ...rest
}) => {
    const buttonIcon = <Icon size={size} icon={icon} color={iconColor} />;
    const localStyles = computeButtonStyles({
        disabled,
        color,
        size,
        outline,
        variant: 'iconButton'
    });

    return (
        <BaseButton
            justifyContent={justify}
            disabled={disabled || !onPress}
            onPress={disabled ? undefined : onPress}
            {...localStyles}
            {...rest}
        >
            {iconPosition === 'left' && buttonIcon}
            {children && (
                <Text color={textColorVariants[color ?? 'black']}>
                    {children}
                </Text>
            )}
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

const computeButtonStyles = (props: Partial<ButtonProps>) => {
    let buttonStyles = {
        ...styles.base,
        ...buttonSizeStyles[props.size ?? 'full'],
        backgroundColor: props.disabled ? 'gray' : props.color
    };
    if (props.outline) {
        buttonStyles = { ...buttonStyles, ...styles.buttonOutline };
    }
    if (props.variant === 'iconButton') {
        buttonStyles = { ...buttonStyles, ...styles.iconButton };
    }

    return buttonStyles;
};

const styles = createStyles({
    base: {
        alignItems: 'center',
        justifyContent: 'center',
        paddingHorizontal: 'l',
        paddingVertical: 'm',
        borderRadius: 'md'
    },
    iconButton: {
        flexDirection: 'row',
        gap: 's'
    },
    buttonOutline: {
        borderColor: 'darkGray',
        borderWidth: 0.3
    }
});

const buttonSizeStyles = createStyles({
    xs: {
        paddingHorizontal: 's',
        paddingVertical: 'xxs'
    },
    sm: {
        minWidth: 80,
        paddingVertical: 'xs',
        paddingHorizontal: 'm'
    },
    md: {
        minWidth: 115
    },
    lg: {
        minWidth: 150
    },
    full: {
        minWidth: '100%'
    }
});
