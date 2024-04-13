import React from 'react';
import { Text, StyleSheet, ScrollView, SafeAreaView, View } from 'react-native';
import { TextInput } from 'react-native-gesture-handler';
import { Button } from '@/components/button';

const Search = () => {
    return (
        <SafeAreaView>
            <ScrollView className='px-[8%] pb-[9%] w-full'>
                <Text className='text-5xl font-bold text-black'>Search</Text>
                <View className='flex flex-row gap-4'>
                    <TextInput placeholder="Search" className='w-48 p-2 mt-2 border border-black rounded-lg' />
                    <Button>Search</Button>
                </View>
            </ScrollView>
        </SafeAreaView>
    );
};

export default Search;