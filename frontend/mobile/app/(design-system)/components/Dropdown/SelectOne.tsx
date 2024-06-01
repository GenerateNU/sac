import React, { useState } from 'react';
import { FlatList, StyleSheet, TouchableOpacity } from 'react-native';

import { faAngleDown, faAngleUp } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Item } from '@/types/item';

import { Colors } from '../../shared/colors';
import { Box } from '../Box/Box';
import { Text } from '../Text/Text';

interface SelectOneProps {
    data: Item[];
    placeholder: string;
    title?: string;
    onChange: (value: Item) => void;
}

export const SelectOne: React.FC<SelectOneProps> = ({
    data,
    placeholder,
    title,
    onChange
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

    let dropdownHeight = data.length * 50.5;
    if (data.length >= 5) {
        dropdownHeight = 252.5;
    }

    return (
        <Box style={styles.container}>
            {title && <Text paddingBottom="xs">{title}</Text>}
            <TouchableOpacity
                style={styles.box}
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
        borderRadius: 12
    },
    box: {
        backgroundColor: 'white',
        width: '100%',
        borderColor: Colors.gray,
        borderWidth: 1,
        justifyContent: 'center',
        paddingHorizontal: 10,
        borderRadius: 10,
        height: 52
    },
    contentContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        paddingHorizontal: 5
    },
    dropdown: {
        backgroundColor: 'white',
        borderColor: Colors.gray,
        borderWidth: 1,
        width: '100%',
        borderRadius: 10,
        marginTop: 7
    },
    item: {
        padding: 15
    },
    divider: {
        height: 1,
        backgroundColor: Colors.gray
    }
});
