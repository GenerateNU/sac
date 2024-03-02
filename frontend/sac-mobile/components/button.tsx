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
