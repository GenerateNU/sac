import EditScreenInfo from '@/components/EditScreenInfo';
import { Text, View } from '@/components/Themed';

const TabOneScreen = () => {
  return (
    <View className="items-center justify-center flex-1">
      <Text className="text-xl font-bold">Tab One</Text>
      <EditScreenInfo path="app/(home)/index.tsx" />
    </View>
  );
};

export default TabOneScreen;
