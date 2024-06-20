import React from 'react';
import { StyleSheet, TouchableOpacity } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

import { router } from 'expo-router';

import {
    IconDefinition,
    faCalendar,
    faSignOutAlt,
    faUser
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';

import { GlobalLayout } from '@/src/app/design-system/components/GlobalLayout/GlobalLayout';
import { resetClub } from '@/src/store/slices/clubSlice';
import { resetEvent } from '@/src/store/slices/eventSlice';
import { resetAccessToken } from '@/src/store/slices/globalSlice';
import { resetUser } from '@/src/store/slices/userSlice';
import { useAppDispatch, useAppSelector } from '@/src/store/store';

import { Box, Colors, SACColors, Spacing, Text } from '../../design-system';

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
    const dispatch = useAppDispatch();
    const user = useAppSelector((state) => state.user);

    return (
        <GlobalLayout>
            <SafeAreaView style={styles.container}>
                <Box
                    paddingTop="xxl"
                    flexDirection="row"
                    gap="m"
                    alignItems="center"
                >
                    <Box flexDirection="column" gap="xxs">
                        <Text variant="subheader-2">{user.name}</Text>
                        <Text>{user.email}</Text>
                    </Box>
                </Box>
                <Box width="100%">
                    {/* <ProfileItem
                        onPress={() => router.push('/app/user/detail/')}
                        icon={faUser}
                        text="Edit Profile"
                    />
                    <ProfileItem
                        icon={faHeart}
                        onPress={() => router.push('/app/user/interest/')}
                        text="Edit Interests"
                    /> */}
                    <ProfileItem
                        icon={faUser}
                        onPress={() => router.push('/app/user/following/')}
                        text="Following"
                    />
                    <ProfileItem
                        icon={faCalendar}
                        onPress={() => router.push('/app/user/events/')}
                        text="Upcoming Events"
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
                        onPress={() => {
                            dispatch(resetAccessToken());
                            dispatch(resetEvent());
                            dispatch(resetUser());
                            dispatch(resetClub());
                        }}
                    />
                </Box>
            </SafeAreaView>
        </GlobalLayout>
    );
};

const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        alignItems: 'flex-start',
        justifyContent: 'flex-start',
        marginVertical: Spacing.m,
        gap: Spacing.l
    },
    title: {
        fontSize: 20,
        fontWeight: 'bold',
        marginBottom: 20
    }
});

export default ProfilePage;
