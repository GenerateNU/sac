import { Text, TouchableOpacity, View } from 'react-native';

import { router } from 'expo-router';

import { Club } from '@/types/club';

import { Button } from './button';
import { pluralize } from '@/lib/utils';

const ClubCard = ({ club }: { club: Club }) => {

    return (
        <TouchableOpacity
            style={{ width: 'auto' }}
            onPress={() => router.push(`/club/${club.id}`)}
        >
            <View style={{ flexDirection: 'row', justifyContent: 'space-between', padding: 16, backgroundColor: 'white', borderWidth: 1, borderColor: '#CBD5E0', borderRadius: 8 }}>
                <View style={{ flexDirection: 'row', alignItems: 'center', gap: 8 }}>
                    <View style={{ height: 40, width: 40, borderWidth: 2, borderColor: '#718096', backgroundColor: '#F7FAFC', borderRadius: 20, alignItems: 'flex-start' }} />
                    <View>
                        <Text style={{ fontSize: 16, fontWeight: 'bold', lineHeight: 20, color: '#1A202C', width: 200 }}>
                            {club.name}
                        </Text>
                        <Text style={{ fontSize: 14, fontWeight: 'normal', lineHeight: 20, color: '#1A202C' }}>
                            {club.num_members} {pluralize(club.num_members, 'member', 'members')}
                        </Text>
                    </View>
                </View>
                <Button style={{ alignItems: 'flex-start' }}>Follow</Button>
            </View>
        </TouchableOpacity>
    );
};

export { ClubCard };