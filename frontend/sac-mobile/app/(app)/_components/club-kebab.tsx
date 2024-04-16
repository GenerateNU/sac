import React from 'react';
import { Platform } from 'react-native';

import { NativeActionEvent } from '@react-native-menu/menu';

import { Kebab } from '@/components/kebab';

const ClubKebab = () => {
    const items = [
        {
            id: 'share',
            title: 'Share Club',
            image: Platform.select({
                ios: 'square.and.arrow.up',
                android: 'share-variant'
            })
        },
        {
            id: 'report',
            title: 'Report Club',
            image: Platform.select({
                ios: 'person.crop.circle.badge.exclamationmark.fill',
                android: 'person-circle-outline'
            })
        }
    ];

    const onPress = ({ nativeEvent }: NativeActionEvent) => {
        console.warn(JSON.stringify(nativeEvent));
    };

    return <Kebab items={items} onPressAction={onPress} />;
};

export { ClubKebab };
