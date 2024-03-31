import calendarize from "calendarize";
import { useContext, useRef, useState } from "react";
import { Dimensions, Text, TouchableOpacity, View, ViewProps } from "react-native";
import { DateContext } from "../context";
import Animated, { useAnimatedScrollHandler, useSharedValue } from "react-native-reanimated";

const CALENDAR_WIDTH = Math.round(Dimensions.get('screen').width * .84);
const CALENDAR_PADDING = Math.round(Dimensions.get('screen').width * .08);
const CALENDAR_TOTAL_WIDTH = CALENDAR_WIDTH + CALENDAR_PADDING;

export interface CalendarCarouselProps extends ViewProps {
    onDayPressed: (day: Date) => void,
}
// @refresh reset

export const CalendarCarousel = (p: CalendarCarouselProps) => {
    const carouselItems = [
        new Date(),
        new Date().setMonth(new Date().getMonth() + 1),
        new Date().setMonth(new Date().getMonth() + 2),
        new Date().setMonth(new Date().getMonth() + 3),
    ];

    const SNAP_INTERVAL = 200;
    const scrollXOffset = useSharedValue(0);

    const [isFirstCard, setIsFirstCard] = useState(scrollXOffset.value === 0);
    const [isLastCard, setIsLastCard] = useState(false);
    const [scrollViewWidth, setScrollViewWidth] = useState(0);
    
    const scrollRef = useRef<Animated.ScrollView>(null);
    const scrollHandler = useAnimatedScrollHandler({ 
        'onEndDrag': (evt) => {
            console.log("This is where we need to do infinite scroll logic");
        },
    });
    const handleContentSizeChange = (width: number, height: number) => {
        console.log("Y");
        setScrollViewWidth(width);
    }

    console.log(`ABCDEF ${CALENDAR_PADDING}`)
    return (
        <Animated.ScrollView
            className="flex flex-row gap-[33px]"
            ref={scrollRef}
            horizontal
            onScroll={scrollHandler}
            showsHorizontalScrollIndicator={false}
            snapToInterval={CALENDAR_TOTAL_WIDTH}
            onContentSizeChange={handleContentSizeChange}>
                {carouselItems.map((val, ind, arr) => {
                    return (
                        <ExpandedCalendar className={`w-[346px]`} onDayPressed={() => {}}/>
                    )
                })}
        </Animated.ScrollView>
    )
}

export interface ExpandedCalendarProps extends ViewProps {
    onDayPressed: (day: Date) => void
}

export const ExpandedCalendar = (p: ExpandedCalendarProps) => {
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
                        style={{ color: textColor, verticalAlign: "middle" }}>
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
        <View {...p} className={`flex flex-column gap-3`}>
            {calendarHeader}
            {calendarRows}
        </View>
    )
}

