import React, { useState } from 'react';
import {View, ScrollView, Text, StyleSheet, DimensionValue} from 'react-native';
import { Dropdown } from 'react-native-element-dropdown';

// Library Component
// https://www.npmjs.com/package/react-native-element-dropdown?activeTab=code

type Item = {
    label: string, 
    value: string,
}
type ListOfItem = {
    title: string,
    item: Array<Item>,   // list of dropdown items
    placeholder: string, // placeholder
    onChangeText: (...event: any[]) => void;
    value: string;
    onSubmitEditing: () => void;
    search?: boolean; // true for enable search
    height: DimensionValue;
    error?: boolean;
}

export const DropdownComponent = (props: ListOfItem) => {
  const [value, setValue] = useState(null);
  const [isFocus, setIsFocus] = useState(false);

  const renderLabel = () => {
    if (value || isFocus) {
      return;
    }
    return null;
  };

  const borderColor = props.error ? 'red' : 'black';
  const borderWidth = props.error ? 1 : 0.5;

  const styles = StyleSheet.create({
    container: {
      backgroundColor: 'white',
      height: props.height,
    },
    dropdown: {
      height: '85%',
      borderColor: borderColor,
      borderWidth: borderWidth,
      borderRadius: 12,
      paddingHorizontal: '5%',
    },
    placeholderStyle: {
      fontSize: 14,
      color: '#CDCBCB',
      borderRadius: 12,
    },
    selectedTextStyle: {
      fontSize: 14,
    },
    inputSearchStyle: {
      height: 40,
      fontSize: 14,
      borderRadius: 11,
    },
    containerStyle: {
      borderRadius: 12,
      marginTop: '2%',
      borderColor: '#CDCBCB',
    },
    itemTextStyle: {
      fontSize: 14,
      paddingHorizontal: '3.5%',
      borderBottomColor: 'grey',
    },
    itemContainerStyle: {
      borderBottomWidth: 0.8,
      borderColor: '#CDCBCB',
    }
  });

  return (
    <ScrollView style={styles.container}>
        <Text className="pb-[2%]">{props.title}</Text>
      {renderLabel()}
      <Dropdown
        style={[styles.dropdown, isFocus && { borderColor: 'black' }]}
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
        placeholder={!isFocus ? props.placeholder : 'Select Year'}
        searchPlaceholder="Search"
        onFocus={() => setIsFocus(true)}
        onBlur={() => setIsFocus(false)}
        onChange={props.onChangeText}
        value={props.value}
      />
    </ScrollView>
  );
};
