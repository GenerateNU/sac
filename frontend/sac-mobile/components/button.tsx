import { Text, TouchableOpacity, TouchableOpacityProps } from 'react-native';

import { VariantProps, cva } from 'class-variance-authority';

import { cn } from '@/lib/utils';

const buttonVariants = {
    variant: {
        default: ['bg-gray-500', 'text-white'],
        destructive: ['bg-red-500', 'text-white font-bold'],
        secondary: ['bg-white', 'text-gray'],
        outline: ['border border-gray-600 text-gray-500 font-medium'],
        gray: ['bg-card-bg', 'text-black']
    },
    size: {
        default: 'h-10 px-4 py-2',
        sm: 'h-9 rounded-xl px-[5%]',
        lg: 'h-12 rounded-xl px-8 py-[4%] w-[47%]',
        icon: 'h-10 w-10',
        screenwide: 'h-12 rounded-xl px-8 py-[4%]'
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
    extends TouchableOpacityProps,
        VariantProps<typeof buttonStyles> {}

const Button = ({ children, variant, size, ...props }: ButtonProps) => {
    return (
        <TouchableOpacity
            className={cn(buttonStyles({ variant, size }), props.className)}
            {...props}
        >
            <Text
                className={cn(
                    ...buttonVariants.variant[variant ?? 'default'],
                    'text-center'
                )}
            >
                {children}
            </Text>
        </TouchableOpacity>
    );
};

Button.displayName = 'Button';

export { Button, buttonVariants };
