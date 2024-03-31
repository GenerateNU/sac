
import calendarize from 'calendarize';
import React, { useContext, useRef, useState } from 'react';
import { Button, Image, Dimensions, PanResponder, SafeAreaView, StyleSheet, Text, TouchableOpacity, TouchableOpacityProps, View, ViewProps, Platform, StatusBar, LayoutChangeEvent } from 'react-native';
import Animated, { useAnimatedStyle, useSharedValue, withSpring, withTiming } from 'react-native-reanimated';
import { DateContext } from './context';

import CalendarViewHeader, {CalendarViewHeaderProps} from './components/header';
import { CalendarCarousel, ExpandedCalendar } from './components/grid';
/*
fixme/bug list:
text spacing on calendar grid is weird/wrong
text-white tailwind class doesnt work on calendar cell, (???) have to use style
*/

const EXPANDED_HEIGHT = Dimensions.get('screen').height * .60;
const SHRUNK_HEIGHT = Dimensions.get('screen').height * .40;


// @refresh reset
const CalendarView = () => {
    console.log("Reload check...");

    const [currentDate, setCurrentDate] = useState(new Date());

    const height = useSharedValue(SHRUNK_HEIGHT);
    const expanded = useSharedValue(false);
  
    const context = useSharedValue({ y: 0 })
    const panResponder = useRef(
      PanResponder.create({
        onMoveShouldSetPanResponder: () => true,
        onPanResponderGrant: (e, gestureState) => {
          context.value = { y: height.value }
        },
        onPanResponderMove: (evt, gestureState) => {
          height.value = gestureState.dy + context.value.y;
          // FIXME: height.value should get clamped; even tho onPanResponderRelease will
          // always put it at shrunk or expanded height, we shouldn't have user be able
          // to really shrink or expand it cause thats gross. Clamp doesn't work for me tho
          // so lets fix that
        },
        onPanResponderRelease: () => {
          const distanceToExpanded = Math.abs(height.value - EXPANDED_HEIGHT);
          const distanceToShrunk = Math.abs(height.value - SHRUNK_HEIGHT);

          if (distanceToExpanded < distanceToShrunk) {
            height.value = withTiming(EXPANDED_HEIGHT, { duration: 300});
            expanded.value = true;
          } else {
            height.value = withTiming(SHRUNK_HEIGHT, { duration: 300 });
            expanded.value = false;
          }
        },
      }),
    ).current;
      
    const styleSheet = useAnimatedStyle(() => {
      return {
        height: height.value
      }
    });


    // FIXME: we need to figure out offsets and layout better
    // FIXME: calendar is off center for some reason
    // FIXME: align-center is not working anywhere

    /*
    <ExpandedCalendar 
              className={'absolute left-[8%] w-full top-[130px]'}
              onDayPressed={(date) => {
                // FIXME: why is this so slow
                setCurrentDate(date)
              }}/>*/
              
    const notchPadding = Platform.OS == "android" && StatusBar.currentHeight !== undefined 
      ? StatusBar.currentHeight : 0;

    return (
      <DateContext.Provider value={currentDate}>
        <View className={`pt-[${Math.round(notchPadding)}] flex h-full w-full`}>
          <Animated.View 
            className="px-[8%] overflow-hidden"
            style={styleSheet}>
              <CalendarViewHeader />
              <CalendarCarousel onDayPressed={() => console.log("X")}/>
          
          </Animated.View>
          <View 
              className="w-full h-[25px]"
              {...panResponder.panHandlers}> 
              <View className="ml-[8%] w-[84%] bg-gray-400 h-[1px]" />
              <View className="ml-[45%] w-[10%] bg-gray-400 rounded-full h-[6px] top-2" />
          </View>
        </View>
      </DateContext.Provider>
      
    );
  };
export default CalendarView;