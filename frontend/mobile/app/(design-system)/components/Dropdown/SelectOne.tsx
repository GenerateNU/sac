import React, { useState } from 'react';
import { FlatList, StyleSheet, TouchableOpacity } from 'react-native';

import { faAngleDown, faAngleUp } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import ErrorIcon from '@/app/(design-system)/assets/svg/Error.svg';
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
    value: Item;
    onChange: (value: Item) => void;
    error?: string;
}

export const SelectOne: React.FC<SelectOneProps> = ({
    data,
    placeholder,
    title,
    value,
    onChange,
    error = ''
}) => {
    const [dropdownVisible, setDropdownVisible] = useState(false);

    const handleSelect = (item: Item) => {
        onChange(item);
        setDropdownVisible(false);
    };

    const renderItem = ({ item, index }: { item: Item; index: number }) => {
        const isLastItem = index === data.length - 1;
        return (
            <>
                <TouchableOpacity
                    style={styles.item}
                    onPress={() => handleSelect(item)}
                >
                    <Text style={{ flex: 1 }}>{item.value}</Text>
                </TouchableOpacity>
                {!isLastItem && <Box style={styles.divider} />}
            </>
        );
    };

    let dropdownHeight = data.length * 53;
    if (data.length >= 5) {
        dropdownHeight = 265;
    }

    return (
        <Box width="100%" borderRadius="medium">
            {title && <Text paddingBottom="xs">{title}</Text>}
            <TouchableOpacity
                style={[
                    styles.box,
                    { width: '100%' },
                    { borderColor: error ? Colors.red : Colors.gray }
                ]}
                onPress={() => setDropdownVisible(!dropdownVisible)}
            >
                <Box style={styles.contentContainer}>
                    <Text variant="body-1" color={!value ? 'gray' : 'black'}>
                        {value ? value.value : placeholder}
                    </Text>
                    <FontAwesomeIcon
                        color={Colors.gray}
                        icon={dropdownVisible ? faAngleUp : faAngleDown}
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
        paddingHorizontal: Spacing.xxs,
        width: '100%'
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
        padding: Spacing.m,
        width: '100%'
    },
    divider: {
        height: 1,
        backgroundColor: Colors.gray
    }
});
