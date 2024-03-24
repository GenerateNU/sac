import { Text, View } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

import { cn } from '@/lib/utils';

const wordmarkVariants = {
    variant: {
        default: 'text-black',
        secondary: 'text-white'
    }
};

const wordmarkStyle = cva(['text-2xl', 'font-bold'], {
    variants: wordmarkVariants,
    defaultVariants: {
        variant: 'default'
    }
});

export interface WordmarkProps extends VariantProps<typeof wordmarkStyle> {}

const Wordmark = ({ variant }: WordmarkProps) => {
    return (
        <View className={`flex flex-row pt-[3%] pb-[5.5%]`}>
            <Text className={cn(wordmarkStyle({ variant }))}>Wordmark</Text>
        </View>
    );
};

export default Wordmark;
