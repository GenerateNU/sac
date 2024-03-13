import React from 'react';
import { Pressable, SafeAreaView, StyleSheet, Text, View } from 'react-native';

import { router } from 'expo-router';

import Wordmark from '@/components/Wordmark';
import { Button } from '@/components/button';

const Welcome = () => {
    const redirect = () => {
        router.push('/(auth)/login');
    };

    return (
        <SafeAreaView style={styles.container}>
            <Wordmark />
            <View style={styles.imageHolder}></View>
            <Text style={styles.header}>Welcome to StudCal</Text>
            <Text className="leading-8" style={styles.description}>
                Discover, follow, and join all the clubs & events Northeastern
                has to offer
            </Text>
            <View style={styles.buttonAlign}>
                <Button
                    children="Get Started"
                    size="lg"
                    variant="default"
                    onPress={() => router.push('/(auth)/login')}
                />
            </View>
        </SafeAreaView>
    );
};

export default Welcome;

const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        marginBottom: '8%',
        marginLeft: "8%",
        marginRight: "8%"
    },
    header: {
        height: '18%',
        fontSize: 50,
        marginTop: '10%',
        fontWeight: 'bold',
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
        fontSize: 25
    },
    button: {
        marginTop: '3%',
        height: '5%'
    },
    buttonAlign: {
        flexDirection: 'row',
        justifyContent: 'flex-end'
    }
});
