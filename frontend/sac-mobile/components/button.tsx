import { Pressable, PressableProps, Text } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

import { cn } from '@/lib/utils';

const buttonVariants = {
    variant: {
        default: ['bg-black', 'text-white'],
        none: [],
        destructive: ['bg-red-500', 'text-white font-bold'],
        secondary: ['text-white'],
        outline: ['border border-gray-600 text-gray-500 font-medium'],
        gray: ['bg-card-bg', 'text-black'],
        icon: ['bg-transparent'],
        underline: ['border-b font-bold border-gray-600 text-lg'],
        pill: ['text-black'],
        menu: ['text-black text-lg']
    },
    size: {
        default: 'h-10 px-4 py-2',
        sm: 'h-9 rounded-xl px-[5%]',
        lg: 'h-12 rounded-xl px-8 py-[4%] w-[47%]',
        icon: 'h-10 w-10',
        screenwide: 'h-12 rounded-xl px-8 py-[4%]',
        pill: 'h-7 py-1 px-2'
    }
};

const buttonStyles = cva(
    ['rounded-md', 'flex', 'items-center', 'justify-center'],
    {
        variants: buttonVariants,
        defaultVariants: {
            variant: 'default',
            size: 'default'
        }
    }
);

export interface ButtonProps
    extends PressableProps,
        VariantProps<typeof buttonStyles> {
            color?: string;
        }

const Button = ({ children, variant, size, color, ...props }: ButtonProps) => {
    return (
        <Pressable
            className={cn(buttonStyles({ variant, size }), props.className)}
            {...props}
            style={color ? {
                backgroundColor: color,
            }: {}}
        >
            <Text
                className={cn(
                    ...(buttonVariants.variant[variant ?? 'default'] ??
                        buttonVariants.variant.default),
                    'text-center'
                )}
            >
                {children as React.ReactNode}
            </Text>
        </Pressable>
    );
};

Button.displayName = 'Button';

export { Button, buttonVariants };
