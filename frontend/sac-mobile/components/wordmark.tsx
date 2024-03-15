import { Text, View } from 'react-native';

import { Button } from '@/components/button';
import { cn } from '@/lib/utils';

interface WordmarkProps {
    textColor?: string;
    backgroundColor?: string;
    func?: () => void;
    title?: string;
    // FIX? Naming this to className is not working
    additionalClasses?: string;
}

const Wordmark = ({textColor, backgroundColor, func, title, additionalClasses}: WordmarkProps) => {
    return (
        <View className={`flex flex-row justify-between mx-auto w-full pt-[3%] pb-[5.5%] ${additionalClasses}`}>
            <View>
                <Text className={`text-2xl font-bold ${textColor}`}>
                    Wordmark
                </Text>
            </View>
            {func && (
                <Button onPress={func} variant="secondary" size="sm">
                    {title}
                </Button>
            )}
        </View>
    );
};

export default Wordmark;