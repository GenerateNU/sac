import React from 'react';
import {
    Platform,
    SafeAreaView,
    StatusBar,
    StyleSheet,
    Text,
    View
} from 'react-native';
import { CalendarProvider, ExpandableCalendar } from 'react-native-calendars';
import { Theme } from 'react-native-calendars/src/types';

import Agenda from '@/app/(calendar)/components/agenda';
import HamburgerButton from '@/app/(calendar)/components/button';
import { getMarkedDates } from '@/app/(calendar)/components/mock';

// FIXME FIXME FIXME FIXME: everything needs to be redone/checked spacing-wise. lots of errors

// @refresh reset
const CalendarView = () => {
    const notchPadding =
        Platform.OS == 'android' && StatusBar.currentHeight !== undefined
            ? StyleSheet.create({
                  padding: { paddingTop: StatusBar.currentHeight }
              })
            : StyleSheet.create({ padding: { paddingTop: 0 } });

    const calendarTheme: Theme = {
        selectedDayBackgroundColor: '#000000',
        todayTextColor: '#000000',
        arrowColor: '#000000',
        textMonthFontWeight: 'bold'
    };

    return (
        <SafeAreaView
            style={notchPadding.padding}
            className="bg-white flex w-full h-full"
        >
            <CalendarProvider date={'2024-04-07'}>
                <View className="w-full px-[5%] h-12 gap-x-3 flex flex-row justify-start items-center">
                    <HamburgerButton className="w-12 h-12" />
                    <Text className="font-bold text-3xl">Events</Text>
                </View>
                <ExpandableCalendar
                    theme={calendarTheme}
                    markedDates={getMarkedDates()}
                />
                <Agenda />
            </CalendarProvider>
        </SafeAreaView>
    );
};

export default CalendarView;
