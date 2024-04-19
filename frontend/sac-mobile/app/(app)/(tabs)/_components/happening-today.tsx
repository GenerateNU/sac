import { View, Text, ScrollView } from 'react-native'
import React from 'react'
import { Event } from '@/types/event'
import { useEvents } from '@/hooks/use-event'
import { EventHomePageCard } from './event-homepage-card'

const HappeningToday = () => {
    const { data: events, isLoading, error } = useEvents()

    if (isLoading) {
        return <Text>Loading...</Text>
    }

    if (error) {
        return <Text>Error: {error.message}</Text>
    }

    return (
        <ScrollView horizontal showsHorizontalScrollIndicator={false} className="px-[6%] pt-[3%]">
            {events?.map((event: Event) => (
                <EventHomePageCard event={event} key={event.id} />
            ))}
        </ScrollView>
    )

}

export { HappeningToday }