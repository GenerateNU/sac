import { Text, View } from '@/components/Themed';
import { Link, Stack } from 'expo-router';

const NotFoundScreen = () => {
  return (
    <>
      <Stack.Screen options={{ title: 'Oops!' }} />
      <View className="justify-center flex-1 p-4 align-center">
        <Text className="text-xl font-bold">This screen doesn't exist.</Text>

        <Link href="/" className="py-3 mt-3">
          <Text className="text-lg text-[#2e78b7]">Go to home screen!</Text>
        </Link>
      </View>
    </>
  );
};

export default NotFoundScreen;
