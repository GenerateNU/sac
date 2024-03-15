import React from 'react';
import { ScrollView, Text, View } from 'react-native';
import { Button } from '@/components/button';

import { useAuthStore } from '@/hooks/use-auth';
import { SafeAreaView } from 'react-native-safe-area-context';

const Home = () => {
    const { logout } = useAuthStore();
    return (
        <SafeAreaView className="bg-neutral-500 h-[100%]" edges={['top']}>
            <View className="flex-1">
                <View className="px-[8%] pb-[10%]">
                    <View className="pt-[1%]">
                    </View>
                    <View className="pt-[20%] pb-[6%]">
                    </View>
                </View>

                <ScrollView className="bg-white pt-[13%] pb-[2%] flex-1 rounded-tl-3xl rounded-tr-3xl px-[8%]">
                    <View className="pb-[8%] flex-row justify-between">
                        <Text className="text-black font-bold text-4xl">Club Name</Text>
                    </View>
                    <View className="pb-[8%]">
                        <Text className="text-black font-bold">Club Description</Text>
                        <Text className="text-black">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</Text>
                    </View>
                    <ScrollView horizontal={true} className = "pb-[10%]">
                        <View className="flex-row">
                            {/* To be replaced by event components */}
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">Event 1</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">Event 2</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">Event 3</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">Event 4</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">Event 5</Text>
                            </Button>
                        </View>
                    </ScrollView>

                    <ScrollView horizontal={true} className = "pb-[10%]">
                        <View className="flex-row">
                            {/* To be replaced by e-board components */}
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">E-board 1</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">E-board 2</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">E-board 3</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">E-board 4</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">E-board 5</Text>
                            </Button>
                        </View>
                    </ScrollView>

                    <ScrollView horizontal={true}>
                        <View className="flex-row">
                            {/* To be replaced by FAQ components */}
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">FAQ 1</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">FAQ 2</Text>
                            </Button>
                            <Button variant="outline" size="sm" className="mr-2">
                                <Text className="text-black">FAQ 3</Text>
                            </Button>
                        </View>
                    </ScrollView>

                </ScrollView>
            </View>
        </SafeAreaView>
    );
};

export default Home;
