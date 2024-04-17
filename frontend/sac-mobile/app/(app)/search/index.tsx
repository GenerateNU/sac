import React, { useEffect, useState } from 'react';
import { SafeAreaView, ScrollView, Text } from 'react-native';
import { ClubCard, ClubCardNoTags } from './_components/club-card';
import { queryClubs } from '@/services/club';
import { Club } from '@/types/club';
import { SearchTopView } from './_components/top-view';
import { EventCard, HorizontalEventCard } from './_components/event-card';
import { TopSearch } from './_components/top-search';
import SearchHomeView from './_components/home-view';
import SearchRecents from './_components/recents';
import SearchResult from './_components/search-result';

type SearchQueryState = "NoQuery" | "Loading" | "Done";

const SearchPage = () => {
    const [searchBarActive, setSearchBarActive] = useState(false);
    const [searchQueryState, setSearchQueryState] = useState<SearchQueryState>("NoQuery");
    const [searchQuery, setSearchQuery] = useState("");

    const [clubs, setClubs] = useState<Club[]>([]);
    
    useEffect(() => {
        if (searchQuery == "") {
            setSearchQueryState("NoQuery");
            return;
        } else {
            setSearchQueryState("Loading");
            queryClubs(searchQuery).then((clubs) => {
                // FIXME: major race condition - try making a search, then deleting/clearing it
                // before its loaded.
                setClubs(clubs)
                setSearchQueryState("Done");                
            }).catch((x) => {
                console.log(":(")
            })
        }
    }, [searchQuery])

    // useEffect(() => {
    //     if (searchQuery == "") { return; }
        
    //     queryClubs(searchQuery).then((clubs) => {
    //         setClubs(clubs)
    //         setSearchState("done");
    //     }).catch((x) => {
    //         console.log(":(")
    //     })
    // }, [searchQuery])

    // let bottomView: JSX.Element;
    
    /*if (searchState === 'initialized') {
        bottomView = (
            <Text>Search!!!</Text>
        );
    } else if (searchState === 'loading') {
        bottomView = (
            <Text>Loading...</Text>
        );
    } else {
        bottomView = (
            <ScrollView className="w-full px-[3%] flex flex-col" onScroll={(evt) => {setSearchBarActive(false)}}>
                {clubs.map((val, ind, array) => {
                    return <ClubCard className="mb-4" club={val}/>
                })}
            </ScrollView>);
    }
*/
    let bottomElement: JSX.Element = <></>;

    if (!searchBarActive && searchQueryState == "NoQuery") {
        bottomElement = <SearchHomeView/>
    } else if (!searchBarActive && searchQueryState == "Loading") {
        bottomElement = <Text>Loading</Text>
    } else if (!searchBarActive && searchQueryState == "Done") {
        bottomElement = <Text>Clubs shown</Text>
    } else if (searchBarActive && searchQueryState == "NoQuery") {
        bottomElement = <SearchRecents/>
    } else if (searchBarActive && searchQueryState == "Loading") {
        bottomElement = <Text>Loading</Text>
    } else if (searchBarActive && searchQueryState == "Done") {
        bottomElement = <SearchResult onClubViewAllPressed={() => {console.log("XXX")}}/>
    }

    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
            <SearchTopView 
                active={searchBarActive} 
                onSearchBoxGainFocus={() => { setSearchBarActive(true) }}
                onSearchBoxLoseFocus={() => { setSearchBarActive(false) }}
                onSearchBoxTextChange={(txt) => { setSearchQuery(txt) }}/>
            {bottomElement}
        </SafeAreaView>
    );
};

export default SearchPage;
