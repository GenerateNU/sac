import { Pressable, StyleSheet, Text, View } from 'react-native';

type ButtonProps = {
    title: string;
    backgroundColor?: string;
    color?: string;
    onPress?: () => void;
    borderColor?: string;
};

export const Button = (props: ButtonProps) => {
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
            borderColor: props.borderColor || 'gray',
            paddingTop: '6%',
            paddingBottom: '6%'
        }
    });

    return (
        <Pressable onPress={props.onPress} style={styles.button}>
            <View style={styles.border}>
                <Text style={styles.title}>{props.title}</Text>
            </View>
        </Pressable>
    );
};
