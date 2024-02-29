import React, { useState } from 'react';
import {View, ScrollView, Text, StyleSheet} from 'react-native'; 
import { Dropdown } from 'react-native-element-dropdown';

type Item = {
    label: string, 
    value: string,
}
type ListOfItem = {
    title: string, 
    item: Array<Item>,
    placeholder: string,
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

  return (
    <ScrollView style={styles.container}>
        <Text>{props.title}</Text>
      {renderLabel()}
      <Dropdown
        style={[styles.dropdown, isFocus && { borderColor: 'black' }]}
        placeholderStyle={styles.placeholderStyle}
        selectedTextStyle={styles.selectedTextStyle}
        inputSearchStyle={styles.inputSearchStyle}
        data={props.item}
        search
        maxHeight={300}
        labelField="label"
        valueField="value"
        placeholder={!isFocus ? props.placeholder : ''}
        searchPlaceholder="Search..."
        value={value}
        onFocus={() => setIsFocus(true)}
        onBlur={() => setIsFocus(false)}
        onChange={item => {
          setValue(value);
          setIsFocus(false);
        }}
      />
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    backgroundColor: 'white',
    height: '20%',
  },
  dropdown: {
    height: '80%',
    borderColor: 'gray',
    borderWidth: 0.5,
    borderRadius: 8,
    paddingHorizontal: '3%',
  },
  placeholderStyle: {
    fontSize: 14,
    color: 'gray'
  },
  selectedTextStyle: {
    fontSize: 14,
  },
  inputSearchStyle: {
    height: 40,
    fontSize: 14,
  },
});