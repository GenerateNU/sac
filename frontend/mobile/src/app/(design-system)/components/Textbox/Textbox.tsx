import React from 'react';
import { TextInput } from 'react-native';

import { BoxProps, createBox } from '@shopify/restyle';

import ErrorIcon from '@/src/app/(design-system)/assets/svg/Error.svg';
import { Box } from '@/src/app/(design-system)/components/Box/Box';
import { Text } from '@/src/app/(design-system)/components/Text/Text';
import { TextVariants } from '@/src/app/(design-system)/components/Text/TextVariants';
import { Theme, createStyles } from '@/src/app/(design-system)/theme';

type TextboxProps = {
    title?: string;
    placeholder?: string;
    helperText?: string;
    error?: string;
    autoFocus?: boolean;
    readOnly?: boolean;
    inputMode?: 'text' | 'numeric' | 'tel' | 'search' | 'email';
    onChangeText?: (() => void) | ((value: any) => void);
    value?: string;
    maxLength?: number;
    secureTextEntry?: boolean;
    children?: React.ReactNode;
} & BoxProps<Theme>;

const BaseTextInput = createBox<Theme, TextboxProps>(TextInput);

export const TextBox: React.FC<TextboxProps> = ({
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
    secureTextEntry = false
}) => {
    return (
        <Box>
            <Text color="black" variant="body-1">
                {title}
            </Text>
            <BaseTextInput
                placeholder={placeholder}
                autoFocus={autoFocus}
                readOnly={readOnly}
                inputMode={inputMode}
                onChangeText={onChangeText}
                value={value}
                maxLength={maxLength}
                secureTextEntry={secureTextEntry}
                {...styles.textInput}
                opacity={readOnly ? 0.5 : 1}
                borderColor={error ? 'darkRed' : 'gray'}
            />
            <Box gap="xxs" flexDirection="row" alignItems="center">
                {error && <ErrorIcon />}
                <Text color={error ? 'darkRed' : 'black'} variant="caption-1">
                    {error ?? helperText}
                </Text>
            </Box>
        </Box>
    );
};

const styles = createStyles({
    textInput: {
        borderWidth: 1,
        borderRadius: 'sm',
        backgroundColor: 'white',
        padding: 'm',
        marginBottom: 'xs',
        marginTop: 'xs',
        minWidth: '100%',
        ...TextVariants['body-1']
    }
});
