import { StyleSheet, Text, View } from 'react-native';

const DiscoverPage = () => {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>Discover</Text>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center'
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold'
    },
    separator: {
        marginVertical: 30,
        height: 1,
        width: '80%'
    }
});

export default DiscoverPage;
