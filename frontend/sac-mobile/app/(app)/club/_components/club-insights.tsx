import { pluralize } from "@/lib/utils";
import React from "react";
import { View, Text } from "react-native";
import CircleClock from '@/assets/images/svg/circle-clock.svg';
import MemberCount from '@/assets/images/svg/members.svg';
import PeaceHand from '@/assets/images/svg/peace-hand.svg';

const HourlyCommitment = ({ numHours }: { numHours: number }) => {
    return (
        <View className="flex-col items-center justify-center gap-y-2" style={{ width: 118 }}>
            <CircleClock width={24} height={24} />
            <Text className="font-semibold text-center text-black">{numHours} {pluralize(numHours, 'Hour', 'Hours')}</Text>
            <Text className="text-xs font-normal text-center text-neutral-400">Weekly Time Commitment</Text>
        </View>
    );
}

const Membership = ({ numMembers }: { numMembers: number }) => {
    const membersRange = numMembers > 100 ? '100+' : numMembers > 50 ? '50 - 100' : numMembers > 25 ? '25 - 50' : '1 - 25';

    return (
        <View className="flex-col items-center justify-center gap-y-2" style={{ width: 118 }}>
            <MemberCount width={24} height={24} />
            <Text className="font-semibold text-center text-black">{membersRange}</Text>
            <View>
                <Text className="text-xs font-normal text-center text-neutral-400 ">Member</Text>
                <Text className="text-xs font-normal text-center text-neutral-400 ">Count</Text>
            </View>
        </View>
    );
}

const OneWord = ({ word }: { word: string }) => {
    return (
        <View className="flex-col items-center justify-center gap-y-2" style={{ width: 118 }}>
            <PeaceHand width={24} height={24} />
            <Text className="font-semibold text-center text-black">{word}</Text>
            <Text className="text-xs font-normal text-center text-neutral-400 ">One Word To Describe Us</Text>
        </View>
    );
}

const ClubInsights = () => {
    return (
        <View className="flex-row justify-between pb-5">
            <View
                className="col-span-1 pt-3 pb-3 bg-white rounded-lg shadow-sm"
                style={{ width: '32%' }}
            >
                <HourlyCommitment numHours={5} />
            </View>
            <View
                className="col-span-1 pt-3 pb-3 bg-white rounded-lg shadow-sm"
                style={{ width: '32%' }}
            >
                <Membership numMembers={100} />
            </View>
            <View
                className="col-span-1 pt-3 pb-3 bg-white rounded-lg shadow-sm"
                style={{ width: '32%' }}
            >
                <OneWord word={'Builders'} />
            </View>
        </View>
    )
}

export { ClubInsights }