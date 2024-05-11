import { useRef, useState } from 'react';
import {
    Pressable,
    Text,
    TextInput,
    TouchableOpacity,
    View,
    ViewProps
} from 'react-native';

import { router } from 'expo-router';

import { MaterialCommunityIcons } from '@expo/vector-icons';

import SearchIcon from '@/assets/images/svg/searchNew.svg';
import { Button } from '@/components/button';

export interface SearchBarProps extends ViewProps {
    searchBoxEditable: boolean;
    /** Will only be shown if searchBox is not editable. */
    searchBoxText?: string;
    focusOnLayout?: boolean;
    onSearchBoxPressed?: () => 'Focus' | undefined;
    onSearchSubmitted?: (query: string) => void;
    cancelOrBackButton?: 'Cancel' | 'Back';
}

export const SearchBar = (p: SearchBarProps) => {
    const [text, setText] = useState('');
    const textInputRef: React.Ref<TextInput> = useRef(null);

    // FIXME: needs to filter, users cannot search with '/' in the query
    const backButton = (
        <Button
            onPress={() => {
                router.back();
            }}
            size={'icon'}
            variant={'icon'}
            className="flex items-center w-2 mr-4 justify-center"
        >
            <MaterialCommunityIcons
                name="chevron-left"
                size={32}
                color="#1a1a1a"
            />
        </Button>
    );

    const cancelButton = (
        <TouchableOpacity
            onPress={() => {
                console.log('Cancel button pressed.');

                router.back();
            }}
            className="ml-4 flex flex-row justify-center items-center"
        >
            <Text className="text-center">Cancel</Text>
        </TouchableOpacity>
    );

    return (
        <View {...p} className="flex flex-row w-full h-12 items-center">
            {p.cancelOrBackButton === 'Back' ? backButton : <></>}
            <Pressable
                className="flex-grow"
                onPress={() => {
                    if (
                        p.onSearchBoxPressed &&
                        p.onSearchBoxPressed() === 'Focus'
                    ) {
                        textInputRef.current?.focus();
                    }
                }}
            >
                <View
                    {...p}
                    className="flex-grow bg-white h-12 px-4 flex flex-row rounded-lg border-zinc-200 border-2 items-center"
                >
                    <SearchIcon className="w-8 h-8" />
                    <TextInput
                        className={'ml-2 h-8 flex-grow'}
                        multiline={false}
                        editable={p.searchBoxEditable}
                        placeholder={'Search'}
                        ref={textInputRef}
                        value={
                            !p.searchBoxEditable && p.searchBoxText
                                ? p.searchBoxText
                                : text
                        }
                        onChangeText={setText}
                        onSubmitEditing={() => {
                            if (p.onSearchSubmitted) {
                                p.onSearchSubmitted(text);
                            }
                        }}
                        onLayout={() => {
                            // FIXME: works when going from discover, but not when coming back from results or results/expanded
                            if (
                                p.focusOnLayout !== undefined &&
                                p.focusOnLayout
                            ) {
                                textInputRef.current?.focus();
                            }
                        }}
                    />
                </View>
            </Pressable>
            {p.cancelOrBackButton === 'Cancel' ? cancelButton : <></>}
        </View>
    );
};
