import clsx, { ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';

/**
 * Nativewind CSS classnames generator
 * @param inputs - a list of classnames
 * @returns a string of Nativewind CSS classnames
 */
export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

// list of items for dropdown menu
export type Item = {
    label: string;
    value: string;
};

// list of graduation year
export const graduationYear = () => {
    var year = new Date().getFullYear();
    const graduationYear: Item[] = [];
    for (let i = 0; i < 5; i++) {
        graduationYear.push({
            label: String(year + i),
            value: String(year + i)
        });
    }
    return graduationYear;
};

export const college = [
    {label: "College of Arts, Media and Design", value: "CAMD"},
    {label: "D'Amore-McKim School of Business", value: "DMSB"},
    {label: "Khoury College of Computer Sciences", value: "KCCS"},
    {label: "College of Engineering", value: "CE"},
    {label: "Bouvé College of Health Sciences", value: "BCHS"},
    {label: "School of Law", value: "SL"},
    {label: "College of Professional Studies", value: "CPS"},
    {label: "College of Science", value: "CS"},
    {label: "College of Social Sciences and Humanities", value: "CSSH"},
]