import { Club } from '@/types/club';
import { Contact } from '@/types/contact';

export const clubs: Club[] = [];

for (let i = 1; i <= 10; i++) {
    const club: Club = {
        id: i.toString(),
        updated_at: new Date(),
        created_at: new Date(),
        name: `Club ${i}`,
        preview: `Preview ${i}`,
        description: `Description ${i}`,
        num_members: i * 10,
        is_recruiting: i % 2 === 0,
        recruitment_cycle:
            i % 4 === 0
                ? 'always'
                : i % 3 === 0
                  ? 'fallSpring'
                  : i % 2 === 0
                    ? 'spring'
                    : 'fall',
        recruitment_type:
            i % 3 === 0
                ? 'unrestricted'
                : i % 2 === 0
                  ? 'tryout'
                  : 'application',
        application_link: `https://example.com/club${i}`,
        logo: `https://example.com/club${i}.png`
    };
    clubs.push(club);
}

export const contacts: Contact[] = [];

for (let i = 1; i <= 10; i++) {
    const contact: Contact = {
        id: i.toString(),
        created_at: new Date(),
        updated_at: new Date(),
        content: `Contact ${i}`,
        type: Math.random() > 0.5 ? 'email' : 'slack'
    };
    contacts.push(contact);
}

export const clubContacts: { [key: string]: Contact[] } = {};

for (let i = 1; i <= 10; i++) {
    const clubContactsArray: Contact[] = [];
    const contactIndices: number[] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9].filter(
        (index) => index % i === 0
    );
    contactIndices.slice(0, 2).forEach((index) => {
        clubContactsArray.push(contacts[index]);
    });
    clubContacts[i.toString()] = clubContactsArray;
}
