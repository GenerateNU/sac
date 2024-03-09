import { Text, TouchableOpacity, TouchableOpacityProps } from 'react-native';
import { VariantProps, cva } from 'class-variance-authority';
import { cn } from '@/lib/utils';

const buttonVariants = {
    variant: {
        default: ['bg-blue-500', 'text-white'],
        destructive: ['bg-red-500', 'text-white font-bold'],
        secondary: ['bg-gray-500', 'text-white italic'],
        outline: ['border border-gray-500 text-gray-500 font-bold'],
    },
    size: {
        default: 'h-10 px-4 py-2',
        sm: 'h-9 rounded-md px-3',
        lg: 'h-11 rounded-md px-8',
        icon: 'h-10 w-10',
    },
};

const buttonStyles = cva(
    ['rounded-md', 'flex', 'items-center', 'justify-center'],
    {
        variants: buttonVariants,
        defaultVariants: {
            variant: 'default',
            size: 'default',
        },
    }
);

export const Button = ({
    children,
    variant,
    size,
    ...props
}: TouchableOpacityProps & VariantProps<typeof buttonStyles>) => {
    return (
        <TouchableOpacity
            className={cn(buttonStyles({ variant, size }), props.className)}
            {...props}
        >
            <Text className={cn(...buttonVariants.variant[variant ?? 'default'], 'text-center')}>
                {children}
            </Text>
        </TouchableOpacity>
    );
};
export default Button;