import { Box, Button, Spacing, Text } from "@/src/app/(design-system)";
import { SafeAreaView } from "react-native";
import { faArrowsRotate } from "@fortawesome/free-solid-svg-icons";

type EventPageErrorProps = {
    refetch: React.Dispatch<React.SetStateAction<boolean>>
};

const EventPageError = ({ refetch }: EventPageErrorProps) => {
    return (
        <SafeAreaView style={{ margin: Spacing.m, gap: 20 }}>
            <Box width='75%' marginTop='m' gap='m'>
                <Text variant='header-1'>
                    Oops, something went wrong!
                </Text>
            </Box>
            <Box>
                <Button 
                    variant="iconButton"
                    icon={faArrowsRotate}
                    color="black"
                    iconPosition="left"
                    size="full"
                    onPress={() => refetch((prev) => !prev)}
                >
                    Refresh
                </Button>
            </Box>
        </SafeAreaView>
    );
};

export default EventPageError;