
import calendarize from 'calendarize';
import React, { useRef } from 'react';
import { Button, Dimensions, PanResponder, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import Animated, { useAnimatedStyle, useSharedValue, withSpring, withTiming } from 'react-native-reanimated';

const EXPANDED_HEIGHT = Dimensions.get('screen').height * .30;
const SHRUNK_HEIGHT = Dimensions.get('screen').height * .20;

type ExpandedCalendarProps = {
  // YYYY-MM
  yearMonth: Date
}

const ExpandedCalendar = (p: ExpandedCalendarProps) => {
  const calendarDays = calendarize(p.yearMonth);

  const calendarHeaderCells: JSX.Element[] = [];
  ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'].forEach((header, index, array) => {
    calendarHeaderCells.push((
        <Text className="w-[8%] h-4 font-semibold text-xs text-center">
            {header}
        </Text>
    ));
  });

  const calendarHeader = (
      <View className="mb-[-4px] flex flex-row justify-between">
          {calendarHeaderCells}
      </View>
  );

  // Calendar view entries
  const calendarRows = calendarDays.map((week, index, array) => {
    const calendarCells = week.map((day, index2, array2) => {
      const display = day !== 0;
      return (
          <TouchableOpacity
            onPress={() => {}}
            className="w-[8%] h-4">
            <Text className='font-normal text-sm text-center'>
              {display ? day : ""}
            </Text>
          </TouchableOpacity> 
      )
    });
  
    return (
      <View className="flex flex-row justify-between">
        {calendarCells}
      </View>
    );
  });

  return (
    <View className="flex flex-column gap-3">
      {calendarHeader}
      {calendarRows}
    </View>
  )
}

// @refresh reset
const CalendarView = () => {
    console.log("Reload check...");
    const height = useSharedValue(EXPANDED_HEIGHT);
    const expanded = useSharedValue(false);

    const context = useSharedValue({ y: 0 })
    const panResponder = useRef(
      PanResponder.create({
        onMoveShouldSetPanResponder: () => true,
        onPanResponderGrant: (e, gestureState) => {
          context.value = { y: height.value }
        },
        onPanResponderMove: (evt, gestureState) => {
          //console.log(gestureState.dy);
          height.value = gestureState.dy + context.value.y;
          console.log(height.value)
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

    return (
      <View className="flex h-full w-full">
        <Animated.View 
          className="px-[8%] overflow-hidden flex flex-column flex-end bg-gray-400"
          style={styleSheet}>
            <ExpandedCalendar yearMonth={new Date(2024,2)}/>
        </Animated.View>
        <View 
            className="bg-red-200 h-[25px]"
            {...panResponder.panHandlers}> 
        </View>
      </View>
    );
  };
  
  const styles = StyleSheet.create({
    container: {
      flex: 1,
      alignItems: 'center',
      justifyContent: 'center',
    },
    titleText: {
      fontSize: 14,
      lineHeight: 24,
      fontWeight: 'bold',
    },
    box: {
      height: 150,
      width: 150,
      backgroundColor: 'red',
      borderRadius: 5,
    },
  });

export default CalendarView;