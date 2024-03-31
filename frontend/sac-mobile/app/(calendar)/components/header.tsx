import { useContext } from "react";
import { DateContext } from "../context";
import { TouchableOpacity, Text, Image, ViewProps, View } from "react-native";

export interface CalendarViewHeaderProps extends ViewProps {
    // FIXME: A) should this be dynamic
    // B) are these names correct
    onMenuButtonPressed?: () => {}
    onNotificationButtonPressed?: () => {}
    onBoxButtonPressed?: () => {}
}
  
const CalendarViewHeader = (p: CalendarViewHeaderProps) => {
    const date = useContext(DateContext);
    // FIXME: how to do images
    return (
    <View {...p}
        className="w-full h-12 flex flex-row justify-between">
        <View className="flex flex-row gap-3 justify-start items-center">
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
            <Text className='font-bold text-3xl'>{date.toLocaleString('default', { month: 'long' })}</Text>
        </View>
        <View className="flex flex-row gap-4 justify-end items-center">
            <TouchableOpacity
            onPress={() => {
                if (p.onNotificationButtonPressed !== undefined) {
                p.onNotificationButtonPressed()
                }
            }}>
            <Image style={{ resizeMode: 'contain' }} className="w-6 h-6" source={require("../../../assets/images/bell.png")}/>
            </TouchableOpacity>
            <TouchableOpacity
            onPress={() => {
                if (p.onMenuButtonPressed !== undefined) {
                p.onMenuButtonPressed()
                }}}>
            <Image style={{ resizeMode: 'contain' }} className="w-6 h-6" source={require("../../../assets/images/box.png")}/>
            </TouchableOpacity>
        </View>
    </View>)
}

export default CalendarViewHeader;