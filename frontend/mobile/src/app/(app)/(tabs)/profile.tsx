import React from 'react';
import { StyleSheet, TouchableOpacity } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import {
    IconDefinition,
    faHeart,
    faSignOutAlt,
    faUser
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';
import { Avatar } from '@rneui/base';

import { GlobalLayout } from '@/src/app/(design-system)/components/GlobalLayout/GlobalLayout';

import { Box, Colors, SACColors, Spacing, Text } from '../../(design-system)';

type ProfileItemProps = {
    icon: IconDefinition;
    text: string;
    textColor?: SACColors;
    onPress?: () => void;
};

const ProfileItem = ({ icon, text, textColor, onPress }: ProfileItemProps) => (
    <TouchableOpacity onPress={onPress}>
        <Box
            gap="l"
            paddingVertical="s"
            flexDirection="row"
            alignItems="center"
        >
            <FontAwesomeIcon
                size={24}
                color={textColor ? Colors.darkRed : Colors.gray}
                icon={icon}
            />
            <Text color={textColor}>{text}</Text>
        </Box>
    </TouchableOpacity>
);

const ProfilePage = () => {
    return (
        <GlobalLayout>
            <SafeAreaView style={styles.container}>
                <Box
                    paddingTop="xxl"
                    flexDirection="row"
                    gap="m"
                    alignItems="center"
                >
                    <Avatar
                        source={{
                            uri: 'https://media1.popsugar-assets.com/files/thumbor/S6_lryTon-0orhMkLrw6m1yIFww=/1500x1500/filters:format_auto():quality(85):extract_cover()/2017/02/28/003/n/1922441/ace7eba458b602264e7940.39836479_edit_img_image_43244124_1488304081.jpg'
                        }}
                        size={60}
                        rounded
                    />
                    <Box flexDirection="column" gap="xxs">
                        <Text variant="subheader-2">Quokka</Text>
                        <Text>quokka@northeastern.edu</Text>
                    </Box>
                </Box>
                <Box width="100%">
                    <ProfileItem
                        onPress={() => router.push('/user/detail/')}
                        icon={faUser}
                        text="Edit Profile"
                    />
                    <ProfileItem
                        icon={faHeart}
                        onPress={() => router.push('/user/interest/')}
                        text="Edit Interests"
                    />
                    <Box
                        width="100%"
                        height={1}
                        backgroundColor="gray"
                        marginVertical="s"
                    />
                    <ProfileItem
                        icon={faSignOutAlt}
                        text="Logout"
                        textColor="darkRed"
                    />
                </Box>
            </SafeAreaView>
        </GlobalLayout>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: 'column',
        alignItems: 'flex-start',
        justifyContent: 'flex-start',
        marginVertical: Spacing.m,
        gap: Spacing.l,
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold',
        marginBottom: 20
    }
});

export default ProfilePage;
