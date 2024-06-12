import { MAJORS, COLLEGE, GRADUATION_CYCLE } from "../consts/user-info";
import { UserInfoData } from "../types/userInfo";

type MockUserType = {
    name: string | null,
    email: string,
    major0: string | null,
    major1: string | null,
    major2: string | null,
    college: string | null,
    startYear: number | null,
    graduationYear: number | null,
    graduationCycle: string | null,
};

export const convertUserData = (userData: MockUserType): Partial<UserInfoData> => {
    const data: Partial<UserInfoData> = {
        name: userData.name || '',
        email: userData.email || '',
    };
    if (userData.college !== null) data.college = userData.college ? COLLEGE.find((item) => userData.college === item.key) || { key: '', value: '' } : { key: '', value: '' };
    if (userData.graduationCycle !== null) data.graduationCycle = userData.graduationCycle ? GRADUATION_CYCLE.find((item) => item.key === userData.graduationCycle) || { key: '', value: '' } : { key: '', value: '' };
    if (userData.startYear !== null) data.startYear = { key: String(userData.startYear), value: String(userData.startYear) };
    if (userData.graduationYear !== null) data.graduationYear = { key: String(userData.graduationYear), value: String(userData.graduationYear) };
    const majors = [userData.major0, userData.major1, userData.major2]
        .filter((major) => major !== null)
        .map((major) => {
            const majorItem = MAJORS.find((item) => item.key === major);
            return majorItem ? majorItem : { key: '', value: '' };
        });
    if (majors.length > 0) {
        data.major = majors;
    }
    return data;
};
