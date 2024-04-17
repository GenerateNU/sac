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


const DiscoverPage = () => {    
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
