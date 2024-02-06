import FontAwesome from '@expo/vector-icons/FontAwesome';
import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { useFonts } from 'expo-font';
import { Stack, router } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { useEffect } from 'react';
export {
  // Catch any errors thrown by the Layout component.
  ErrorBoundary,
} from 'expo-router';
import { useAuthStore } from '@/hooks/use-auth';
import { getItemAsync } from 'expo-secure-store';
import { User } from '@/types/user';


export const unstable_settings = {
  // Ensure that reloading on `/modal` keeps a back button present.
  initialRouteName: '(app)',
};

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  const [loaded, error] = useFonts({
    SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
    ...FontAwesome.font,
  });

  // Expo Router uses Error Boundaries to catch errors in the navigation tree.
  useEffect(() => {
    if (error) throw error;
  }, [error]);

  useEffect(() => {
    if (loaded) {
      SplashScreen.hideAsync();
    }
  }, [loaded]);

  if (!loaded) {
    return null;
  }

  return <RootLayoutNav />;
}

function RootLayoutNav() {
  const { isLoggedIn } = useAuthStore();

  useEffect(() => {
    const checkLoginStatus = async () => {
      const accessToken = await getItemAsync('accessToken');
      const refreshToken = await getItemAsync('refreshToken');
      const savedUser = await getItemAsync('user');
      const user: User = savedUser ? JSON.parse(savedUser) : null;

      if (accessToken && refreshToken) {
        // Set the logged-in state and other user info
        useAuthStore.getState().login({ accessToken, refreshToken }, user);
      }
    };

    checkLoginStatus();
  }, []);

  useEffect(() => {
    if (!isLoggedIn) {
      router.push("/(auth)/login")
    }

  }, [isLoggedIn]);

  return (
    <Stack>
      <Stack.Screen name="(app)" options={{ headerShown: false }} />
      <Stack.Screen name="(auth)" options={{ headerShown: false }} />
    </Stack>
  );
}
