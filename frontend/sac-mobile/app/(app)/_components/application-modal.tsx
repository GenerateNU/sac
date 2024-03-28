import { View, Text, TouchableOpacity, Modal, StyleSheet } from "react-native";
import { Button } from "@/components/button";
export interface AskAQuestionModalProps {
    onClose: () => void;
}

const AskAQuestionModal = ({ onClose }: AskAQuestionModalProps) => {
    return (
        <Modal
            transparent={true}
            visible={true}
            onRequestClose={onClose}
        >
            <View className='flex-1 justify-center items-center bg-black/50'>
                <View className='bg-white opacity-100 rounded-lg pt-4 pb-5 w-80'>
                    <View className='items-center'>
                        <TouchableOpacity className='justify-center items-center absolute left-2 bg-black/30 rounded-full w-[10%] aspect-square' onPress={onClose}>
                            <Text>X</Text>
                        </TouchableOpacity>
                        <Text className=''>Send through</Text>
                        <View className='border-b w-[100%] p-2' />
                    </View>
                    <View className='p-4'>
                        <TouchableOpacity>
                            <Text>Open in Browser</Text>
                        </TouchableOpacity>
                        <TouchableOpacity>
                            <Text>Send To Email</Text>
                        </TouchableOpacity>
                        <TouchableOpacity>
                            <Text>Open in Browser and Send to Email</Text>
                        </TouchableOpacity>
                    </View>
                    <View className='border-b w-[100%] p-2' />
                    <View className='p-4'>
                        <Button onPress={() => onClose()}> Send</Button>
                    </View>
                </View>
            </View>
        </Modal>
    );
}

export { AskAQuestionModal };
