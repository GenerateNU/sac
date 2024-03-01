import React, { useState } from 'react';
import {View, ScrollView, Text, StyleSheet} from 'react-native'; 
import { Dropdown } from 'react-native-element-dropdown';

// Library Component
// for more customization --> https://www.npmjs.com/package/react-native-element-dropdown?activeTab=code

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
        search
        maxHeight={300}
        labelField="label"
        valueField="value"
        placeholder={!isFocus ? props.placeholder : 'Select Year'}
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
    height: 80,
  },
  dropdown: {
    height: '85%',
    borderColor: 'black',
    borderWidth: 0.5,
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
    borderRadius: 12,
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