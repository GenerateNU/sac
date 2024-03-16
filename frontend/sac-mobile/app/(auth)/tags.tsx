import { Text, View, SafeAreaView, ScrollView } from 'react-native';
import { Button } from '@/components/button';
import Wordmark from '@/components/wordmark';
import React, { useState } from 'react';


type TagsData = {
    tags: String[];
}

const listOfTags = ["Pre-med","Prelaw","Judaism","Christianity","Hinduism","Islam","Latin America","African American","Asian American","LGBTQ","Performing Arts","Visual Arts","Creative Writing","Music","Soccer","Hiking","Climbing","Lacrosse","Mathematics","Physics","Biology","Chemistry","Environmental Science","Geology","Neuroscience","Psychology","Software Engineering","Artificial Intelligence","Data Science","Mechanical Engineering","Electrical Engineering","Industrial Engineering","Volunteerism","Environmental Advocacy","Human Rights","Community Outreach","Journalism","Broadcasting","Film","Public Relations","Other"]

const Tags = () => {
    const [toggle,setToggle] = React.useState(false);
    const toggleIt =()=>{
        setToggle(!toggle)
    }
    return (
        <SafeAreaView>
            <View className="px-[8%] pt-[4%]">
            <View className="pb-[3%]"><Wordmark/></View>
            <Text className="text-5xl font-bold">What are you interested in?</Text>
            <Text className="text-xl pt-[3%] pb-[4%]">Select one or more</Text>
            <ScrollView className="h-[62%] pt-[3%]">
                <View className="flex-row flex-wrap mb-[2%]">
                {listOfTags.map((text, index) => (
                <Button 
                variant={toggle? "default" : "outline"}
                size="tags">{text}</Button>))}
                </View>
            </ScrollView>
            <View className="flex-row justify-end pt-[5%]">
                <Button
                    size="lg"
                >Finish</Button></View>
            </View>
        </SafeAreaView>
    );
}

export default Tags;

