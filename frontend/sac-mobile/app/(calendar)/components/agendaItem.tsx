import React, { useCallback } from 'react';
import {
    Alert,
    Button,
    StyleSheet,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

interface ItemProps {
    item: any;
}

const AgendaItem = (props: ItemProps) => {
    const { item } = props;

    const buttonPressed = useCallback(() => {
        Alert.alert('Show me more');
    }, []);

    const itemPressed = useCallback(() => {
        Alert.alert(item.title);
    }, []);

    if (!item) {
        return (
            <View style={styles.emptyItem}>
                <Text style={styles.emptyItemText}>
                    No Events Planned Today
                </Text>
            </View>
        );
    }

    return (
        <TouchableOpacity onPress={itemPressed} style={styles.item}>
            <View>
                <Text style={styles.itemHourText}>{item.hour}</Text>
                <Text style={styles.itemDurationText}>{item.duration}</Text>
            </View>
            <View>
              <Text style={styles.itemTitleText}>{item.title}</Text>
              <Text style={styles.itemText}>{item.location}</Text>
            </View>
            <View style={styles.itemButtonContainer}>
                <Button color={'grey'} title={'Info'} onPress={buttonPressed} />
            </View>
        </TouchableOpacity>
    );
};

export default React.memo(AgendaItem);

const styles = StyleSheet.create({
    item: {
        marginTop: 4,
        marginBottom: 4,
        marginLeft: 16,
        marginRight: 8,
        padding: 20,
        backgroundColor: 'white',
        shadowColor: 'black',
        shadowOpacity: 0.25,
        shadowOffset: { width: 0, height: 2 },
        flexDirection: 'row',
        borderRadius: 12
    },
    itemHourText: {
        color: 'black'
    },
    itemText: {
      color: 'black',
      paddingLeft: 16
  },
    itemDurationText: {
        color: 'grey',
        fontSize: 12,
        marginTop: 4,
        marginLeft: 'auto'
    },
    itemTitleText: {
        color: 'black',
        marginLeft: 16,
        fontWeight: 'bold',
        fontSize: 16
    },
    itemButtonContainer: {
        flex: 1,
        alignItems: 'flex-end'
    },
    emptyItem: {
        paddingLeft: 20,
        height: 52,
        justifyContent: 'center',
        borderBottomWidth: 1,
        borderBottomColor: 'lightgrey'
    },
    emptyItemText: {
        color: 'lightgrey',
        fontSize: 14
    }
});
