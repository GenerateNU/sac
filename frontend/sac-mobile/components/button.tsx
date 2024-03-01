import { Pressable, StyleSheet, Text, View } from 'react-native';

type ButtonProps = {
    title: string;
    backgroundColor?: string;
    color?: string;
    onPress?: () => void;
    borderColor?: string;
    fullWidth?: boolean;
};

const Button = (props: ButtonProps) => {
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
        }
    });

    return (
        <Pressable onPress={props.onPress} style={styles.button}>
            <View
                className="border-1 rounded-xl pt-4 pb-4"
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
