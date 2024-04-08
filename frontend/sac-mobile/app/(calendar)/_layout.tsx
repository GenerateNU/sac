
import calendarize from 'calendarize';
import React, { useContext, useRef, useState } from 'react';
import { Button, Image, Dimensions, PanResponder, SafeAreaView, StyleSheet, Text, TouchableOpacity, TouchableOpacityProps, View, ViewProps, Platform, StatusBar, LayoutChangeEvent } from 'react-native';
import Animated, { useAnimatedStyle, useSharedValue, withSpring, withTiming } from 'react-native-reanimated';
import { DateContext, DateContextState } from './context';

import CalendarViewHeader, {CalendarViewHeaderProps} from './components/header';
import { CalendarCarousel, CalendarDays } from './components/grid';
import DayScrollView from '@/app/(calendar)/components/dayScrollView';
import Agenda from './components/agenda';

// FIXME FIXME FIXME FIXME: everything needs to be redone/checked spacing-wise. lots of errors

const EXPANDED_HEIGHT = Dimensions.get('screen').height * .40;
const SHRUNK_HEIGHT = Dimensions.get('screen').height * .20;

// @refresh reset
const CalendarView = () => {
    const [dateContextState, setDateContextState] = useState({
        topViewDate: new Date(),
        selectedDate: new Date()
    } as DateContextState);

    const height = useSharedValue(SHRUNK_HEIGHT);
    const [expanded, setExpanded] = useState(true);
  
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
            setExpanded(true);
          } else {
            height.value = withTiming(SHRUNK_HEIGHT, { duration: 300 });
            setExpanded(false);
          }
        },
      }),
    ).current;
      
    const topSheetHeight = useAnimatedStyle(() => {
      return {
        height: height.value
      }
    });
    
    const notchPadding = Platform.OS == "android" && StatusBar.currentHeight !== undefined 
      ? StatusBar.currentHeight : 0;

    // FIXME: Id like to do pt-[notchPadding px] but it is not working.
    return (
      <DateContext.Provider value={dateContextState}>
        <SafeAreaView className={`pt-12 flex h-full w-full`}>
          <Animated.View 
            className="px-[8%] bg-green-500 w-full flex items-center flex-col overflow-hidden"
            style={topSheetHeight}>
              <CalendarViewHeader/>
              <CalendarCarousel monthly={expanded} 
                onDayPressed={(day) => {
                  setDateContextState({
                    topViewDate: day,
                    selectedDate: day
                  });
                }}
                onNewMonthDisplayed={(day) => {
                  setDateContextState({
                    topViewDate: day,
                    selectedDate: dateContextState.selectedDate
                  });
                }}/>
          </Animated.View>
          <View 
              className="w-full h-[25px]"
              {...panResponder.panHandlers}> 
              <View className="ml-[8%] w-[84%] bg-gray-400 h-[1px]" />
              <View className="ml-[45%] w-[10%] bg-gray-400 rounded-full h-[6px] top-2" />
          </View>
          <Agenda/>
        </SafeAreaView>
      </DateContext.Provider>
      
    );
  };
export default CalendarView;