import { BoxProps, VariantProps, createBox, createRestyleComponent, createVariant } from "@shopify/restyle";
import { Theme } from "../theme";
import { GestureResponderEvent, TouchableOpacity } from "react-native";
import { Text } from "../Text/Text";
import { defaultColor, textColorVariants } from "../Colors";

const buttonVariant = createVariant<Theme, 'buttonVariants'>({
    themeKey: 'buttonVariants'
});

type ButtonProps = VariantProps<Theme, 'buttonVariants'> &
    BoxProps<Theme> & {
        children?: React.ReactNode;
        onPress?: (event: GestureResponderEvent) => void;
        disabled?: boolean;
        icon?: string;
    };

const ButtonBase = createRestyleComponent<ButtonProps, Theme>([buttonVariant], TouchableOpacity);

export const Button: React.FC<ButtonProps> = ({ variant, children, onPress, disabled, icon, ...rest }) => {
    return (
        <ButtonBase 
            variant={disabled ? 'disabled' : variant}
            onPress={disabled ? undefined : onPress}
            disabled={disabled}
            {...rest}
        >
            
            <Text color={textColorVariants[variant ?? defaultColor]}>
                {children}
            </Text>
        </ButtonBase>
    )
}