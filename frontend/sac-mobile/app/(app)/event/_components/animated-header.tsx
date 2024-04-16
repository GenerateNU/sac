import React, { useState, useEffect } from 'react';
import { Animated, View, Text, StyleSheet } from 'react-native';

const AnimatedHeader = ({ scrollY, headerHeight }: any) => {
    const [translateY] = useState(new Animated.Value(0));

    useEffect(() => {
        const unsubscribe = scrollY.addListener((event: any) => {
            const offsetY = event.nativeEvent.contentOffset.y;
            Animated.event(
                [{ nativeEvent: { contentOffset: { y: offsetY } } }],
                { useNativeDriver: true } // Optimize for performance
            )(translateY);
        });

        return unsubscribe; // Cleanup function to prevent memory leaks
    }, [scrollY, headerHeight]);

    const interpolatedStyles = {
        transform: [
            {
                translateY: translateY.interpolate({
                    inputRange: [0, headerHeight],
                    outputRange: [0, -headerHeight],
                    extrapolate: 'clamp', // Prevent overshooting
                }),
            },
        ],
    };

    return (
        <Animated.View style={[styles.header, interpolatedStyles]}>
            <Text style={styles.headerText}>Your App Name</Text>
        </Animated.View>
    );
};

const styles = StyleSheet.create({
    header: {
        backgroundColor: '#fff',
        padding: 16,
        position: 'absolute',
        top: 0,
        left: 0,
        right: 0,
        zIndex: 100, // Ensure header stays on top
    },
    headerText: {
        fontSize: 20,
        fontWeight: 'bold',
    },
});

export { AnimatedHeader };
