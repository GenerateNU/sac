import { Text, View } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

const wordmarkVariants = {
    color: {
        default: 'text-white',
        primary: 'text-black',
        secondary: 'text-gray-500'
    }
};

const wordmarkStyles = cva(['text-2xl', 'font-bold'], {
    variants: wordmarkVariants,
    defaultVariants: { color: 'default' }
});

export interface WordmarkProps extends VariantProps<typeof wordmarkStyles> {}

const Wordmark = ({ color }: WordmarkProps) => {
    return (
        <View className="flex-row items-center">
            <Text className={wordmarkStyles({ color })}>SAC</Text>
        </View>
    );
};

Wordmark.displayName = 'Wordmark';

export { Wordmark, wordmarkVariants };
