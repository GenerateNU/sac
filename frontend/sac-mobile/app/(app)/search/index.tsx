import React, { useEffect, useState } from 'react';
import { Dimensions, ImageBackground, Platform, SafeAreaView, ScrollView, StatusBar, StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { SearchBar } from '../search/_components/searchbar';
import { ClubCard } from '@/components/club-card';
import { LeftArrow } from '@/components/left-arrow';
import LinearGradient from 'react-native-linear-gradient';
import { Button } from '@/components/button';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { router } from 'expo-router';
import Animated, { useAnimatedStyle, useSharedValue, withSpring, withTiming } from 'react-native-reanimated';
import { queryClubs } from '@/services/club';
import { Club } from '@/types/club';
import { ToggleButton } from './_components/togglebutton';

const SearchPageStyles = StyleSheet.create({
    // Because the bottom view's top border is rounded, we do these two styles to ensure that there's no gap in the space
    // that is left exposed because of the roundedness.
    topView: {
        height: Dimensions.get("screen").height * .2 + 24
    },
    bottomView: {
        height: Dimensions.get("screen").height * .8,
        marginTop: -24
    }
});

type SearchState = 'initialized' | 'loading' | 'done';

const SearchPage = () => {
    const [searchBarFocused, setSearchBarFocused] = useState(false);

    const [searchState, setSearchState] = useState<SearchState>("initialized");
    const [searchQuery, setSearchQuery] = useState("");

    const [clubs, setClubs] = useState<Club[]>([]);
    
    useEffect(() => {
        if (searchQuery == "") { return; }
        
        setSearchState("loading");
        queryClubs(searchQuery).then((clubs) => {
            setClubs(clubs)
            setSearchState("done");
        }).catch((x) => {
            console.log(":(")
        })
    }, [searchQuery])

    let bottomView: JSX.Element;
    
    if (searchState === 'initialized') {
        bottomView = (
            <Text>Search!!!</Text>
        );
    } else if (searchState === 'loading') {
        bottomView = (
            <Text>Loading...</Text>
        );
    } else {
        bottomView = (
            <ScrollView className="w-full px-[3%] flex flex-col" onScroll={(evt) => {setSearchBarFocused(false)}}>
                {clubs.map((val, ind, array) => {
                    return <ClubCard className="mb-4" club={val}/>
                })}
            </ScrollView>);
    }

    return (
        <SafeAreaView className="bg-red-500 w-full h-full">
            <View className="w-full" style={SearchPageStyles.topView}>
                <ImageBackground className="w-full h-full flex flex-row justify-center items-center" resizeMode="cover" source={require("@/assets/images/campus.png")} >
                     <Text className="font-bold text-white z-40 text-2xl drop-shadow-lg">Search for</Text>
                     <ToggleButton className="z-40 ml-3 h-8 w-[55%]" left="Clubs" right="Events" />
                </ImageBackground>
                {/* Overlay for image background */}
                <LinearGradient colors={["rgba(0,0,0,0.25)", "rgba(0,0,0,0.7)"]} className="w-full h-full absolute left-0 top-0 bottom-0 right-0"></LinearGradient>
            </View>
            <View className="z-10 w-full bg-white rounded-t-3xl flex flex-col items-center" style={SearchPageStyles.bottomView}>
                <View className="flex flex-row w-full h-12 mt-[-24px] px-[5%]">
                    <Button
                        onPress={() => {
                            console.log('Back Button Pressed');
                            router.back();
                        }}
                        size={'icon'}
                        variant={'icon'}
                        className="flex items-center justify-center bg-[#C8102E] rounded-full">
                        <MaterialCommunityIcons
                            name="chevron-left"
                            size={32}
                            color="rgb(255 255 255)"
                        />
                    </Button>
                    <SearchBar 
                        className='ml-3 flex-grow'
                        onGainFocus={() => {
                            setSearchBarFocused(true);
                        }}
                        onLoseFocus={() => {
                            setSearchBarFocused(false);
                        }}
                        onTextChange={(text) => {
                            setSearchQuery(text);
                        }}
                        focused={searchBarFocused}/>
                </View>
                { bottomView }
            </View>
            
        </SafeAreaView>
    );
};

export default SearchPage;
