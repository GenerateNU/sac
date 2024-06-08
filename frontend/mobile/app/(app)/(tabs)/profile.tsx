import { StyleSheet, Text, View } from 'react-native';

import { ClubPage } from '@/app/(design-system)/components/ClubPage/ClubPage';
import { GlobalLayout } from '@/app/(design-system)/components/GlobalLayout/GlobalLayout';

import {
    mockClub,
    mockContacts,
    mockEvents,
    mockTags,
    pointOfContactMocks
} from '../(tabs)/mockClubPageData';

const ProfilePage = () => {
    return (
        <GlobalLayout>
            {/* <View style={styles.container}>
                <Text style={styles.title}>Profile</Text>
            </View> */}
            <ClubPage
                club={mockClub}
                tags={mockTags}
                events={mockEvents}
                contacts={mockContacts}
                pointOfContacts={pointOfContactMocks}
                color="darkBlue"
            />
        </GlobalLayout>
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

export default ProfilePage;
