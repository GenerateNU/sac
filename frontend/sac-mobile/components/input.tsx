import { GestureResponderEvent, Text, TextInput, TextInputProps, TouchableOpacity, View } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

import { Button } from '@/components/button';
import { cn } from '@/lib/utils';

const inputVariants = {
    variant: {
        default: ['pt-[4.5%]', 'pb-[4.5%]', 'pl-[5%]', 'border', 'rounded-xl'],
        faq: [
            'bg-gray-100',
            'rounded-lg',
            'py-[2%]',
            'pl-[4.5%]',
            'pr-[2%]',
            'flex-row',
            'justify-between'
        ]
    }
};

const inputStyles = cva(['w-full'], {
    variants: inputVariants,
    defaultVariants: {
        variant: 'default'
    }
});

export interface InputProps
    extends TextInputProps,
        VariantProps<typeof inputStyles> {
    title?: string;
    error?: boolean;
}

const Input = ({ title, error, variant, ...props }: InputProps) => {
    const borderColor = error ? 'border-red-600' : 'border-gray-500';

    let inputComponent = null;
    if (variant === 'faq') {
        inputComponent = (
            <View
                className={cn(inputStyles({ variant }), props.className)}
                {...props}
            >
                <TextInput
                    className="bg-transparent"
                    placeholder={props.placeholder}
                    autoCapitalize={props.autoCapitalize || 'none'}
                    autoCorrect={props.autoCorrect}
                    onChangeText={props.onChangeText}
                    value={props.value}
                    secureTextEntry={props.secureTextEntry || false}
                />
                {/* @ts-ignore */}
                <Button size="faq" variant="faq" onPress={(e: GestureResponderEvent) => props.onSubmitEditing!(e)} />
            </View>
        );
    } else {
        inputComponent = (
            <TextInput
                className={cn(
                    ...inputVariants.variant[variant ?? 'default'],
                    borderColor
                )}
                autoCapitalize={props.autoCapitalize || 'none'}
                autoCorrect={props.autoCorrect}
                placeholder={props.placeholder}
                onChangeText={props.onChangeText}
                value={props.value}
                secureTextEntry={props.secureTextEntry || false}
                onSubmitEditing={props.onSubmitEditing}
            />
        );
    }

    return (
        <View>
            {title && <Text className="pb-[2%]">{title}</Text>}
            {inputComponent}
        </View>
    );
};

export default Input;
