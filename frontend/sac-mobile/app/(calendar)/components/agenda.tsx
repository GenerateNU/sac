import React, { useCallback } from 'react';
import { AgendaList } from 'react-native-calendars';

import AgendaItem from '@/app/(calendar)/components/agendaItem';

import { agendaItems, getDates } from './mock';

const Agenda = () => {
    const dates = getDates(2024, 4);

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
