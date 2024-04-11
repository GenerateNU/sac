import { View, Text } from 'react-native'
import React from 'react'
import { Stack } from 'expo-router'

const Layout = () => {
    return (
        <Stack>
            <Stack.Screen name="[id]" options={{ headerShown: false }} />
            <Stack.Screen name="faq/[id]" options={{ headerShown: false, presentation: 'modal' }} />

        </Stack>
    )
}

export default Layout