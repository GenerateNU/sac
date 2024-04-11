import { useEffect, useRef, useState } from "react";
import { GestureResponderEvent, TextInput, Touchable, TouchableOpacity, TouchableWithoutFeedback, View, ViewProps } from "react-native";

import SearchIcon from '@/assets/images/svg/search.svg';
import { Svg } from "react-native-svg";
import { MaterialCommunityIcons } from "@expo/vector-icons";

const PLACEHOLDERS = [
    "uh",
    "uhhhh",
    "ummmmmmm",
    "idk"
];

export interface SearchBarProps extends ViewProps {
    onSearchBarGainFocus?: () => void,
    onSearchBarLoseFocus?: (text: string) => void,
    focused: boolean
}

// @refresh reset
export const SearchBar = (p: SearchBarProps) => {
    const [placeholderIdx, setPlaceholderIdx] = useState(0);
    useEffect(() => {
        const interval = setInterval(() => {
            setPlaceholderIdx((placeholderIdx + 1) % PLACEHOLDERS.length);
        }, 3000);

        return () => { clearInterval(interval); }
    });

    const [text, onChangeText] = useState("");
    const textInputRef: React.Ref<TextInput> = useRef(null);

    if (p.focused) {
        textInputRef.current?.focus();
    } else {
        textInputRef.current?.blur();
    }

    const clearButton = p.focused ? (
        <TouchableOpacity
            onPress={(evt) => {
                onChangeText("")
            }}>
            <MaterialCommunityIcons
                name="close"
                size={24}
                color="#9d9d9d"
            />
        </TouchableOpacity>
        
    ) : <></>;

    // This is wrapped in a touchable so the textinput will focus when any part of the
    // searchbar is pressed, not just the text input.
    return (
        <TouchableWithoutFeedback
            onPress={(evt: GestureResponderEvent) => {
                textInputRef.current?.focus();
            }}>
            <View {...p} className="bg-white h-12 px-4 flex flex-row rounded-lg border-zinc-200 border-2 items-center">
                <SearchIcon className="w-8 h-8"></SearchIcon>
                <TextInput ref={textInputRef}
                    className={"ml-2 h-8 flex-grow"}
                    multiline={false}
                    onBlur={(e: any) => {
                        if (p.onSearchBarLoseFocus) {
                            p.onSearchBarLoseFocus(text);
                        }
                    }}
                    onFocus={(e: any) => {
                        if (p.onSearchBarGainFocus) {
                            p.onSearchBarGainFocus()
                        }
                    }}
                    placeholder={PLACEHOLDERS[placeholderIdx]}
                    onChangeText={onChangeText}
                    value={text}/>
                {clearButton}
            </View>
        </TouchableWithoutFeedback>
    )
}