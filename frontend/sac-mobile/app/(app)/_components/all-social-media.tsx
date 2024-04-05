import { Button } from '@/components/button';
import { useAdmin } from '@/hooks/use-admin';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import {Text, View } from 'react-native';

const SocialMediaIcons = () => {

    const SlackIcon = ({ color }: { color: string }) => (
        <MaterialCommunityIcons name="slack" size={24} color={color} />
    );
    
    const EmailIcon = ({ color }: { color: string }) => (
        <MaterialCommunityIcons name="email" size={24} color={color} />
    );
    
    const InstagramIcon = ({ color }: { color: string }) => (
        <MaterialCommunityIcons name="instagram" size={24} color={color} />
    );
    
    const { data: eboardmembers, isLoading, error } = useAdmin("sac");

    // TODO: Implement social media click handling
    function handleSocialMediaClick(arg0: string): void {
        throw new Error('Function not implemented.');
    }

    if (isLoading) {
        return <Text>Loading...</Text>;
    }

    if (error) {
        return <Text>Error: {error.message}</Text>;
    }

    return (
        <View className='flex-row pb-[3%]'>
            <Button variant="secondary" onPress={() => handleSocialMediaClick('email')}>
                <EmailIcon color="black" />
            </Button>

            <Button variant="secondary" onPress={() => handleSocialMediaClick('instagram')}>
                <InstagramIcon color="black" />
            </Button>

            <Button variant="secondary" onPress={() => handleSocialMediaClick('slack')}>
                <SlackIcon color="black" />
            </Button>
        </View>
    );
}

export { SocialMediaIcons };