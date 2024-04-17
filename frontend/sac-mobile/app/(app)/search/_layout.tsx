import React, { createContext, useContext, useEffect, useState } from 'react';
import { Stack } from 'expo-router';
import { Club } from '@/types/club';
import { queryClubs } from '@/services/club';
import DiscoverPage from './discover';

export type QueryState = "NoQuery" | "Loading" | "Done";
export type SearchContextState = {
    query: string,
    queryState: QueryState,
    clubs: Club[],
}

// @refresh reset

export const SearchContext = createContext<SearchContextState>({
    query: "",
    queryState: "NoQuery",
    clubs: []
});

const SearchLayout = () => {
    const [query, setQuery] = useState("");
    const [queryState, setQueryState] = useState<QueryState>("NoQuery");
    const [clubs, setClubs] = useState<Club[]>([]);



    useEffect(() => {
        if (query == "") {
            setQueryState("NoQuery");
            return;
        } else {
            setQueryState("Loading");
            queryClubs(query).then((clubs) => {
                // FIXME: major race condition - try making a search, then deleting/clearing it
                // before its loaded.
                setClubs(clubs)
                setQueryState("Done");                
            }).catch((x) => {
                console.log(":(")
            })
        }
    }, [query])

    
    return (
        <SearchContext.Provider value={{
            query: query,
            queryState: queryState,
            clubs: clubs
        }}>
            <Stack screenOptions={{ headerShown: false }}>
                <Stack.Screen 
                    name="discover"/>
            </Stack>
        </SearchContext.Provider>
        
    );
};

export default SearchLayout;
