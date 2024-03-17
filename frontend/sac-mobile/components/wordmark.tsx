import { Text, View } from 'react-native';
import { Button } from '@/components/button';

interface WordmarkProps {
    textColor?: string;
    func?: () => void;
    buttonText?: string;
    additionalClasses?: string;
}

const Wordmark = ({textColor, func, buttonText, additionalClasses}: WordmarkProps) => {
    return (
        <View className={`flex flex-row justify-between mx-auto w-full pt-[3%] pb-[5.5%] ${additionalClasses}`}>
            <View>
                <Text className={`text-2xl font-bold ${textColor}`}>
                    Wordmark
                </Text>
            </View>
            {func && (
                <Button onPress={func} variant="secondary" size="sm">
                    {buttonText}
                </Button>
            )}
        </View>
    );
};

export default Wordmark;