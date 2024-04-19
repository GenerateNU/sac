import { View, Text, ScrollView } from 'react-native'
import React from 'react'
import { useClubs } from '@/hooks/use-club'
import { ClubHomePageCard } from './club-homepage-card'

const RecommendedClubs = () => {
    const { data: clubs, isLoading, error } = useClubs()

    if (isLoading) {
        return <Text>Loading...</Text>
    }

    if (error) {
        return <Text>Error: {error.message}</Text>
    }

    return (
        <View>
            <ScrollView className="px-[6%] pt-[3%]" horizontal showsHorizontalScrollIndicator={false}>
                {clubs?.map((club) => (
                    <ClubHomePageCard key={club.id} club={club} />
                ))}
            </ScrollView>
        </View>
    )

}

export { RecommendedClubs }