import { DimensionValue, Pressable, StyleSheet, Text, View } from 'react-native';

type ButtonProps = {
    title: string; // text of the button
    backgroundColor?: string; // background color
    color?: string; // text color
    onPress?: () => void;
    borderColor?: string;
    fullWidth?: boolean;
    padding?: DimensionValue; // padding
};

const Button = (props: ButtonProps) => {
    const padding = props.padding + '%'
    const styles = StyleSheet.create({
        button: {
            backgroundColor: props.backgroundColor || 'gray',
            fontSize: 24,
            fontWeight: 'bold',
            borderRadius: 14,
            width: props.fullWidth ? '100%' : '47.5%',
            borderColor: 'black',
        },
        title: {
            color: props.color || 'black',
            textAlign: 'center',
        },
        border: {
            borderWidth: 1,
            borderColor: props.borderColor || 'gray',
            paddingTop: props.padding || '9%',
            paddingBottom: props.padding || '9%',
        }
    });

    return (
        <Pressable onPress={props.onPress} style={styles.button}>
            <View
                className="border-1 rounded-xl"
                style={styles.border}
            >
                <Text className="px-[5%]" style={styles.title}>
                    {props.title}
                </Text>
            </View>
        </Pressable>
    );
};

export default Button;

// import { Text, TouchableOpacity, TouchableOpacityProps } from 'react-native';

// import { VariantProps, cva } from 'class-variance-authority';

// import { cn } from '@/lib/utils';

// const buttonVariants = {
//     variant: {
//         default: ['bg-blue-500', 'text-white'],
//         destructive: ['bg-red-500', 'text-white font-bold'],
//         secondary: ['bg-gray-500', 'text-white italic'],
//         outline: ['border border-gray-500 text-gray-500 font-bold']
//     },
//     size: {
//         default: 'h-10 px-4 py-2',
//         sm: 'h-9 rounded-md px-3',
//         lg: 'h-11 rounded-md px-8',
//         icon: 'h-10 w-10'
//     }
// };

// const buttonStyles = cva(
//     ['rounded-md', 'flex', 'items-center', 'justify-center'],
//     {
//         variants: buttonVariants,
//         defaultVariants: {
//             variant: 'default',
//             size: 'default'
//         }
//     }
// );

// export interface ButtonProps
//     extends TouchableOpacityProps,
//         VariantProps<typeof buttonStyles> {}

// const Button = ({ children, variant, size, ...props }: ButtonProps) => {
//     return (
//         <TouchableOpacity
//             className={cn(buttonStyles({ variant, size }), props.className)}
//             {...props}
//         >
//             <Text
//                 className={cn(
//                     ...buttonVariants.variant[variant ?? 'default'],
//                     'text-center'
//                 )}
//             >
//                 {children}
//             </Text>
//         </TouchableOpacity>
//     );
// };

// Button.displayName = 'Button';

// export { Button, buttonVariants };
