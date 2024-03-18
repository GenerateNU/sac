import React, {useEffect, useState} from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { MultiSelect } from 'react-native-element-dropdown';
import { Item } from '@/types/item';

interface MultiSelectProps {
    title: string;
    item: Array<Item>;
    placeholder: string;
    onSubmitEditing: () => void;
    search?: boolean;
    error?: boolean;
    onChange: (selectedItems: Item[]) => void;
    maxSelect: number;
}

const MultiSelectComponent = (props: MultiSelectProps) => {
    const borderColor = props.error ? 'red' : 'black';
    const borderWidth = props.error ? 1 : 0.5;
    const marginBottom = props.error ? '0.5%' : '2%';
    const [selected, setSelected] = useState(Array<Item>);

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
          marginBottom: marginBottom,
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
            maxSelect={props.maxSelect}
            placeholder={props.placeholder}
            searchPlaceholder="Search..."
            value={selected}
            onChange={item => {
              setSelected(item);
              props.onChange(item);
            }}
            selectedStyle={styles.selectedStyle}
        />
      </View>
    );
};

export default MultiSelectComponent;