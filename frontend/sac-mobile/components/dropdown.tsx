import React, { useState } from 'react';
import { DimensionValue, StyleSheet, Text, View } from 'react-native';
import { Dropdown } from 'react-native-element-dropdown';

import { Item } from '@/types/item';

// Library Component
// https://www.npmjs.com/package/react-native-element-dropdown?activeTab=code

interface DropdownProps {
    title: string;
    item: Array<Item>; // list of dropdown items
    placeholder: string; // placeholder
    onChangeText: (...event: any[]) => void;
    value: Item;
    onSubmitEditing: () => void;
    search?: boolean; // true to enable search
    height?: DimensionValue;
    error?: boolean;
}

export const DropdownComponent = (props: DropdownProps) => {
    const [isFocus, setIsFocus] = useState(false);
    const borderColor = props.error ? 'red' : 'black';
    const borderWidth = props.error ? 1 : 0.5;

    const styles = StyleSheet.create({
        container: {
            height: props.height || 78
        },
        dropdown: {
            height: 50,
            borderColor: borderColor,
            borderWidth: borderWidth,
            borderRadius: 12,
            paddingHorizontal: '5%'
        },
        placeholderStyle: {
            fontSize: 14,
            color: '#CDCBCB',
            borderRadius: 12
        },
        selectedTextStyle: {
            fontSize: 14
        },
        inputSearchStyle: {
            height: 40,
            fontSize: 14,
            borderRadius: 11,
            marginLeft: 8
        },
        containerStyle: {
            borderRadius: 12,
            marginTop: '2%',
            borderColor: '#CDCBCB',
            overflow: 'hidden',
            borderWidth: 1
        },
        itemTextStyle: {
            fontSize: 14,
            paddingHorizontal: '3.5%'
        },
        itemContainerStyle: {
            borderBottomWidth: 1,
            borderColor: '#CDCBCB'
        },
        isFocus: {
            borderColor: 'black'
        }
    });

    return (
        <View style={styles.container}>
            <Text className="pb-[2%]">{props.title}</Text>
            <Dropdown
                style={[styles.dropdown, isFocus && styles.isFocus]}
                placeholderStyle={styles.placeholderStyle}
                selectedTextStyle={styles.selectedTextStyle}
                inputSearchStyle={styles.inputSearchStyle}
                containerStyle={styles.containerStyle}
                itemTextStyle={styles.itemTextStyle}
                itemContainerStyle={styles.itemContainerStyle}
                data={props.item}
                search={props.search || false}
                maxHeight={300}
                labelField="label"
                valueField="value"
                placeholder={props.placeholder}
                searchPlaceholder="Search"
                onFocus={() => setIsFocus(true)}
                onBlur={() => setIsFocus(false)}
                onChange={props.onChangeText}
                value={props.value}
            />
        </View>
    );
};
