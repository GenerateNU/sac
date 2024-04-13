import React from 'react';
import { SafeAreaView, Text } from 'react-native';

import { Link, Stack, useLocalSearchParams } from 'expo-router';

const ClubPage = () => {
    const { id } = useLocalSearchParams<{ id: string }>();

    return (
        <SafeAreaView>
            <Stack.Screen options={{ title: `${id}` }} />
            <Text>ClubPage</Text>
            <Link href={{ pathname: `/club/faq/${id}` }}>
                <Text>FAQ</Text>
            </Link>
        </SafeAreaView>
    );
};

export default ClubPage;
