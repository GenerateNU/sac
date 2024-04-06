import { Text } from 'react-native'
import React from 'react'

const Title = ({ title }: { title: string }) => {
    return (
        <Text className="text-3xl font-bold">
            {title}
        </Text>
    )
}

export { Title }