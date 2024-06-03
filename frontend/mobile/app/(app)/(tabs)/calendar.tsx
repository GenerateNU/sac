import { Box } from '@/app/(design-system)';
import Calendar from '@/app/(design-system)/components/Calendar/Calendar';
import { EventCard } from '@/app/(design-system)/components/EventCard/EventCard';
import { EventInformation} from '@/utils/mock';
import { Tag } from '@generatesac/lib';
import { useCallback } from 'react';
import { SafeAreaView } from 'react-native';
import { CalendarProvider } from 'react-native-calendars-sac';
import { Theme } from 'react-native-calendars-sac/src/types';

const calendarTheme: Theme = {
    selectedDayBackgroundColor: '#000000',
    todayTextColor: '#000000',
    arrowColor: '#000000',
    textMonthFontWeight: 'bold'
};

const tags: Tag[] = [
    {
        id: '1',
        created_at: new Date('2021-09-01T00:00:00.000Z'),
        updated_at: new Date('2021-09-01T00:00:00.000Z'),
        name: 'Software',
        category_id: '1'
    },
    {
        id: '2',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Free Food',
        category_id: '2'
    },
    {
        id: '3',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Panel Discussion',
        category_id: '3'
    },
    {
        id: '4',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Seminar',
        category_id: '4'
    },
    {
        id: '5',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Hackathon',
        category_id: '5'
    }
];

// We pull events in a date range from the backend (1 month)
// Show two weeks before today and two weeks after today

const CalendarPage = () => {
    const renderEvent = useCallback((event: EventInformation) => {
        return (
            <Box marginVertical='xxs'>
                <EventCard 
                    event="Spring Showcase"
                    variant="calendar"
                    tags={tags}
                    club="Generate"
                    startTime={new Date('2021-09-01T00:00:00.000Z')}
                    endTime={new Date('2021-09-01T00:00:00.000Z')}
                    logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                    image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                />
             </Box>

        )
    }, [])

    return (
        <SafeAreaView style={{ flex: 1 }}>
            <Calendar />
        </SafeAreaView>
    );
};

export default CalendarPage;
