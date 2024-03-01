import { StyleSheet, Text, TextInput, View } from 'react-native';

type InputProps = {
    title: string;
    placeholder: string;
};

export const Input = (props: InputProps) => {
    const styles = StyleSheet.create({
        input: {
            borderRadius: 10,
            borderWidth: 1,
            paddingTop: '5%',
            paddingBottom: '5%',
            width: '100%'
        }
    });
    return (
        <View >
            <Text>{props.title}</Text>
            <View className="items-center">
                <TextInput
                    style={styles.input}
                    placeholder={props.placeholder}
                ></TextInput>
            </View>
        </View>
    );
};

export default Input;
