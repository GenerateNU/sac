
import React, { useRef, useState } from 'react';
import { Dimensions, PanResponder, SafeAreaView, StyleSheet, Platform, StatusBar, View, Text } from 'react-native';
import { useAnimatedStyle, useSharedValue, withTiming } from 'react-native-reanimated';
import { DateContextState } from './context';
import { CalendarProvider, ExpandableCalendar } from 'react-native-calendars';
import HamburgerButton from './components/button';
import { Theme } from 'react-native-calendars/src/types';


import Agenda from './components/agenda';
import { getMarkedDates } from './components/mock';

// FIXME FIXME FIXME FIXME: everything needs to be redone/checked spacing-wise. lots of errors

const EXPANDED_HEIGHT = Dimensions.get('screen').height * .40;
const SHRUNK_HEIGHT = Dimensions.get('screen').height * .20;

// @refresh reset
const CalendarView = () => {
    const notchPadding = Platform.OS == "android" && StatusBar.currentHeight !== undefined
        ? StyleSheet.create({ padding: { paddingTop: StatusBar.currentHeight } }) : StyleSheet.create({ padding: { paddingTop: 0 } });

    const calendarTheme: Theme = {
        selectedDayBackgroundColor: "#000000", 
        todayTextColor: "#000000", 
        arrowColor: "#000000",
        textMonthFontWeight: "bold",
    };

    return (
        <SafeAreaView style={notchPadding.padding} className="bg-white flex w-full h-full">
            <CalendarProvider
                date={"2024-04-07"}>
                <View className="w-full px-[5%] h-12 gap-x-3 flex flex-row justify-start items-center">
                    <HamburgerButton className="w-12 h-12" />
                    <Text className='font-bold text-3xl'>Events</Text>
                </View>   
                <ExpandableCalendar theme={calendarTheme} markedDates={getMarkedDates()}/>
                <Agenda />
            </CalendarProvider>
        </SafeAreaView>
    );
};

export default CalendarView;