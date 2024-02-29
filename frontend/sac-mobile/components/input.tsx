import { StyleSheet, Text, TextInput, View } from 'react-native';

type InputProps = {
    title: string;
    placeholder: string;
};

export const Input = (props: InputProps) => {
    const styles = StyleSheet.create({
        input: {
            borderRadius: 20, 
            borderWidth: 1, 
            paddingTop: '5%',
            paddingBottom: '5%',
        }
    })
    return (
        <View>
            <Text>{props.title}</Text>
            <TextInput style={styles.input} placeholder={props.placeholder}></TextInput>
        </View>
    );
};

