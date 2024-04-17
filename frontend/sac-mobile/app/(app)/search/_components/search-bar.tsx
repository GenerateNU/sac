import { Button } from "@/components/button";
import { MaterialCommunityIcons } from "@expo/vector-icons";
import { router } from "expo-router";
import { useContext, useEffect, useRef, useState } from "react";
import { Text, GestureResponderEvent, TextInput, TouchableOpacity, TouchableWithoutFeedback, View, ViewProps, Pressable } from "react-native";

import SearchIcon from '@/assets/images/svg/search.svg';
import { SearchContext } from "../_layout";

export interface SearchBarProps extends ViewProps {
    searchBoxEditable: boolean,
    onSearchBoxPressed?: () => ("Focus" | undefined),
    onSearchSubmitted?: (query: string) => void,
    cancelOrBackButton?: "Cancel" | "Back",
}

export const SearchBar = (p: SearchBarProps) => {
    const [text, setText] = useState("");
    const textInputRef: React.Ref<TextInput> = useRef(null);


    const searchBoxMargin = p.cancelOrBackButton == "Cancel" ? "ml-2" : "mr-2";
    const backButton = (
        <Button
            onPress={() => {
                router.back();
            }}
            size={'icon'}
            variant={'icon'}
            className="flex items-center w-2 mr-4 justify-center">
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
                console.log("Cancel button pressed.");
                
                router.back();
            }}
            className="ml-4 flex flex-row justify-center items-center">
            <Text className="text-center">Cancel</Text>
        </TouchableOpacity>
    );

    return (
        <View {...p} className="flex flex-row w-full h-12 items-center">
            {p.cancelOrBackButton === "Back" ? backButton : <></>}
            <Pressable className="flex-grow"
                onPress={(evt: GestureResponderEvent) => {
                   if (p.onSearchBoxPressed && p.onSearchBoxPressed() == "Focus") {
                     textInputRef.current?.focus();
                   } 
                }}>
                <View {...p} className="flex-grow bg-white h-12 px-4 flex flex-row rounded-lg border-zinc-200 border-2 items-center">
                    <SearchIcon className="w-8 h-8"></SearchIcon>
                    <TextInput
                        className={"ml-2 h-8 flex-grow"}
                        multiline={false}     
                        editable={p.searchBoxEditable}
                        placeholder={"Search"}
                        ref={textInputRef}
                        value={text}
                        onChangeText={setText}
                        onSubmitEditing={(evt) => {
                            if (p.onSearchSubmitted) {
                                p.onSearchSubmitted(text)
                            }
                        }}/>
                </View>
            </Pressable>
            {p.cancelOrBackButton === "Cancel" ? cancelButton : <></>}
        </View>);
}