import { SafeAreaView, StyleSheet, Text } from 'react-native';

import { GlobalLayout } from '@/app/(design-system)/components/GlobalLayout/GlobalLayout';
import { useCallback, useEffect, useState } from 'react';

const HomePage = () => {
    return (
        <GlobalLayout>
            <SafeAreaView style={styles.container}>
                <Text style={styles.title}>Home</Text>
                {/* <ScrollView showsVerticalScrollIndicator={false}>
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
                        endTime={date}
                        logo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT800M6T7YVq_f6W49g_UNL29US7gC63nTitg&s"
                        image="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfT6p6kCen0-GphSDogJRd2KoYjg0-QQWuAw9e5JBIBEVTc3Hxho_UwRsZ0IrEi6Ap5oo&usqp=CAU"
                    />
                </ScrollView> */}
            </SafeAreaView>
        </GlobalLayout>
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
