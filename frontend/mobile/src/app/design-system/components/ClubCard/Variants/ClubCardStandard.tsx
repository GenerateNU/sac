import { TouchableOpacity } from "react-native-gesture-handler";
import { createStyles } from "../../../theme";
import { Box } from "../../Box/Box";
import { ClubCardProps } from "../ClubCard";
import { Text } from "../../Text/Text";
import { ClubIcon } from "../../ClubIcon/ClubIcon";
import { router } from "expo-router";
import { useAppDispatch } from "@/src/store/store";
import { resetEvent } from "@/src/store/slices/eventSlice";

const ClubCardStandard = ({ logo, name, id }: ClubCardProps) => {
    const dispatch = useAppDispatch();

    return (
        <Box {...styles.cardContainer}>
            <TouchableOpacity
                onPress={() => { 
                    dispatch(resetEvent());
                    router.navigate(`/app/club/${id}`)
                }}
            >
                <Box { ...styles.cardSubContainer}>
                   <ClubIcon imageUrl={logo} size={55} />
                   <Text>
                        {name.length > 28 ? name.slice(0, 28).trim() + '...' : name}
                   </Text>
                </Box>
            </TouchableOpacity>
        </Box>
    )
}

const styles = createStyles({
    cardContainer: {
        shadowColor: 'black',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.1,
        shadowRadius: 2,
        backgroundColor: 'white',
        borderRadius: 'md',
        width: '100%',
    },
    cardSubContainer: {
        width: '100%',
        borderRadius: 'md',
        padding: 'm',
        backgroundColor: 'white',
        flexDirection: 'row',
        alignItems: 'center',
        gap: 's',
    }
});

export default ClubCardStandard;