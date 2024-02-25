import React from 'react';
import { Pressable, SafeAreaView, StyleSheet, Text, View } from 'react-native';

import Wordmark from '@/components/Wordmark';
import Button from '@/components/button';
import { useAuthStore } from '@/hooks/use-auth';
import { router } from 'expo-router';

const Welcome = () => {

  const redirect = () => {
    router.push('/(auth)/login');
  }

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
                <Button title="Get Started" color="white" buttonfunc={redirect}/>
            </View>
        </SafeAreaView>
    );
};

export default Welcome;

const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        marginBottom: '10%',
        marginLeft: 30,
        marginRight: 30
    },
    header: {
        color: 'black',
        height: '20%',
        fontSize: 50,
        marginTop: '5%',
        fontWeight: 'bold'
    },
    imageHolder: {
        backgroundColor: 'black',
        height: '45%',
        width: '100%',
        borderRadius: 20
    },
    description: {
        color: 'black',
        height: '15%',
        fontSize: 23
    },
    button: {
        height: '10%'
    },
    buttonAlign: {
      flexDirection: 'row',
      justifyContent: 'flex-end'
    }
});
