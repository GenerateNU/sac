import { StyleSheet, Text, View } from 'react-native';

type WordmarkProps = {
    textColor?: string;
    backgroundColor?: string;
}

export default const Wordmark = (props: WordmarkProps) => {
    const styles = StyleSheet.create({
        wordmark: {
            fontSize: 24,
            paddingTop: '2.5%',
            paddingBottom: '7.5%',
            fontWeight: 'bold',
            color: props.textColor
        },
        wordmarkView: {
            backgroundColor: props.backgroundColor,
        }
    });

    return (
        <View style={styles.wordmarkView}>
            <Text style={styles.wordmark}>Wordmark</Text>
        </View>
    );
};