import { Pressable, StyleSheet, Text, View } from 'react-native';

type ButtonProps = {
    title: string;
    backgroundColor?: string;
    color?: string;
    onPress?: () => void;
    borderColor?: string;
};

const Button = (props: ButtonProps) => {
    const styles = StyleSheet.create({
        button: {
            backgroundColor: props.backgroundColor || 'gray',
            fontSize: 24,
            fontWeight: 'bold',
            borderRadius: 15,
            width: '47.5%',
            borderColor: 'black'
        },
        title: {
            color: props.color || 'black',
            textAlign: 'center'
        },
        border: {
            borderWidth: 1,
            borderRadius: 15,
            borderColor: props.borderColor || 'gray'
        }
    });

    return (
        <Pressable onPress={props.onPress} style={styles.button}>
            <View
                className="border-1 rounded-xl pt-[10%]"
                style={styles.border}
            >
                <Text className="pb-[10%] px-[5%]" style={styles.title}>
                    {props.title}
                </Text>
            </View>
        </Pressable>
    );
};

export default Button;
