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

import { Colors } from '../../shared/colors';
import { Box } from '../Box/Box';
import { Text } from '../Text/Text';

interface MultiSelectProps {
    data: Item[];
    placeholder: string;
    title: string;
    onChange: (values: Item[]) => void;
    limit?: number;
}

export const MultiSelect: React.FC<MultiSelectProps> = ({
    data,
    placeholder,
    title,
    onChange,
    limit = 10
}) => {
    const [selectedItems, setSelectedItems] = useState<Item[]>([]);
    const [dropdownVisible, setDropdownVisible] = useState(false);
    const [searchTerm, setSearchTerm] = useState('');

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
            style={styles.container}
            onPress={handleComponentPress}
        >
            <Text style={styles.title}>{title}</Text>
            <Box style={styles.box}>
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
                        onChangeText={setSearchTerm}
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
        borderRadius: 12
    },
    box: {
        backgroundColor: 'white',
        width: '100%',
        borderColor: Colors.gray,
        borderWidth: 1,
        justifyContent: 'center',
        paddingHorizontal: 10,
        borderRadius: 10
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
        width: '100%',
        borderRadius: 10,
        marginTop: 5
    },
    title: {
        fontSize: 18,
        marginBottom: 10
    },
    item: {
        flexDirection: 'row',
        alignItems: 'center',
        paddingHorizontal: 10,
        paddingVertical: 15
    },
    selectedItem: {
        backgroundColor: Colors.gray
    },
    itemText: {
        flex: 1,
        marginLeft: 5
    },
    input: {
        flexDirection: 'row',
        alignItems: 'center',
        paddingVertical: 15,
        paddingHorizontal: 5,
        justifyContent: 'space-around'
    },
    placeholder: {
        flex: 1,
        color: Colors.gray
    },
    tag: {
        backgroundColor: Colors.gray,
        borderRadius: 8,
        paddingVertical: 5,
        paddingHorizontal: 10,
        marginRight: 5,
        marginBottom: 5,
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
        borderRadius: 10,
        paddingHorizontal: 10,
        margin: 5
    },
    selectButtonText: {
        color: 'white'
    },
    divider: {
        height: 1,
        backgroundColor: Colors.gray
    }
});
