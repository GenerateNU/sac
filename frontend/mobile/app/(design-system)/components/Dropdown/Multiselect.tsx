import React, { useState } from 'react';
import {
    FlatList,
    StyleSheet,
    TextInput,
    TouchableOpacity
} from 'react-native';

import {
    faAngleDown,
    faAngleUp,
    faXmark
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Item } from '@/types/item';

import { Border } from '../../shared/border';
import { Colors } from '../../shared/colors';
import { Spacing } from '../../shared/spacing';
import { Box } from '../Box/Box';
import { Text } from '../Text/Text';

interface MultiSelectProps {
    data: Item[];
    placeholder: string;
    title: string;
    onChange: (values: Item[]) => void;
    limit?: number;
    error?: boolean;
}

export const MultiSelect: React.FC<MultiSelectProps> = ({
    data,
    placeholder,
    title,
    onChange,
    limit = 10,
    error = false
}) => {
    const [selectedItems, setSelectedItems] = useState<Item[]>([]);
    const [dropdownVisible, setDropdownVisible] = useState(false);
    const [searchTerm, setSearchTerm] = useState('');
    const [dropdownHeight, setDropdownHeight] = useState<number>(
        (data.length + 1) * 51.5
    );

    const calculateDropdownHeight = (itemsCount: number) => {
        let height = (itemsCount + 1) * 52;
        if (itemsCount >= 5) {
            height = 255;
        }
        return height;
    };

    const handleSearch = (text: string) => {
        const filteredData = text
            ? data.filter((item) =>
                  item.value.toLowerCase().includes(text.toLowerCase())
              )
            : data;
        const newDropdownHeight = calculateDropdownHeight(filteredData.length);
        setDropdownHeight(newDropdownHeight);
        setSearchTerm(text);
    };

    const toggleItem = (item: Item) => {
        const index = selectedItems.findIndex(
            (selectedItem) => selectedItem.key === item.key
        );
        if (index !== -1) {
            setSelectedItems((prevSelectedItems) => {
                const newSelectedItems = prevSelectedItems.filter(
                    (selectedItem) => selectedItem.key !== item.key
                );
                onChange(newSelectedItems);
                return newSelectedItems;
            });
        } else {
            if (selectedItems.length < limit) {
                setSelectedItems((prevSelectedItems) => {
                    const newSelectedItems = [...prevSelectedItems, item];
                    onChange(newSelectedItems);
                    return newSelectedItems;
                });
            }
        }
    };

    const renderItem = ({ item, index }: { item: Item; index: number }) => {
        const isSelected = selectedItems.some(
            (selectedItem) => selectedItem.key === item.key
        );
        const isLastItem = index === filteredData.length - 1;

        return (
            <Box>
                <TouchableOpacity
                    style={[styles.item, isSelected && styles.selectedItem]}
                    onPress={() => toggleItem(item)}
                >
                    <Text style={styles.itemText}>{item.value}</Text>
                </TouchableOpacity>
                {!isLastItem && <Box style={styles.divider} />}
            </Box>
        );
    };

    const filteredData = searchTerm
        ? data.filter((item) =>
              item.value.toLowerCase().includes(searchTerm.toLowerCase())
          )
        : data;

    const handleComponentPress = () => {
        setDropdownVisible(!dropdownVisible);
    };

    return (
        <TouchableOpacity
            style={[
                styles.container,
                { height: dropdownVisible ? dropdownHeight : undefined }
            ]}
            onPress={handleComponentPress}
        >
            <Text style={styles.title}>{title}</Text>
            <Box
                style={[
                    styles.box,
                    { borderColor: error ? Colors.red : Colors.gray }
                ]}
            >
                <Box style={styles.contentContainer}>
                    <Box
                        width={selectedItems.length === 0 ? '0%' : '80%'}
                        marginVertical="s"
                        alignItems="center"
                        flexDirection="row"
                        flexWrap="wrap"
                    >
                        {selectedItems.map((item) => (
                            <TouchableOpacity
                                key={item.key}
                                style={styles.tag}
                                onPress={() => toggleItem(item)}
                            >
                                <Text style={styles.tagText}>{item.value}</Text>
                                <FontAwesomeIcon
                                    icon={faXmark}
                                    color={Colors.black}
                                />
                            </TouchableOpacity>
                        ))}
                    </Box>
                    <TouchableOpacity
                        onPress={handleComponentPress}
                        style={styles.input}
                    >
                        {selectedItems.length === 0 && (
                            <Text style={styles.placeholder}>
                                {placeholder}
                            </Text>
                        )}
                        <FontAwesomeIcon
                            icon={dropdownVisible ? faAngleUp : faAngleDown}
                            color={Colors.gray}
                        />
                    </TouchableOpacity>
                </Box>
            </Box>
            {dropdownVisible && (
                <Box style={styles.dropdown}>
                    <TextInput
                        placeholder="Search..."
                        style={styles.searchInput}
                        value={searchTerm}
                        onChangeText={handleSearch}
                    />
                    <FlatList
                        data={filteredData}
                        renderItem={renderItem}
                        keyExtractor={(item) => item.key}
                    />
                </Box>
            )}
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    container: {
        width: '100%',
        borderRadius: Border.base
    },
    box: {
        backgroundColor: 'white',
        width: '100%',
        borderWidth: 1,
        justifyContent: 'center',
        paddingHorizontal: Spacing.xs,
        borderRadius: Spacing.xs
    },
    contentContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'space-between'
    },
    dropdown: {
        backgroundColor: 'white',
        borderColor: Colors.gray,
        borderWidth: 1,
        borderRadius: Border.base,
        marginTop: Spacing.xxs
    },
    title: {
        fontSize: 18,
        marginBottom: Spacing.xs
    },
    item: {
        flexDirection: 'row',
        alignItems: 'center',
        paddingHorizontal: Spacing.s,
        paddingVertical: Spacing.m
    },
    selectedItem: {
        backgroundColor: Colors.gray
    },
    itemText: {
        flex: 1,
        marginLeft: Spacing.xxs
    },
    input: {
        flexDirection: 'row',
        alignItems: 'center',
        paddingVertical: Spacing.m,
        paddingHorizontal: Spacing.xxs,
        justifyContent: 'space-around'
    },
    placeholder: {
        flex: 1,
        color: Colors.gray
    },
    tag: {
        backgroundColor: Colors.gray,
        borderRadius: Border.base,
        paddingVertical: Spacing.xxs,
        paddingHorizontal: Spacing.xs,
        marginRight: Spacing.xxs,
        marginBottom: Spacing.xxs,
        flexDirection: 'row',
        alignItems: 'center',
        gap: 2
    },
    tagText: {
        color: Colors.black
    },
    searchInput: {
        height: 40,
        borderColor: Colors.gray,
        borderWidth: 1,
        borderRadius: Border.base,
        paddingHorizontal: Spacing.s,
        margin: Spacing.xxs
    },
    selectButtonText: {
        color: 'white'
    },
    divider: {
        height: 1,
        backgroundColor: Colors.gray
    }
});

export default MultiSelect;
