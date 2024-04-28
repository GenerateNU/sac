import { Club } from '@/types/club';
import { Event } from '@/types/event';
import { Tag } from '@/types/tag';

import { clubs } from './clubs';
import { tags } from './tags';

const Locations = [
    '2435 E. North St., Greenville, SC',
    '1234 Main Rd., New York, NY',
    '5678 Elm St., Los Angeles, CA',
    '91011 Pine Ct., Chicago, IL',
    '1213 Oak Ave., Houston, TX',
    '1415 Maple Dr., Phoenix, AZ',
    '1617 Birch Ln., Philadelphia, PA',
    '1819 Cedar Rd., San Antonio, TX',
    '2021 Spruce St., San Diego, CA',
    '2223 Willow Dr., Dallas, TX'
];

const generateLoremIpsum = (paragraphs = 1, sentencesPerParagraph = 3) => {
    const loremIpsumText =
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.';
    const sentences = loremIpsumText.split('. ');

    const getRandomSentence = () => {
        return sentences[Math.floor(Math.random() * sentences.length)];
    };

    let paragraphsText = '';
    for (let i = 0; i < paragraphs; i++) {
        let paragraph = '';
        for (let j = 0; j < sentencesPerParagraph; j++) {
            paragraph += getRandomSentence() + ' ';
        }
        paragraphsText += paragraph.trim() + '\n\n';
    }

    return paragraphsText.trim();
};

export const events: Event[] = [];

for (let i = 1; i <= 10; i++) {
    events.push({
        id: i.toString(),
        updated_at: new Date(),
        created_at: new Date(),
        name: `Event ${i}`,
        start_time: new Date(),
        end_time: new Date(),
        location: Locations[i - 1],
        preview: `Preview ${i}`,
        content: generateLoremIpsum(
            Math.floor(Math.random() * 3) + 1,
            Math.floor(Math.random() * 3) + 1
        ),
        meeting_link: `http://www.example.com/event${i}`,
        is_recurring: false,
        event_type: 'open'
    });
}

// Event Hosts

export const eventHosts: { [key: string]: Club[] } = {};

for (let i = 1; i <= 10; i++) {
    const hostsArray: Club[] = [];
    const hostIndices: number[] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9].filter(
        (index) => index % i === 0
    );
    hostIndices.forEach((index) => {
        hostsArray.push(clubs[index]);
    });
    eventHosts[i.toString()] = hostsArray;
}

// Event Tags

export const eventTags: { [key: string]: Tag[] } = {};

for (let i = 1; i <= 10; i++) {
    const tagsArray: Tag[] = [];
    const tagIndices: number[] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9].filter(
        (index) => index % i === 0
    );
    tagIndices.forEach((index) => {
        tagsArray.push(tags[index]);
    });
    eventTags[i.toString()] = tagsArray;
}
