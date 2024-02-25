import React from 'react';
import {View, Text, StyleSheet, Button} from 'react-native'; 
import { useAuthStore } from '@/hooks/use-auth';

const Welcome = () => {
    return (
        <View style={styles.container}>
          <Text style={styles.wordmark}>Wordmark</Text>
          <Text style={styles.header}>Welcome to StudCal</Text>
          <Text style={styles.description}>Discover, follow, and join all the clubs & events Northeastern has to offer</Text>
          <Button title="Login" />
        </View>
      )
};

export default Welcome;

const styles = StyleSheet.create({
    container: {
      flex: 1, 
      flexDirection: 'column',
      marginTop: '5%', 
      marginBottom: '10%', 
      marginLeft: 30, 
      marginRight: 30,
    },
    wordmark: {
      marginTop: 60, 
      fontSize: 20, 
      flex: 1, 
    },
    header: {
      color: 'black',
      flex: 1.7,
      fontSize: 60, 
    },
    description: {
      color: 'black', 
      flex: 1, 
      fontSize: 18,
    },
  });