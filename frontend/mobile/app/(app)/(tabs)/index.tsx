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
                <Box marginLeft="xl">
                    <EventCard
                        event="Spring Showcase"
                        variant="small"
                        tags={tags}
                        club="Generate Product Development Studio"
                        startTime={new Date()}
                        endTime={new Date()}
                        logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                    />
                </Box>
                <EventCard
                    event="Spring Showcase"
                    variant="big"
                    tags={tags}
                    club="Generate Product Development Studio"
                    startTime={new Date()}
                    endTime={new Date()}
                    logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                    image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                />
                <EventCard
                    event="Spring Showcase"
                    variant="club"
                    tags={tags}
                    club="Generate"
                    startTime={new Date()}
                    endTime={new Date()}
                    logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                    image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                />
                <EventCard
                    event="Spring Showcase"
                    variant="calendar"
                    tags={tags}
                    club="Generate"
                    startTime={new Date()}
                    endTime={new Date()}
                    logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                    image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
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
