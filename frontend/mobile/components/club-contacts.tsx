import { Text, View } from 'react-native';

import { useClubContactsByClubIDQuery } from '@sac/lib/api/club';
import { Contact } from '@sac/lib/api/contact';
import { uuid } from '@sac/lib/types/uuid';

import { ContactIcon } from './contact-icon';

const ClubContacts = ({ clubID }: { clubID: uuid }) => {
    const {
        data: clubContacts,
        error,
        isLoading
    } = useClubContactsByClubIDQuery(clubID);

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        console.error(error);
        return <Text>Error</Text>;
    }

    if (!clubContacts || clubContacts.length === 0) {
        return <Text>No contacts found</Text>;
    }

    return (
        <View className="flex flex-row items-center justify-start pt-2 space-x-3">
            {clubContacts.map((contact: Contact) => (
                <View key={contact.id}>
                    <ContactIcon type={contact.type} key={contact.id} />
                </View>
            ))}
        </View>
    );
};

export { ClubContacts };
