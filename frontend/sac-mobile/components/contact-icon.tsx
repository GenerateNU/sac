import { Text, View } from 'react-native';

import Email from '@/assets/images/svg/email.svg';
import Instagram from '@/assets/images/svg/instagram.svg';
import Slack from '@/assets/images/svg/slack.svg';
import { ContactType } from '@/types/contact';

const contactIcons = {
    email: Email,
    slack: Slack,
    instagram: Instagram
};

type ContactIconProps = {
    type: ContactType;
    width?: number;
    height?: number;
};

const ContactIcon = ({ type, width = 20, height = 20 }: ContactIconProps) => {
    const Icon = contactIcons[type];

    if (!Icon) {
        return <Text>Unknown Contact Type</Text>;
    }

    return (
        <View className="pr-[2%]">
            {Icon && <Icon width={width} height={height} />}
        </View>
    );
};

export { ContactIcon };
