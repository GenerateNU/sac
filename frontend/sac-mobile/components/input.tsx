import { Text, TextInput, TextInputProps, View } from 'react-native';

interface InputProps extends TextInputProps {
    title: string;
    error?: boolean;
}

const Input = ({ title, error, ...props }: InputProps) => {
    const borderColor = error ? 'border-red-600' : 'border-gray-500';
    return (
        <View>
            <Text className="pb-[2%]">{title}</Text>
            <TextInput
                className={`pt-[4.5%] pb-[4.5%] pl-[5%] w-full border rounded-xl ${borderColor}`}
                autoCapitalize={props.autoCapitalize || 'none'}
                autoCorrect={props.autoCorrect}
                placeholder={props.placeholder}
                onChangeText={props.onChangeText}
                value={props.value}
                secureTextEntry={props.secureTextEntry || false}
                onSubmitEditing={props.onSubmitEditing}
            />
        </View>
    );
};

export default Input;
