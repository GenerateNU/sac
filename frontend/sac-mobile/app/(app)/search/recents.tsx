import { router } from 'expo-router';
import React from 'react';
import { Pressable, SafeAreaView, Text, TouchableOpacity, View } from 'react-native';

import { SearchBar } from '@/app/(app)/search/_components/search-bar';
import SearchIcon from '@/assets/images/svg/search.svg';

const Recents = () => {
    const searchRecents = [
        'UI/UX Design',
        "Women's Club Volleyball",
        'NU Italian Club',
        'Generate'
    ].map((val, ind, arr) => {
        return (
            <Pressable
                onPress={(evt) => {
                    router.push(`/search/results/${val}`);
                }}
            >
                <View className="w-full flex flex-row mb-4 items-center">
                    <SearchIcon className="w-5 h-5 mr-2"></SearchIcon>
                    <Text className="text-base flex-grow">{val}</Text>
                    <TouchableOpacity
                        onPress={(evt) => {
                            console.log('Recents removed.');
                        }}
                    >
                        {/* <MaterialCommunityIcons
                            name="close" size={20} className="w-5 h-5">
                        </MaterialCommunityIcons> */}
                    </TouchableOpacity>
                </View>
            </Pressable>
        );
    });
    return (
        <SafeAreaView className="w-full h-full bg-[#fafafa] pt-[50px] flex flex-col px-[5%]">
            <SearchBar
                searchBoxEditable={true}
                cancelOrBackButton="Cancel"
                onSearchBoxPressed={() => {
                    return 'Focus';
                }}
                onSearchSubmitted={(text) => {
                    router.push(`/search/results/${text}`);
                }}
                focusOnLayout={true}
            />
            <View className="flex flex-col mt-6">
                <Text className="font-bold text-base mb-2">Recents</Text>
                {searchRecents}
            </View>
        </SafeAreaView>
    );
};

export default Recents;
