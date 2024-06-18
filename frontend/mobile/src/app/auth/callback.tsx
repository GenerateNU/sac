import { Linking } from "react-native";
import { Box, Text } from "../design-system"
import { useEffect } from "react";
import { useLocalSearchParams } from "expo-router";
import { useAppSelector } from "@/src/store/store";

const OAuthCallback = () => {
    const params = useLocalSearchParams();
    const { sacLoginSession } = useAppSelector((state) => state.user);

    useEffect(() => {
        const { code, session_state, state } = params;
        console.log(params)
        console.log(sacLoginSession);
        if (!code || !session_state || !state) {
            console.log('failed!')
        }

    }, [params])

    return (
        <Box>
            <Text>
                Callback!!!!!!!
            </Text>
        </Box>
    )
}

export default OAuthCallback;