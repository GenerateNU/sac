import { SafeAreaView, StyleSheet, Text } from 'react-native';

import { GlobalLayout } from '@/app/(design-system)/components/GlobalLayout/GlobalLayout';

const HomePage = () => {
    return (
        <GlobalLayout>
            <SafeAreaView style={styles.container}>
                <Text style={styles.title}>Home</Text>
            </SafeAreaView>
        </GlobalLayout>
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

export default HomePage;
