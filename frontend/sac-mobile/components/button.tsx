import { Pressable, StyleSheet, Text, View } from 'react-native';

type ButtonProps = {
    title: string;
    backgroundColor?: string;
    color?: string;
    buttonfunc?: () => void;
};

const Button = (props: ButtonProps) => {
    const styles = StyleSheet.create({
        button: {
            backgroundColor: props.backgroundColor || 'gray',
            fontSize: 24,
            paddingTop: '5%',
            paddingBottom: '5%',
            fontWeight: 'bold',
            borderRadius: 15,
            width: '45%'
        },
        title: {
            color: props.color || 'black',
            textAlign: 'center'
        }
    });

    return (
        <Pressable onPress={props.buttonfunc} style={styles.button}>
            <Text style={styles.title}>{props.title}</Text>
        </Pressable>
    );
};

export default Button;
