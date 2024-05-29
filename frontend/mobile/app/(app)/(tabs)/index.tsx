import { StyleSheet, Text, View } from 'react-native';

import { faChevronUp } from '@fortawesome/free-solid-svg-icons';

import { Button } from '@/app/(design-system)/Button/Button';

const HomePage = () => {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>Home Page</Text>
            <Button
                variant="iconButton"
                icon={faChevronUp}
                iconColor={'black'}
                size={'small'}
            />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center'
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold'
    },
    separator: {
        marginVertical: 30,
        height: 1,
        width: '80%'
    }
});

export default HomePage;
