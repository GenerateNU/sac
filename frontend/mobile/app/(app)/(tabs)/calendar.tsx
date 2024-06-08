import { SafeAreaView } from 'react-native';

import Calendar from '@/app/(design-system)/components/Calendar/Calendar';

// We pull events in a date range from the backend (1 month)
// Show two weeks before today and two weeks after today

const CalendarPage = () => {
    return (
        <SafeAreaView style={{ flex: 1 }}>
            <Calendar />
        </SafeAreaView>
    );
};

export default CalendarPage;
