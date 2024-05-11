import React from 'react';
import {
    Pressable,
    SafeAreaView,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { router } from 'expo-router';

import { SearchBar } from '@/app/(app)/search/_components/search-bar';
import SearchIcon from '@/assets/images/svg/searchNew.svg';

const Recents = () => {
    const searchRecents = [
        'UI-UX Design',
        "Women's Club Volleyball",
        'NU Italian Club',
        'Generate'
    ].map((val) => {
        return (
            <Pressable
                onPress={() => {
                    router.push(`/search/results/${val}`);
                }}
            >
                <View className="w-full flex flex-row mb-4 items-center">
                    <SearchIcon className="w-5 h-5 mr-2" />
                    <Text className="text-base flex-grow">{val}</Text>
                    <TouchableOpacity
                        onPress={() => {
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
        <SafeAreaView className="w-full h-full bg-[#fafafa]">
            <View className="px-[5%]">
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
            </View>
        </SafeAreaView>
    );
};

export default Recents;
