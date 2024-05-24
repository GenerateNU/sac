import { StyleSheet, Text, View, SafeAreaView } from 'react-native';
import React from "react";
import {Editor} from "../../components/editor";
import { SafeAreaProvider } from 'react-native-safe-area-context';

const HomePage = () => {
    return (
        <SafeAreaView style={styles.container}>
            <Editor />
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
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
