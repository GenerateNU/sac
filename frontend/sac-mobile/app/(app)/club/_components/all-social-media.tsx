import { Button } from '@/components/button';
import { useClub, useClubAdmin } from '@/hooks/use-club';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { Text, View } from 'react-native';
import SlackIcon from '@/assets/images/svg/slack.svg';
import InstagramIcon from '@/assets/images/svg/instagram.svg';
import EmailIcon from '@/assets/images/svg/email.svg';

const SocialMediaIcons = () => {


    const { data: eboardmembers, isLoading, error } = useClubAdmin("sac");

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
            <Button variant="icon" onPress={() => handleSocialMediaClick('email')}>
                <EmailIcon color="black" />
            </Button>

            <Button variant="icon" onPress={() => handleSocialMediaClick('instagram')}>
                <InstagramIcon color="black" />
            </Button>

            <Button variant="icon" onPress={() => handleSocialMediaClick('slack')}>
                <SlackIcon color="black" />
            </Button>
        </View>
    );
}

export { SocialMediaIcons };