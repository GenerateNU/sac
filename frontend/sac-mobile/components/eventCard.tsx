import { Text, TouchableOpacity, TouchableOpacityProps, View } from 'react-native';

import { cn } from '@/lib/utils';
import { Button } from './button';


export interface EventCardProps
    {
        title: string
        time: string
        date: string
        location: string
     }

const EventCard = ({title, time, date, location }: EventCardProps) => {
    return (
        <View
            className='flex text-center bg-card-bg text-white justify-end items-start mr-3 rounded-lg p-4 w-80 h-40'
        >
            <View className='flex-row px-3 items-end space-x-10'>
                <View>
                    <Text className={cn(`text-lg font-bold mb-2`, { fontFamily: 'OpenSans-SemiBold' })}>{title}</Text>
                    <Text className='text-sm text-gray-600'>{time}</Text>
                    <Text className='text-sm text-gray-600'>{date}</Text>
                    <Text className='text-sm text-gray-600'>{location}</Text>
                </View>
                <Button className='bg-[#747474] rounded-lg'>Register</Button>
                
            </View>

        </View>
    );
};

EventCard.displayName = 'Event Card';

export { EventCard };
