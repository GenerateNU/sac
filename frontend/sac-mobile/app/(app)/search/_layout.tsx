import React, { createContext, useContext, useEffect, useState } from 'react';
import { Stack } from 'expo-router';
import { Club } from '@/types/club';
import { queryClubs } from '@/services/club';
import DiscoverPage from '../(tabs)/discover';

// @refresh reset
const SearchLayout = () => {    
    return (
        <Stack screenOptions={{ headerShown: false }}>
            <Stack.Screen 
                name="results/[query]"
                options={{headerShown: false}} />
            <Stack.Screen
                name="recents"
                options={{headerShown: false}} />
            <Stack.Screen
                name="results/expanded/[query]"
                options={{headerShown: false}} />
        </Stack>        
    );
};

export default SearchLayout;
