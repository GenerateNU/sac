
import calendarize from 'calendarize';
import React, { useContext, useRef, useState } from 'react';
import { Button, Image, Dimensions, PanResponder, SafeAreaView, StyleSheet, Text, TouchableOpacity, TouchableOpacityProps, View, ViewProps } from 'react-native';
import Animated, { useAnimatedStyle, useSharedValue, withSpring, withTiming } from 'react-native-reanimated';
import { DateContext } from './context';

/*
fixme/bug list:
text spacing on calendar grid is weird/wrong
text-white tailwind class doesnt work on calendar cell, have to use style
*/

const EXPANDED_HEIGHT = Dimensions.get('screen').height * .60;
const SHRUNK_HEIGHT = Dimensions.get('screen').height * .40;

interface CalendarViewHeaderProps extends ViewProps {
  // FIXME: A) should this be dynamic
  // B) are these names correct
  onMenuButtonPressed?: () => {}
  onNotificationButtonPressed?: () => {}
  onBoxButtonPressed?: () => {}
}

const CalendarViewHeader = (p: CalendarViewHeaderProps) => {
  const date = useContext(DateContext);
  // FIXME: how to do images
  return (
  <View {...p}
    className="w-full h-12 flex flex-row justify-between">
      <View className="flex flex-row gap-3 justify-start items-center">
        <TouchableOpacity
          onPress={() => {
            if (p.onMenuButtonPressed !== undefined) {
              p.onMenuButtonPressed()
            }
          }}>
          <View className='px-2 py-3 w-12 h-12 flex flex-col justify-between'>
            <View className='bg-black w-full h-1'></View>
            <View className='bg-black w-2/3 h-1'></View>
            <View className='bg-black w-1/3 h-1'></View>
          </View>
        </TouchableOpacity>
        <Text className='font-bold text-3xl'>{date.toLocaleString('default', { month: 'long' })}</Text>
      </View>
      <View className="flex flex-row gap-4 justify-end items-center">
        <TouchableOpacity
          onPress={() => {
            if (p.onNotificationButtonPressed !== undefined) {
              p.onNotificationButtonPressed()
            }
          }}>
          <Image style={{ resizeMode: 'contain' }} className="w-6 h-6" source={require("../../assets/images/bell.png")}/>
        </TouchableOpacity>
        <TouchableOpacity
          onPress={() => {
            if (p.onMenuButtonPressed !== undefined) {
              p.onMenuButtonPressed()
            }}}>
          <Image style={{ resizeMode: 'contain' }} className="w-6 h-6" source={require("../../assets/images/box.png")}/>
        </TouchableOpacity>
      </View>
  </View>)

}
interface ExpandedCalendarProps extends ViewProps {
  onDayPressed: (day: Date) => void
}

const ExpandedCalendar = (p: ExpandedCalendarProps) => {
  const currentDate = useContext(DateContext);
  const calendarDays = calendarize(currentDate);

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
      const date = new Date(
        currentDate.getFullYear(),
        currentDate.getMonth(),
        day);

      const display = day !== 0;
      const marked = (currentDate != undefined && day == currentDate.getDate());

      // FIXME: text-white tailwind class isnt working so i have to use this VV
      const textColor = marked ? "white" : "black";

      // FIXME: text spacing on calendar grid is weird/wrong
      return (
          <TouchableOpacity
            onPress={() => { p.onDayPressed(date) }}
            className={`${marked ? "bg-black" : ""} rounded flex flex-col items-center w-[8%] h-4`}>
              <Text
              className={`text-sm font-normal text-center`}
              style={{color:textColor, verticalAlign: "middle"}}>
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
    <View {...p} className="flex flex-column gap-3">
      {calendarHeader}
      {calendarRows}
    </View>
  )
}

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
    return (
      <DateContext.Provider value={currentDate}>
        <SafeAreaView className="flex pt-12 h-full w-full">
          <Animated.View 
            className="px-[8%] overflow-hidden"
            style={styleSheet}>
              <CalendarViewHeader />
              <ExpandedCalendar 
              className={'absolute left-[8%] w-full top-[130px]'}
              onDayPressed={(date) => {
                // FIXME: why is this so slow
                setCurrentDate(date)
              }}/>
          
          </Animated.View>
          <View 
              className="w-full h-[25px]"
              {...panResponder.panHandlers}> 
              <View className="ml-[8%] w-[84%] bg-gray-400 h-[1px]" />
              <View className="ml-[45%] w-[10%] bg-gray-400 rounded-full h-[6px] top-2" />
          </View>
        </SafeAreaView>
      </DateContext.Provider>
      
    );
  };
export default CalendarView;