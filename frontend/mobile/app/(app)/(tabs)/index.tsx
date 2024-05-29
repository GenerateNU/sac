import { SafeAreaView, ScrollView, StyleSheet, Text } from 'react-native';

import { Tag } from '@generatesac/lib';

import { Box } from '@/app/(design-system)';
import { EventCard } from '@/components/event-card';

const tags: Tag[] = [
    {
        id: '1',
        created_at: new Date(),
        updated_at: new Date(),
        name: 'Software Development',
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

const HomePage = () => {
    return (
        <SafeAreaView style={styles.container}>
            <Text style={styles.title}>Home</Text>
            <ScrollView showsVerticalScrollIndicator={false}>
                <Box margin="xl">
                    <EventCard
                        event="Spring Showcase"
                        variant="small"
                        tags={tags}
                        club="Sandbox"
                        startTime={new Date()}
                        endTime={new Date()}
                    />
                </Box>
                <EventCard
                    event="Spring Showcase"
                    variant="big"
                    tags={tags}
                    club="Sandbox"
                    startTime={new Date()}
                    endTime={new Date()}
                />
                <EventCard
                    event="Spring Showcase"
                    variant="horizontal"
                    tags={tags}
                    club="Sandbox"
                    startTime={new Date()}
                    endTime={new Date()}
                />
                <EventCard
                    event="Spring Showcase"
                    variant="calendar"
                    tags={tags}
                    club="Sandbox"
                    startTime={new Date()}
                    endTime={new Date()}
                />
            </ScrollView>
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center'
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold'
    },
    separator: {
        marginVertical: 30,
        height: 1,
        width: '80%'
    }
});

export default HomePage;
