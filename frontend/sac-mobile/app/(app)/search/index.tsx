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

const SearchPage = () => {
    const [searchBarFocused, setSearchBarFocused] = useState(false);
    const [searchQuery, setSearchQuery] = useState("");

    useEffect(() => {
        if (searchQuery == "") { return; }
        console.log(`We are going to do search logic here! ${searchQuery}`)
    }, [searchQuery])

    return (
        <SafeAreaView className="bg-red-500 w-full h-full">
            <View className="w-full" style={SearchPageStyles.topView}>
                <ImageBackground className="w-full h-full flex flex-row justify-center items-center" resizeMode="cover" source={require("@/assets/images/campus.png")} >
                     <Text className="font-bold text-white z-40 text-2xl drop-shadow-lg">Search for</Text>
                </ImageBackground>
                {/* Overlay for image background */}
                <LinearGradient colors={["rgba(0,0,0,0.25)", "rgba(0,0,0,0.7)"]} className="w-full h-full absolute left-0 top-0 bottom-0 right-0"></LinearGradient>
            </View>
            <View className="z-10 w-full bg-white rounded-t-3xl" style={SearchPageStyles.bottomView}>
                <View className="flex flex-row w-full px-[5%] h-12 mt-[-24px]">
                    <Button
                        onPress={() => {
                            console.log('Back Button Pressed');
                            router.back();
                        }}
                        size={'icon'}
                        variant={'icon'}
                        className="flex items-center justify-center bg-[#C8102E] rounded-full"
                    >
                        <MaterialCommunityIcons
                            name="chevron-left"
                            size={32}
                            color="rgb(255 255 255)"
                        />
                    </Button>
                    <SearchBar 
                        className='ml-3 flex-grow'
                        onSearchBarGainFocus={() => {
                            setSearchBarFocused(true);
                        }}
                        onSearchBarLoseFocus={() => {
                            setSearchBarFocused(false);
                        }}
                        focused={searchBarFocused}/>
                </View>
            </View>
        </SafeAreaView>
    )
    // return (
    //     <SafeAreaView className="px-6 flex flex-col w-full h-full pt-2" style={notchPadding.padding}>
    //         <SearchBar onSearchBarLoseFocus={(text: string) => {
    //             setSearchQuery(text)
    //         }}/>
    //         <ScrollView className='mt-6 flex flex-col' showsVerticalScrollIndicator={false}>
    //             <ClubCard club={{
    //                 id: "cfc85717-4f51-43e7-aa4d-06f3ce748cb9",
    //                 created_at: new Date(),
    //                 updated_at: new Date(),
    //                 name: "Generate",
    //                 preview: "alright ig",
    //                 description: "im goated icl",
    //                 numMembers: 250,
    //                 isRecruiting: false,
    //                 recruitmentCycle: 'fall',
    //                 recruitmentType: 'unrestricted',
    //                 applicationLink: "https://google.com",
    //                 logo: "Logo.png"   
    //             }}/>
    //             <View className="bg-transparent w-full h-4" />
    //             <ClubCard club={{
    //                 id: "cfc85717-4f51-43e7-aa4d-06f3ce748cb9",
    //                 created_at: new Date(),
    //                 updated_at: new Date(),
    //                 name: "Generate",
    //                 preview: "alright ig",
    //                 description: "im goated icl",
    //                 numMembers: 250,
    //                 isRecruiting: false,
    //                 recruitmentCycle: 'fall',
    //                 recruitmentType: 'unrestricted',
    //                 applicationLink: "https://google.com",
    //                 logo: "Logo.png"   
    //             }}/>
    //             <View className="bg-transparent w-full h-4" />
    //             <ClubCard club={{
    //                 id: "cfc85717-4f51-43e7-aa4d-06f3ce748cb9",
    //                 created_at: new Date(),
    //                 updated_at: new Date(),
    //                 name: "Generate",
    //                 preview: "alright ig",
    //                 description: "im goated icl",
    //                 numMembers: 250,
    //                 isRecruiting: false,
    //                 recruitmentCycle: 'fall',
    //                 recruitmentType: 'unrestricted',
    //                 applicationLink: "https://google.com",
    //                 logo: "Logo.png"   
    //             }}/>
    //             <View className="bg-transparent w-full h-4" />
    //             <ClubCard club={{
    //                 id: "cfc85717-4f51-43e7-aa4d-06f3ce748cb9",
    //                 created_at: new Date(),
    //                 updated_at: new Date(),
    //                 name: "Generate",
    //                 preview: "alright ig",
    //                 description: "im goated icl",
    //                 numMembers: 250,
    //                 isRecruiting: false,
    //                 recruitmentCycle: 'fall',
    //                 recruitmentType: 'unrestricted',
    //                 applicationLink: "https://google.com",
    //                 logo: "Logo.png"   
    //             }}/>
    //             <View className="bg-transparent w-full h-4" />
    //             <ClubCard club={{
    //                 id: "cfc85717-4f51-43e7-aa4d-06f3ce748cb9",
    //                 created_at: new Date(),
    //                 updated_at: new Date(),
    //                 name: "Generate",
    //                 preview: "alright ig",
    //                 description: "im goated icl",
    //                 numMembers: 250,
    //                 isRecruiting: false,
    //                 recruitmentCycle: 'fall',
    //                 recruitmentType: 'unrestricted',
    //                 applicationLink: "https://google.com",
    //                 logo: "Logo.png"   
    //             }}/>
    //         </ScrollView>
            
    //     </SafeAreaView>
    // );
};

export default SearchPage;
