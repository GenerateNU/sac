import { Text, View, SafeAreaView, ScrollView } from 'react-native';
import { Button } from '@/components/button';
import Wordmark from '@/components/wordmark';
import { useState } from 'react';

type TagsData = {
    tags: String[];
}

const listOfTags = ["Premed","Prelaw","Judaism","Christianity","Hinduism","Islam","Latin America","African American","Asian American","LGBTQ","Performing Arts","Visual Arts","Creative Writing","Music","Soccer","Hiking","Climbing","Lacrosse","Mathematics","Physics","Biology","Chemistry","Environmental Science","Geology","Neuroscience","Psychology","Software Engineering","Artificial Intelligence","DataScience","Mechanical Engineering","ElectricalEngineering","Industrial Engineering","Volunteerism","Environmental Advocacy","Human Rights","Community Outreach","Journalism","Broadcasting","Film","Public Relations","Other"]

const Tags = () => {
    return (
        <SafeAreaView>
            <View className="px-[8%] pt-[4%]">
            <View className="pb-[8%]"><Wordmark/></View>
            <Text className="text-5xl font-bold">What are you interested in?</Text>
            <Text className="text-xl pt-[5%]">Select one or more</Text>
            <ScrollView className="h-[61.5%] pt-[5%]">
                <View className="flex-row flex-wrap">
                {listOfTags.map((text, index) => (
                <Button 
                variant="outline"
                size="tags">{text}</Button>))}
                </View>
            </ScrollView>
            <View className="flex-row justify-end pt-[5%]">
                <Button
                    size="lg"
                    variant="default"
                >Finish</Button></View>
            </View>
        </SafeAreaView>
    );
}

export default Tags;

