import React, { useState } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { MultiSelect } from 'react-native-element-dropdown';
import { Item } from '@/types/item';

interface MultiSelectProps {
    title: string;
    item: Array<Item>; // list of dropdown items
    placeholder: string; // placeholder
    onSubmitEditing: () => void;
    search?: boolean; // true to enable search
    error?: boolean;
    value: Array<Item>;
    onChange?: () => void;
}

const MultiSelectComponent = (props: MultiSelectProps) => {
    const [selected, setSelected] = useState(Array<Item>);
    const borderColor = props.error ? 'red' : 'black';
    const borderWidth = props.error ? 1 : 0.5;

    const styles = StyleSheet.create({
        container: { 
            padding: 0,
        },
        itemContainerStyle: {
            borderBottomWidth: 1,
            borderColor: '#F0F0F0',
        },
        itemTextStyle: {
            fontSize: 14,
            paddingLeft: '2%',
        },
        containerStyle: {
            borderRadius: 14,
            marginTop: 6,
            height: 320,
        },
        dropdown: {
          height: 52,
          borderWidth: borderWidth,
          borderRadius: 12,
          paddingLeft: '5%', 
          paddingRight: '5%',
          marginBottom: '2%',
          borderColor: borderColor,
        },
        placeholderStyle: {
          fontSize: 14,
          color: '#CDCBCB',
        },
        selectedTextStyle: {
          fontSize: 14,
        },
        inputSearchStyle: {
          height: 45,
          fontSize: 14,
          borderRadius: 13,
          paddingLeft: 7,
        },
        selectedStyle: {
          borderRadius: 10,
          marginTop: '1%',
        },
      });

    return (
      <View style={styles.container}>
        <Text className="mb-[2%]">{props.title}</Text>
        <MultiSelect
            style={styles.dropdown}
            placeholderStyle={styles.placeholderStyle}
            selectedTextStyle={styles.selectedTextStyle}
            inputSearchStyle={styles.inputSearchStyle}
            containerStyle={styles.containerStyle}
            itemContainerStyle={styles.itemContainerStyle}
            itemTextStyle={styles.itemTextStyle}
            search={props.search}
            data={props.item}
            activeColor={'#DCDCDC'}
            labelField="label"
            valueField="value"
            placeholder={props.placeholder}
            searchPlaceholder="Search..."
            value={selected}
            onChange={(item : Item[]) => {
              setSelected(item);
            }}
            selectedStyle={styles.selectedStyle}
        />
      </View>
    );
};

export default MultiSelectComponent;