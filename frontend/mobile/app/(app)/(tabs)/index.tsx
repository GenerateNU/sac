import { StyleSheet, View } from 'react-native';

import { Text } from '@/app/(design-system)';
import { RecruitmentInfo } from '@/components/club-recruitment-info';

const HomePage = () => {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>Hello</Text>
            <RecruitmentInfo
                recruitmentText="Fill out our application"
                recruitingType="unrestricted"
                isRecruiting={true}
                recruitmentCycle="always"
                color="aqua"
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
