import { Pressable, StyleSheet, Text, View } from 'react-native';

type InputHeader = {
    text: string;
    fontSize: string;
    color: string;
};
const Header = (props: InputHeader) => {
    const styles = StyleSheet.create({
        text: {
            fontSize: Number(props.fontSize),
            color: props.color,
            fontWeight: 'bold'
        }
    });
    return (
        <View>
            <Text style={styles.text}>{props.text}</Text>
        </View>
    );
};

export default Header;
