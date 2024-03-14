import { Pressable, StyleSheet, Text, View } from 'react-native';

import { Button } from '@/components/button';

interface WordmarkProps {
    textColor?: string;
    backgroundColor?: string;
    button?: boolean;
    func?: () => void;
    title?: string;
};

const Wordmark = (props: WordmarkProps) => {
    return (
        <View className="flex flex-row justify-between mx-auto w-full items-center pt-[3%] pb-[5.5%]">
            <View>
                <Text className={`text-2xl font-bold ${props.textColor}`}>Wordmark</Text>
            </View>
            {props.button && (
                <Button
                    onPress={props.func}
                    variant="secondary"
                    size="sm"
                >{props.title}
                </Button>

            )}
        </View>
    );
};

export default Wordmark;
