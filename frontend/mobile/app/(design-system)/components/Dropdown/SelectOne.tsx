import React, { useState } from 'react';
import { FlatList, StyleSheet, TouchableOpacity } from 'react-native';

import { faAngleDown, faAngleUp } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Item } from '@/types/item';

import { Border } from '../../shared/border';
import { Colors } from '../../shared/colors';
import { Spacing } from '../../shared/spacing';
import { Box } from '../Box/Box';
import { Text } from '../Text/Text';

interface SelectOneProps {
    data: Item[];
    placeholder: string;
    title?: string;
    onChange: (value: Item) => void;
    error?: boolean;
}

export const SelectOne: React.FC<SelectOneProps> = ({
    data,
    placeholder,
    title,
    onChange,
    error = false
}) => {
    const [selectedItem, setSelectedItem] = useState<Item | null>(null);
    const [dropdownVisible, setDropdownVisible] = useState(false);

    const handleSelect = (item: Item) => {
        setSelectedItem(item);
        onChange(item);
        setDropdownVisible(false);
    };

    const renderItem = ({ item, index }: { item: Item; index: number }) => {
        const isLastItem = index === data.length - 1;
        return (
            <Box>
                <TouchableOpacity
                    style={styles.item}
                    onPress={() => handleSelect(item)}
                >
                    <Text>{item.value}</Text>
                </TouchableOpacity>
                {!isLastItem && <Box style={styles.divider} />}
            </Box>
        );
    };

    let dropdownHeight = data.length * 53;
    if (data.length >= 5) {
        dropdownHeight = 265;
    }

    return (
        <Box style={styles.container}>
            {title && <Text paddingBottom="xs">{title}</Text>}
            <TouchableOpacity
                style={[
                    styles.box,
                    { borderColor: error ? Colors.red : Colors.gray }
                ]}
                onPress={() => setDropdownVisible(!dropdownVisible)}
            >
                <Box style={styles.contentContainer}>
                    <Text
                        variant="body-1"
                        color={selectedItem === null ? 'gray' : 'black'}
                    >
                        {selectedItem ? selectedItem.value : placeholder}
                    </Text>
                    <FontAwesomeIcon
                        color={Colors.gray}
                        icon={dropdownVisible ? faAngleDown : faAngleUp}
                    />
                </Box>
            </TouchableOpacity>
            {dropdownVisible && (
                <Box style={[styles.dropdown, { height: dropdownHeight }]}>
                    <FlatList
                        data={data}
                        renderItem={renderItem}
                        keyExtractor={(item) => item.key}
                    />
                </Box>
            )}
        </Box>
    );
};

const styles = StyleSheet.create({
    container: {
        width: '100%',
        borderRadius: Border.medium
    },
    box: {
        backgroundColor: 'white',
        width: '100%',
        borderWidth: 1,
        justifyContent: 'center',
        paddingHorizontal: Spacing.s,
        borderRadius: Border.base,
        height: 52
    },
    contentContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        paddingHorizontal: Spacing.xxs
    },
    dropdown: {
        backgroundColor: 'white',
        borderColor: Colors.gray,
        borderWidth: 1,
        width: '100%',
        borderRadius: Border.base,
        marginTop: Spacing.xs
    },
    item: {
        padding: Spacing.m
    },
    divider: {
        height: 1,
        backgroundColor: Colors.gray
    }
});
