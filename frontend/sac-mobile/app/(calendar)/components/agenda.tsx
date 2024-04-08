import React, { useCallback } from 'react';
import { AgendaList, CalendarProvider } from 'react-native-calendars';

import AgendaItem from './agendaItem';

const Agenda = () => {
    const dates = [
        '2024-04-01',
        '2024-04-02',
        '2024-04-03',
        '2024-04-04',
        '2024-04-05',
        '2024-04-06',
        '2024-04-07',
        '2024-04-08',
        '2024-04-09',
        '2024-04-10',
        '2024-04-11',
        '2024-04-12',
        '2024-04-13',
        '2024-04-14'
    ];
    const agendaItems = [
        {
            title: dates[0],
            data: [{ hour: '12am', duration: '1h', title: 'Club GM' }]
        },
        {
            title: dates[1],
            data: [
                { hour: '4pm', duration: '1h', title: 'Intramural game' },
                { hour: '5pm', duration: '1h', title: 'Generate Team Meeting' },
                { hour: '5pm', duration: '1h', title: 'Other commitment' }
            ]
        },
        {
            title: dates[2],
            data: [
                { hour: '1pm', duration: '1h', title: 'Dance Crew Rehearsal' },
                { hour: '2pm', duration: '1h', title: 'Intramural game' },
                { hour: '3pm', duration: '1h', title: 'Club GM' }
            ]
        },
        {
            title: dates[3],
            data: [{ hour: '12am', duration: '1h', title: 'Ashtanga Yoga' }]
        },
        {
            title: dates[4],
            data: []
        },
        {
            title: dates[5],
            data: [
                {
                    hour: '11pm',
                    duration: '1h',
                    title: 'Generate Team Meeting'
                },
                { hour: '12pm', duration: '1h', title: 'Running Group' }
            ]
        },
        {
            title: dates[6],
            data: [{ hour: '12am', duration: '1h', title: 'Intramural game' }]
        },
        {
            title: dates[7],
            data: []
        },
        {
            title: dates[8],
            data: [
                { hour: '9pm', duration: '1h', title: 'Cheese Club GM' },
                { hour: '10pm', duration: '1h', title: 'Club GM' },
                { hour: '11pm', duration: '1h', title: 'Intramural game' },
                { hour: '12pm', duration: '1h', title: 'Running Group' }
            ]
        },
        {
            title: dates[9],
            data: [
                { hour: '1pm', duration: '1h', title: 'Intramural game' },
                { hour: '2pm', duration: '1h', title: 'Club GM' },
                { hour: '3pm', duration: '1h', title: 'Private Yoga' }
            ]
        },
        {
            title: dates[10],
            data: [{ hour: '12am', duration: '1h', title: 'Gym Workout' }]
        },
        {
            title: dates[11],
            data: [
                { hour: '1pm', duration: '1h', title: 'Intramural game' },
                { hour: '2pm', duration: '1h', title: 'Club GM' },
                { hour: '3pm', duration: '1h', title: 'Private Yoga' }
            ]
        },
        {
            title: dates[12],
            data: [{ hour: '12am', duration: '1h', title: 'Gym Workout' }]
        },
        {
            title: dates[13],
            data: [{ hour: '12am', duration: '1h', title: 'Gym Workout' }]
        }
    ];

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
