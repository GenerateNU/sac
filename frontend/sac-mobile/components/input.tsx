import { StyleSheet, Text, TextInput, View } from 'react-native';

type InputProps = {
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

const Input = (props: InputProps) => {
    const borderColor = props.error ? 'border-red-600' : 'border-gray-500';
    return (
        <View>
            <Text className="pb-[2%]">{props.title}</Text>
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
