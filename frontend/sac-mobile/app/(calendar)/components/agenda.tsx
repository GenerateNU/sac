import React, { useCallback } from 'react';
import { AgendaList, CalendarProvider } from 'react-native-calendars';
import { getDates, agendaItems } from './mock';
import AgendaItem from './agendaItem';
import { StyleSheet } from 'react-native';

const Agenda = () => {
    const dates = getDates(2024, 4)

    const renderItem = useCallback(({ item }: any) => {
        return <AgendaItem item={item} />;
    }, []);

    return (
        <AgendaList
            sections={agendaItems}

            renderItem={renderItem}
            scrollToNextEvent
            // sectionStyle={styles.section}
        />
    );
};

export default Agenda;
