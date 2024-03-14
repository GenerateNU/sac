import { StyleSheet, Text, TextInput, View } from 'react-native';

interface InputProps {
    title: string;
    placeholder: string;
    autoCapitalize?: 'sentences' | 'words' | 'characters';
    autoCorrect: boolean;
    onChangeText: (...event: any[]) => void;
    value: string;
    secureTextEntry?: boolean;
    onSubmitEditing: () => void;
    error?: boolean;
};

const Input = ({title, error, ...props}: InputProps) => {
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
