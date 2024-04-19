import { TouchableOpacity, TouchableOpacityProps, View } from 'react-native';

export const HamburgerButton = (p: TouchableOpacityProps) => {
    return (
        <TouchableOpacity {...p}>
            <View className="px-2 py-3 w-full h-full flex flex-col justify-between">
                <View className="bg-black w-full h-1 rounded-full"></View>
                <View className="bg-black w-2/3 h-1 rounded-full"></View>
                <View className="bg-black w-1/3 h-1 rounded-full"></View>
            </View>
        </TouchableOpacity>
    );
};

export default HamburgerButton;
