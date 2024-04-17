import React, { createContext, useContext, useEffect, useState } from 'react';
import { Stack } from 'expo-router';
import { Club } from '@/types/club';
import { queryClubs } from '@/services/club';
import DiscoverPage from './discover';


// @refresh reset
const SearchLayout = () => {    
    return (
        <Stack screenOptions={{ headerShown: false }}>
            <Stack.Screen 
                name="discover"/>
        </Stack>        
    );
};

export default SearchLayout;
