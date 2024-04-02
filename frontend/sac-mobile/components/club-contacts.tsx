import { Text, View } from 'react-native';

import { useClubContacts } from '@/hooks/use-club-contacts';
import { Contact } from '@/types/contact';
import { uuid } from '@/types/uuid';

import ContactIcon from './contact-icon';

const ClubContacts = ({ clubID }: { clubID: uuid }) => {
    const { data: clubContacts, error, isLoading } = useClubContacts(clubID);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    if (!clubContacts || clubContacts.length === 0) {
        return <Text>No contacts found</Text>;
    }

    return (
        <View className="flex flex-row items-center justify-start space-x-3">
            {clubContacts.map((contact: Contact) => (
                <View key={contact.id}>
                    <ContactIcon contact={contact} key={contact.id} />
                </View>
            ))}
        </View>
    );
};

export default ClubContacts;
