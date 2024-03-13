import { Pressable, StyleSheet, Text, View } from 'react-native';

import { Button } from '@/components/button';

type WordmarkProps = {
    textColor?: string;
    backgroundColor?: string;
    button?: boolean;
    func?: () => void;
    title?: string;
};

const Wordmark = (props: WordmarkProps) => {
    const styles = StyleSheet.create({
        wordmark: {
            fontSize: 24,
            fontWeight: 'bold',
            color: props.textColor
        },
        wordmarkView: {
            backgroundColor: props.backgroundColor,
            flexDirection: 'row'
        }
    });

    return (
        <View className="flex flex-row justify-between mx-auto w-full items-center pt-[2.5%] pb-[5.5%]">
            <View>
                <Text style={styles.wordmark}>Wordmark</Text>
            </View>
            {props.button && (
                <Button
                    onPress={props.func}
                    children= {props.title}
                    variant="secondary"
                    size="sm"
                >
                </Button>

            )}
        </View>
    );
};

export default Wordmark;
