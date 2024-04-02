import { TouchableOpacity } from 'react-native';

import {
    AntDesign,
    Feather,
    MaterialCommunityIcons,
    MaterialIcons
} from '@expo/vector-icons';

import { Contact } from '@/types/contact';

const contactIconMapper = {
    facebook: <AntDesign name="facebook-square" size={24} color="black" />,
    instagram: <AntDesign name="instagram" size={24} color="black" />,
    twitter: <AntDesign name="twitter" size={24} color="black" />,
    linkedin: <AntDesign name="linkedin-square" size={24} color="black" />,
    youtube: <AntDesign name="youtube" size={24} color="black" />,
    github: <AntDesign name="github" size={24} color="black" />,
    slack: <MaterialCommunityIcons name="slack" size={24} color="black" />,
    discord: <MaterialIcons name="discord" size={24} color="black" />,
    email: <MaterialIcons name="email" size={24} color="black" />,
    customSite: <Feather name="link" size={24} color="black" />
} as const;

const ContactIcon = ({ contact }: { contact: Contact }) => {
    return (
        <TouchableOpacity
            className="flex items-center justify-center"
            onPress={() => console.log('Navigating to:', contact.content)}
        >
            {contactIconMapper[contact.type]}
        </TouchableOpacity>
    );
};

export default ContactIcon;
