import calendarize from "calendarize";
import { useContext, useEffect, useRef, useState } from "react";
import { Dimensions, Text, TouchableOpacity, View, ViewProps } from "react-native";
import { DateContext } from "../context";
import Animated, { runOnJS, useAnimatedScrollHandler, useSharedValue } from "react-native-reanimated";
import { addMonths, addWeeks, getDate, isSameDay, setDate } from "date-fns";
import { calendarizeHelper } from "@/lib/utils";

const CALENDAR_WIDTH = Math.round(Dimensions.get('screen').width * .84);
const CALENDAR_PADDING = Math.round(Dimensions.get('screen').width * .08);
const CALENDAR_TOTAL_WIDTH = CALENDAR_WIDTH + CALENDAR_PADDING;

export interface CalendarCarouselProps extends ViewProps {
    monthly: boolean,
    onDayPressed: (day: Date) => void,
    onNewMonthDisplayed: (day: Date) => void,
}

// @refresh reset
export const CalendarCarousel = (p: CalendarCarouselProps) => {
    const dateCtx = useContext(DateContext);

    const carouselItems: Date[] = [];

    for (let i = 0; i < 10; i++) {        if (p.monthly) {
            carouselItems.push(addMonths(dateCtx.selectedDate, i));
        } else {
            carouselItems.push(addWeeks(dateCtx.selectedDate, i));
        }
    }

    const scrollRef = useRef<Animated.ScrollView>(null);
    const currentPage = useSharedValue(0);

    const scrollHandler = useAnimatedScrollHandler({ 
        'onEndDrag': (evt) => {
            // todo: infinite scroll logic
            console.log(`Ending event: ${evt.contentOffset.x}`);
        },
        'onScroll': (evt) => {
            // FIXME: When the user swipes to a new month it should update the view header.
            // This is the code that should do it, but because this is running on UI thread,
            // running into issues where we cannot call p.onNewMonthDisplayed.
            //const page = Math.round(Math.round(evt.contentOffset.x) / CALENDAR_WIDTH);
            //if (page !== currentPage.value && page >= 0 && page < carouselItems.length) {
            //    currentPage.value = page;
            //    runOnJS(() => {
            //        p.onNewMonthDisplayed(carouselItems[page])
            //    })()
            //}
        }
    });

    return (
        <Animated.ScrollView
            className="flex flex-row gap-[33px]"
            ref={scrollRef}
            horizontal
            onScroll={scrollHandler}
            showsHorizontalScrollIndicator={false}
            snapToInterval={CALENDAR_TOTAL_WIDTH}
            {...p}
            >
                {carouselItems.map((val, ind, arr) => {
                    return (<CalendarMonth monthly={p.monthly} weekOrMonth={val} className={`w-[346px]`} onDayPressed={(day) => {
                        p.onDayPressed(day);
                    }}/>);
                })}
        </Animated.ScrollView>
    )
}

export interface CalendarDays extends ViewProps {
    monthly: boolean
    weekOrMonth: Date,
    onDayPressed: (day: Date) => void
}

export const CalendarMonth = (p: CalendarDays) => {
    const dateCtx = useContext(DateContext);
    const calendarDays = calendarizeHelper(p.weekOrMonth, p.monthly ? 'month' : 'week');

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
            const marked = isSameDay(day, dateCtx.selectedDate);

            // FIXME: text-white tailwind class isnt working so i have to use this VV
            
            const textColor = marked ? "white" : "black";
            return (
                <TouchableOpacity
                    onPress={() => { p.onDayPressed(day) }}
                    className={`${marked ? "bg-black" : ""} rounded flex flex-col items-center w-[8%] h-4`}>
                    <Text
                        className={`text-sm font-normal text-center`}
                        style={{ color: textColor, verticalAlign: "middle" }}>
                        {getDate(day)}
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
        <View {...p} className={`bg-yellow-200 flex flex-column gap-3`}>
            {calendarHeader}
            {calendarRows}
        </View>
    )
}

