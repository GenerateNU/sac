import { useRef, useState } from 'react';
import { Animated, Dimensions } from 'react-native';
import { ScrollView } from 'react-native-gesture-handler';
import { SafeAreaView } from 'react-native-safe-area-context';

import { Box } from '../(design-system)';
import { Notification } from './components/notification';
import { UserInfo } from './components/user-info';
import { UserInterests } from './components/user-interest';

const UserDetails = () => {
    const { width } = Dimensions.get('window');

    const [, setScreen] = useState(0);
    const scrollX = useRef(new Animated.Value(0)).current;
    const scrollViewRef = useRef<ScrollView | null>(null);

    const screens = [
        {
            key: 'UserInfo',
            component: (
                <ScrollView
                    contentContainerStyle={{ paddingBottom: 100 }}
                    style={{ width: width }}
                >
                    <UserInfo onPress={() => handleNext(1)} />
                </ScrollView>
            )
        },
        {
            key: 'UserInterests',
            component: (
                <Box width={width}>
                    <UserInterests onPress={() => handleNext(2)} />
                </Box>
            )
        },
        {
            key: 'Notification',
            component: (
                <Box padding="xl" width={width}>
                    <Notification />
                </Box>
            )
        }
    ];

    const handleNext = (screenIndex: number) => {
        setScreen(screenIndex);
        scrollViewRef.current?.scrollTo({
            x: width * screenIndex,
            animated: true
        });
    };

    return (
        <SafeAreaView>
            <Box flexDirection="column" justifyContent="center">
                <Box
                    marginTop="xxl"
                    paddingTop="m"
                    paddingBottom="l"
                    flexDirection="row"
                    justifyContent="center"
                    width="100%"
                    paddingHorizontal="xl"
                >
                    <Box flexDirection="row">
                        {screens.map((_, i) => {
                            const fillWidth = scrollX.interpolate({
                                inputRange: [
                                    width * (i - 1),
                                    width * i,
                                    width * (i + 1)
                                ],
                                outputRange: [0, width * 0.32, width * 0.32],
                                extrapolate: 'clamp'
                            }) as Animated.AnimatedInterpolation<number>;
                            return (
                                <Box
                                    key={i}
                                    height={3}
                                    width={'32%'}
                                    backgroundColor="gray"
                                    borderRadius="sm"
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
                </Box>
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
                    {screens.map((screen) => (
                        <Box key={screen.key} width={width}>
                            {screen.component}
                        </Box>
                    ))}
                </Animated.ScrollView>
            </Box>
        </SafeAreaView>
    );
};

export default UserDetails;
