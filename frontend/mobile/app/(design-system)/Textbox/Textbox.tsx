import { TextInput } from "react-native"
import { Text } from "@/app/(design-system)/Text/Text"
import { VariantProps, createRestyleComponent, createVariant } from "@shopify/restyle"
import { Theme } from "@/app/(design-system)/theme"
import ErrorIcon from "@/app/(design-system)/assets/svg/Error.svg"
import React from "react"
import { Box } from "@/app/(design-system)/Box"

type TextboxProps = {
    title?: string,
    placeholder?: string
    helperText?: string
    error?: string
    autoFocus?: boolean
    readOnly?: boolean
    inputMode?: 'text' | 'numeric' | 'tel' | 'search' | 'email'
    onChangeText?: () => void
    value?: string
    maxLength?: number
    secureTextEntry?: boolean
}

const textBoxVariant = createVariant<Theme, 'textBoxVariants'>({
    themeKey: 'textBoxVariants'
})

type TextBoxVariantProps = VariantProps<Theme, 'textBoxVariants'> & TextboxProps & {
    children?: React.ReactNode
}

const TextInputBase = createRestyleComponent<TextBoxVariantProps, Theme>([textBoxVariant], TextInput);

export const TextBox : React.FC<TextBoxVariantProps> = ({ 
    title,
    placeholder,
    helperText = ' ',
    error,
    autoFocus = false,
    readOnly = false,
    inputMode = 'text',
    onChangeText,
    value,
    maxLength,
    secureTextEntry = false,
}) => {
    return (
        <>
            <Text color="black" variant="body-1">
                {title}
            </Text>
            <TextInputBase
                variant={error ? 'error' : undefined}
                placeholder={placeholder}
                autoFocus={autoFocus}
                readOnly={readOnly}
                inputMode={inputMode}
                onChangeText={onChangeText}
                value={value}
                maxLength={maxLength}
                secureTextEntry={secureTextEntry}
            >
            </TextInputBase>
            <Box gap='xxs' flexDirection="row" alignItems="center">
                {error && <ErrorIcon />}
                <Text color={error ? "darkRed" : "black"} variant="caption-1">
                    {error ?? helperText}
                </Text>
            </Box>
        </>
    )
}