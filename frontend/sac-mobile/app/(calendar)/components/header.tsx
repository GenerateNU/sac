import { useContext } from "react";
import { DateContext } from "../context";
import { TouchableOpacity, Text, Image, ViewProps, View } from "react-native";
import { format } from "date-fns";

export interface CalendarViewHeaderProps extends ViewProps {
    onMenuButtonPressed?: () => {}
}
  
const CalendarViewHeader = (p: CalendarViewHeaderProps) => {
    const dateCtx = useContext(DateContext);
    return (
    <View {...p}
        className="w-full bg-pink-300 h-12 flex flex-row justify-start items-center ">
            <TouchableOpacity
            onPress={() => {
                if (p.onMenuButtonPressed !== undefined) {
                p.onMenuButtonPressed()
                }
            }}>
                <View className='px-2 py-3 w-12 h-12 flex flex-col justify-between'>
                    <View className='bg-black w-full h-1'></View>
                    <View className='bg-black w-2/3 h-1'></View>
                    <View className='bg-black w-1/3 h-1'></View>
                </View>
            </TouchableOpacity>
            <Text className='font-bold text-3xl'>{format(dateCtx.topViewDate, "LLLL yyyy")}</Text>
        </View>);
}

export default CalendarViewHeader;