import { useAppDispatch, useAppSelector } from "@/src/store/store";
import { userApi } from "@generatesac/lib";
import { useEffect, useState } from "react";
import PageError from "../../design-system/components/PageError/PageError";
import { Arrow, Box, Button, Text } from "../../design-system";
import ClubCardStandardSkeleton from "../../design-system/components/ClubCard/Skeletons/ClubCardStandardSkeleton";
import { Clubcard } from "../../design-system/components/ClubCard/ClubCard";
import { SafeAreaView } from "react-native";
import { GlobalLayout } from "../../design-system/components/GlobalLayout/GlobalLayout";
import { Stack, router } from "expo-router";

const Following = () => {
    const { id, following } = useAppSelector((state) => state.user)
    const [getFollowers, { data, isLoading, error }] = userApi.useLazyGetUserFollowingQuery();

    useEffect(() => {
        getFollowers(id)
    }, [following])

    if (error) {
        return (
            <PageError refetch={() => {}}/>
        )
    }
    
    return (
        <>
            <Stack.Screen
                options={{
                    headerTitle: () => (
                        <Text variant="subheader-2">Clubs you Follow</Text>
                    ),
                    headerTransparent: true,
                    headerShown: true,
                    headerLeft: () => <Arrow color="black" />,
                }}
            />
            <SafeAreaView>
                <GlobalLayout>
                    <Box alignItems="center" flexDirection="column" flex={0} marginTop='xl'>
                        {isLoading || !data ?
                            <Box gap='s' width='100%'>
                                <ClubCardStandardSkeleton />
                                <ClubCardStandardSkeleton />
                                <ClubCardStandardSkeleton />
                            </Box>
                            : 
                            <Box gap='s' width='100%'>
                                {data.map((club) => {
                                    if (club.name !== 'SAC') {
                                        return (
                                            <Clubcard
                                                key={club.id}
                                                id={club.id}
                                                name={club.name}
                                                logo={club.logo as string}
                                            />
                                        )
                                    }
                                })}
                                <Box marginTop='m'>
                                    <Button 
                                        variant="standardButton"
                                        color="darkRed"
                                        onPress={() => router.push('/app/calendar')}
                                    >
                                        Explore New Clubs
                                    </Button>
                                </Box>
                            </Box>
                        }
                    </Box>
                </GlobalLayout>
            </SafeAreaView>
        </>
    );
};

export default Following;