import { CalendarProvider, ExpandableCalendar } from "react-native-calendars-sac";
import { GlobalLayout } from "../GlobalLayout/GlobalLayout";
import { Theme } from "react-native-calendars-sac/src/types";
import React, { useEffect, useState } from "react";
import { getMarkedDates } from "@/utils/mock";
import { EventSection, fetchEvents } from "./mockData";
import AgendaList from "./Agenda/AgendaList";

const calendarTheme: Theme = {
    selectedDayBackgroundColor: '#000000',
    todayTextColor: '#000000',
    arrowColor: '#000000',
    textMonthFontWeight: 'bold'
};

export default function Calendar() {
    const [isLoading, setIsLoading] = useState(true);
    const [events, setEvents] = useState<EventSection[]>([]);

    const fetchData = async () => {
        setIsLoading(true);
        setTimeout(() => {
            fetchEvents(123)
            .then((events) => {
                setEvents((oldData) => [...oldData, ...events]);
                setIsLoading(false);
            })
        }, 2500)
    };

    useEffect(() => {
        fetchData();
    }, []);

    return (
        <CalendarProvider 
            date={new Date().toISOString()}
            theme={calendarTheme}
            disabledOpacity={0.5}
        >
            <ExpandableCalendar
                style={{
                    shadowColor: 'gray',
                    shadowOffset: { width: 0, height: 10 },
                    shadowOpacity: 0.1,
                    shadowRadius: 5,
                }}
                theme={calendarTheme}
                allowShadow
                markedDates={getMarkedDates()} // TODO: Add marked dates
                closeOnDayPress={false}
            />
            <GlobalLayout>
                <AgendaList 
                    sections={events}
                    fetchData={fetchData}
                    isLoading={isLoading}
                />
            </GlobalLayout>
        </CalendarProvider>
    );
}
