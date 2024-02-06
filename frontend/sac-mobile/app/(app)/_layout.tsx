import React from 'react'
import { Tabs } from 'expo-router'
import { MaterialCommunityIcons } from '@expo/vector-icons'

const AppLayout = () => {

    return (
        <Tabs>
            <Tabs.Screen
                name="index"
                options={{
                    title: 'Home',
                    headerShown: false,
                    tabBarIcon: ({ color }) => <MaterialCommunityIcons name="home" size={24} color={color} />
                }}
            />
        </Tabs>
    )
}

export default AppLayout