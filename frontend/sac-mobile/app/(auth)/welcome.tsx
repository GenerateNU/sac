import React from 'react';
import { Pressable, SafeAreaView, StyleSheet, Text, View } from 'react-native';

import { router } from 'expo-router';

import Wordmark from '@/components/Wordmark';
import Button from '@/components/button';

const Welcome = () => {
    const redirect = () => {
        router.push('/(auth)/login');
    };

    return (
        <SafeAreaView style={styles.container}>
            <Wordmark />
            <View style={styles.imageHolder}></View>
            <Text style={styles.header}>Welcome to StudCal</Text>
            <Text style={styles.description}>
                Discover, follow, and join all the clubs & events Northeastern
                has to offer
            </Text>
            <View style={styles.buttonAlign}>
                <Button
                    title="Get Started"
                    color="white"
                    onPress={redirect}
                />
            </View>
        </SafeAreaView>
    );
};

export default Welcome;

const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        marginBottom: '10%',
        marginLeft: "8%",
        marginRight: "8%"
    },
    header: {
        height: '20%',
        fontSize: 50,
        marginTop: '10%',
        fontWeight: 'bold'
    },
    imageHolder: {
        backgroundColor: 'gray',
        height: '45%',
        width: '100%',
        borderRadius: 20,
        marginTop: '5%'
    },
    description: {
        height: '15%',
        fontSize: 23
    },
    button: {
        height: '5%'
    },
    buttonAlign: {
        flexDirection: 'row',
        justifyContent: 'flex-end'
    }
});
