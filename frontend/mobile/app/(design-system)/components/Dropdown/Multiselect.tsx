import React, { useState } from 'react';
import {
    FlatList,
    StyleSheet,
    TextInput,
    TouchableOpacity,
    View
} from 'react-native';

import {
    faAngleDown,
    faAngleUp,
    faTimes
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import ErrorIcon from '@/app/(design-system)/assets/svg/Error.svg';
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
    error?: string;
    value: Item[];
}

export const MultiSelect: React.FC<MultiSelectProps> = ({
    data,
    placeholder,
    title,
    onChange,
    limit = 10,
    error = '',
    value
}) => {
    const calculateDropdownHeight = (itemsCount: number) => {
        let height = 45 + Math.min(itemsCount, 5) * 53;
        return height;
    };

    const [dropdownVisible, setDropdownVisible] = useState(false);
    const [searchTerm, setSearchTerm] = useState('');
    const dropdownHeight = calculateDropdownHeight(data.length);

    const filteredData = searchTerm
        ? data.filter((item) =>
              item.value.toLowerCase().includes(searchTerm.toLowerCase())
          )
        : data;

    const renderItem = ({ item, index }: { item: Item; index: number }) => {
        const isSelected = value.some(
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

    const handleComponentPress = () => {
        setDropdownVisible(!dropdownVisible);
    };

    const toggleItem = (item: Item) => {
        const index = value.findIndex(
            (selectedItem) => selectedItem.key === item.key
        );
        if (index !== -1) {
            const newSelectedItems = value.filter(
                (selectedItem) => selectedItem.key !== item.key
            );
            onChange(newSelectedItems);
        } else {
            if (value.length < limit) {
                const newSelectedItems = [...value, item];
                onChange(newSelectedItems);
            }
        }
    };

    return (
        <Box style={[styles.container]}>
            <Text paddingBottom="xs">{title}</Text>
            <Box
                style={[
                    styles.box,
                    { borderColor: error ? Colors.red : Colors.gray }
                ]}
            >
                <TouchableOpacity
                    onPress={handleComponentPress}
                    style={styles.input}
                >
                    <View
                        style={{
                            flexDirection: 'row',
                            alignItems: 'center',
                            flex: 1
                        }}
                    >
                        {value.length === 0 && (
                            <Text style={styles.placeholder}>
                                {placeholder}
                            </Text>
                        )}
                        <View
                            style={{ flexDirection: 'row', flexWrap: 'wrap' }}
                        >
                            {value.map((item) => (
                                <TouchableOpacity
                                    key={item.key}
                                    style={styles.tag}
                                    onPress={() => toggleItem(item)}
                                >
                                    <Text style={styles.tagText}>
                                        {item.value}
                                    </Text>
                                    <FontAwesomeIcon
                                        icon={faTimes}
                                        color={Colors.black}
                                    />
                                </TouchableOpacity>
                            ))}
                        </View>
                    </View>
                    <FontAwesomeIcon
                        icon={dropdownVisible ? faAngleUp : faAngleDown}
                        color={Colors.gray}
                    />
                </TouchableOpacity>
            </Box>
            {dropdownVisible && (
                <Box style={[styles.dropdown, { height: dropdownHeight }]}>
                    <TextInput
                        placeholder="Search..."
                        style={styles.searchInput}
                        value={searchTerm}
                        onChangeText={(text) => setSearchTerm(text)}
                    />
                    <FlatList
                        data={filteredData}
                        renderItem={renderItem}
                        keyExtractor={(item) => item.key}
                    />
                </Box>
            )}
            <Box
                gap="xxs"
                paddingTop="xs"
                flexDirection="row"
                alignItems="center"
            >
                {!dropdownVisible && (
                    <Box gap="xxs" flexDirection="row" alignItems="center">
                        {error && <ErrorIcon />}
                        <Text
                            color={error ? 'darkRed' : 'black'}
                            variant="caption-1"
                        >
                            {error}
                        </Text>
                    </Box>
                )}
            </Box>
        </Box>
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
    dropdown: {
        backgroundColor: 'white',
        borderColor: Colors.gray,
        borderWidth: 1,
        borderRadius: Border.base,
        marginTop: Spacing.xs
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
        justifyContent: 'space-between'
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
    divider: {
        height: 1,
        backgroundColor: Colors.gray
    }
});

export default MultiSelect;
