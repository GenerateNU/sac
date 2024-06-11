import React from 'react';
import { FlatList, Pressable } from 'react-native';

import { router } from 'expo-router';

import { PointOfContact } from '@generatesac/lib';

import { SACColors } from '../../shared/colors';
import { PointOfContact as PointOfContactComponent } from './PointOfContactCard';

interface PointOfContactListProps {
    contacts: PointOfContact[];
    color: SACColors;
}

export const PointOfContactList: React.FC<PointOfContactListProps> = ({
    contacts,
    color
}) => {
    const renderPointOfContact = ({ item }: { item: PointOfContact }) => {
        return (
            <Pressable onPress={() => router.push(`/contact/${item.id}`)}>
                <PointOfContactComponent
                    imageUrl={item.photo_file.file_url}
                    name={item.name}
                    position={item.position}
                    color={color}
                />
            </Pressable>
        );
    };

    return (
        <FlatList
            showsHorizontalScrollIndicator={false}
            horizontal
            contentContainerStyle={{
                gap: 15,
                paddingBottom: 40
            }}
            data={contacts}
            renderItem={renderPointOfContact}
            keyExtractor={(item) => item.id}
        />
    );
};
