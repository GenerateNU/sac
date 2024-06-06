import React, { useRef, useState } from 'react';
import {
    Animated,
    Dimensions,
    ScrollView,
    TouchableOpacity
} from 'react-native';

import { Stack, router } from 'expo-router';

import {
    faCalendarDays,
    faMagnifyingGlass,
    faUserPlus
} from '@fortawesome/free-solid-svg-icons';

import { Box, Text } from '../(design-system)';
import { Button } from '../(design-system)/components/Button/Button';
import { Screen, ScreenProps } from './components/splash-screen';
import { Wordmark } from './components/wordmark';

const WelcomePage = () => {
    const [screen, setScreen] = useState(0);
    const scrollX = useRef(new Animated.Value(0)).current;
    const scrollViewRef = useRef<ScrollView | null>(null);

    const screens: ScreenProps[] = [
        {
            title: 'Discover clubs & events',
            description:
                'Explore clubs and events tailored to your interests and searches',
            icon: faMagnifyingGlass
        },
        {
            title: 'Follow & join clubs',
            description:
                'Join clubs and be in the know on their upcoming events and latest news',
            icon: faUserPlus
        },
        {
            title: 'Add events to your calendar',
            description:
                'Keep track of all your upcoming events by adding them to your calendar',
            icon: faCalendarDays
        }
    ];

    const { width } = Dimensions.get('window');
    const handleNext = () => {
        if (screen < screens.length - 1) {
            setScreen(screen + 1);
            scrollViewRef.current?.scrollTo({
                x: width * (screen + 1),
                animated: true
            });
        } else {
            router.push('/(auth)/signup');
        }
    };

    return (
        <Box flex={1} backgroundColor="white">
            <Stack.Screen
                options={{
                    headerShown: true,
                    headerLeft: () => <Wordmark />,
                    headerRight: () => (
                        <TouchableOpacity
                            onPress={() => router.push('/(auth)/signup')}
                        >
                            <Text>Skip</Text>
                        </TouchableOpacity>
                    )
                }}
            />
            <Box style={{ paddingTop: 200 }} flex={1}>
                <Box flex={2}>
                    <Animated.ScrollView
                        horizontal
                        pagingEnabled
                        scrollEnabled={false}
                        showsHorizontalScrollIndicator={false}
                        ref={scrollViewRef}
                        scrollEventThrottle={16}
                        onScroll={Animated.event(
                            [
                                {
                                    nativeEvent: {
                                        contentOffset: { x: scrollX }
                                    }
                                }
                            ],
                            { useNativeDriver: false }
                        )}
                    >
                        {screens.map((screenItem, index) => (
                            <Box
                                marginTop="xxl"
                                justifyContent="center"
                                padding="xxl"
                                key={index}
                                style={{ width }}
                            >
                                <Screen
                                    title={screenItem.title}
                                    description={screenItem.description}
                                    icon={screenItem.icon}
                                />
                            </Box>
                        ))}
                    </Animated.ScrollView>
                </Box>
                <Box flex={1} paddingHorizontal="xl">
                    <Box
                        flexDirection="row"
                        paddingBottom="xxl"
                        justifyContent="center"
                        alignItems="center"
                    >
                        {screens.map((_, i) => {
                            const fillWidth = scrollX.interpolate({
                                inputRange: [
                                    width * (i - 1),
                                    width * i,
                                    width * (i + 1)
                                ],
                                outputRange: [0, 20, 20],
                                extrapolate: 'clamp'
                            }) as Animated.AnimatedInterpolation<number>;
                            return (
                                <Box
                                    key={i}
                                    height={3}
                                    width={20}
                                    backgroundColor="gray"
                                    borderRadius="small"
                                    marginHorizontal="xxs"
                                    overflow="hidden"
                                >
                                    <Animated.View
                                        style={{
                                            backgroundColor: 'black',
                                            height: '100%',
                                            width: fillWidth
                                        }}
                                    />
                                </Box>
                            );
                        })}
                    </Box>
                    <Button onPress={handleNext} variant="standardButton">
                        Get Started
                    </Button>
                </Box>
            </Box>
        </Box>
    );
};

export default WelcomePage;
