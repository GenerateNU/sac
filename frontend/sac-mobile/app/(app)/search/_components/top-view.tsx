import { Button } from "@/components/button";
import { MaterialCommunityIcons } from "@expo/vector-icons";
import { router } from "expo-router";
import { useEffect, useRef, useState } from "react";
import { Text, GestureResponderEvent, TextInput, TouchableOpacity, TouchableWithoutFeedback, View, ViewProps } from "react-native";

import SearchIcon from '@/assets/images/svg/search.svg';

export interface SearchTopViewProps extends ViewProps {
    onSearchBoxGainFocus?: () => void,
    onSearchBoxLoseFocus?: () => void,
    onSearchBoxTextChange?: (text: string) => void,
    active: boolean
}

export const SearchTopView = (p: SearchTopViewProps) => {
    const [text, onChangeText] = useState("");
    const textInputRef: React.Ref<TextInput> = useRef(null);

    useEffect(() => {
        if (p.onSearchBoxTextChange) {
            p.onSearchBoxTextChange(text)
        }
    }, [text])

    const inactiveElements = (
        <>
            <Button
                onPress={() => {
                    console.log('Back Button Pressed');
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
            <TouchableWithoutFeedback
                onPress={(evt: GestureResponderEvent) => {
                    if (p.onSearchBoxGainFocus) {
                        p.onSearchBoxGainFocus()
                    }
                }}>
                <View {...p} className="flex-grow bg-white h-12 px-4 flex flex-row rounded-lg border-zinc-200 border-2 items-center">
                    <SearchIcon className="w-8 h-8"></SearchIcon>
                    <TextInput
                        className={"ml-2 h-8 flex-grow"}
                        multiline={false}
                        editable={false}
                        placeholder={"Search"}
                        value={text}/>
                </View>
            </TouchableWithoutFeedback>
        </>
    );

    const activeElements = (
        <>
            <TouchableWithoutFeedback
                onPress={(evt: GestureResponderEvent) => {
                    textInputRef.current?.focus();
                }}>
                <View {...p} className="flex-grow bg-white h-12 px-4 flex flex-row rounded-lg border-zinc-200 border-2 items-center">
                    <SearchIcon className="w-8 h-8"></SearchIcon>
                    <TextInput ref={textInputRef}
                        className={"ml-2 h-8 flex-grow"}
                        multiline={false}
                        onBlur={(e: any) => {
                            if (p.onSearchBoxLoseFocus) {
                                p.onSearchBoxLoseFocus()
                            }
                        }}
                        onFocus={(e: any) => {
                            if (p.onSearchBoxGainFocus) {
                                p.onSearchBoxGainFocus()
                            }
                        }}
                        onLayout={() => {
                            if (p.active) {
                                textInputRef.current?.focus();
                            }
                        }}
                        placeholder={"Search"}
                        onChangeText={onChangeText}
                        value={text}/>
                </View>
            </TouchableWithoutFeedback>
            <TouchableOpacity
                onPress={() => {
                    textInputRef.current?.blur();

                    if (p.onSearchBoxLoseFocus) {
                        p.onSearchBoxLoseFocus();
                    }
                }}
                className="ml-4 flex flex-row justify-center items-center">
                <Text className="text-center">Cancel</Text>
            </TouchableOpacity>
        </>
    )

    return (
        <View {...p} className="flex flex-row w-full h-12 items-center">
            {p.active ? activeElements : inactiveElements }
        </View>);
}