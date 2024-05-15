import { Text, TouchableOpacity, View, ViewProps } from 'react-native';

import { router } from 'expo-router';

import SearchIcon from '@/assets/images/svg/searchNew.svg';

export interface TopSearchProps extends ViewProps {
    search: string;
}

export const TopSearch = (p: TopSearchProps) => {
    return (
        <TouchableOpacity
            onPress={() => {
                router.push(`/(app)/search/results/${p.search}`);
            }}
        >
            <View
                {...p}
                className="elevation bg-white rounded-xl flex flex-row p-3 items-center"
            >
                <SearchIcon className="w-5 h-5" />
                <Text className="ml-2 text-xs">{p.search}</Text>
            </View>
        </TouchableOpacity>
    );
};
