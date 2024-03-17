import { Text, View } from 'react-native';

interface WordmarkProps {
    textColor?: string;
}

const Wordmark = ({ textColor }: WordmarkProps) => {
    return (
        <View className={`flex flex-row pt-[3%] pb-[5.5%]`}>
            <Text className={`text-2xl font-bold ${textColor}`}>Wordmark</Text>
        </View>
    );
};

export default Wordmark;
