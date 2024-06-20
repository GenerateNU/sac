import React, { useState } from 'react';
import { TextInput, TouchableOpacity } from 'react-native';
import { StyleSheet } from 'react-native';

import { faArrowLeft, faSearch } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { Colors, SACColors } from '../../shared/colors';
import { createStyles } from '../../theme';
import { Box } from '../Box/Box';
import { Text } from '../Text/Text';
import { TextBox } from '../Textbox/Textbox';

interface SearchBarProps {
    placeholder?: string;
    onSearch: (query: string) => void;
    onCancel?: () => void;
    onBack?: () => void;
    initialQuery?: string;
}

const SearchBar: React.FC<SearchBarProps> = ({
    placeholder = 'Search',
    onSearch,
    onCancel,
    onBack,
    initialQuery = ''
}) => {
    const [query, setQuery] = useState(initialQuery);

    const handleSearch = () => {
        onSearch(query);
    };

    return (
        <Box
            flexDirection="row"
            alignItems="center"
            width={'100%'}
            paddingBottom="s"
        >
            {onBack && (
                <TouchableOpacity onPress={onBack}>
                    <FontAwesomeIcon icon={faArrowLeft} size={16} />
                </TouchableOpacity>
            )}
            <Box
                flexDirection="row"
                alignItems="center"
                flex={1}
                borderRadius="sm"
                paddingHorizontal="s"
                borderColor="gray"
                borderWidth={1}
                marginHorizontal="s"
            >
                <FontAwesomeIcon icon={faSearch} size={16} />
                <TextInput
                    style={styles.input}
                    placeholder={placeholder}
                    value={query}
                    onChangeText={setQuery}
                    onSubmitEditing={handleSearch}
                />
            </Box>
            {query.length > 0 && onCancel && (
                <TouchableOpacity
                    onPress={() => {
                        setQuery('');
                        onCancel && onCancel();
                    }}
                    style={styles.cancelButton}
                >
                    <Text color="gray">Cancel</Text>
                </TouchableOpacity>
            )}
        </Box>
    );
};

const styles = StyleSheet.create({
    input: {
        flex: 1,
        height: 40,
        fontSize: 16,
        marginLeft: 8
    },
    cancelButton: {
        marginLeft: 8
    }
});

export default SearchBar;
