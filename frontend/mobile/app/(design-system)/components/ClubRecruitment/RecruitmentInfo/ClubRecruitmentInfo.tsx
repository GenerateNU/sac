import React from 'react';

import { faCalendar } from '@fortawesome/free-regular-svg-icons/faCalendar';
import { faComments } from '@fortawesome/free-solid-svg-icons/faComments';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { RecruitmentCycle, RecruitmentType } from '@generatesac/lib';

import { Box, SACColors, createStyles } from '@/app/(design-system)';

import { RecruitmentItem } from '../RecruitmentItem/ClubRecruitmentItem';

interface RecruitmentInfoProps {
    color: SACColors;
    recruitmentCycle: RecruitmentCycle;
    recruitingType: RecruitmentType;
    isRecruiting?: boolean;
    recruitmentText: string;
}

export const RecruitmentInfo = ({
    recruitmentText,
    color,
    recruitmentCycle,
    recruitingType,
    isRecruiting = false
}: RecruitmentInfoProps) => {
    return (
        <Box {...styles.recruitmentInfo}>
            <RecruitmentItem
                icon={faCalendar}
                title="Recruitment Cycle"
                text={recruitmentCycle}
                color={color}
            />
            <RecruitmentItem
                icon={faPenToSquare}
                title={
                    isRecruiting
                        ? 'Currently Recruiting'
                        : 'Currently Not Recruiting'
                }
                text={recruitmentText}
                color={color}
            />
            <RecruitmentItem
                icon={faComments}
                title="Application Type"
                text={recruitingType}
                color={color}
            />
        </Box>
    );
};

const styles = createStyles({
    recruitmentInfo: {
        padding: 'l',
        flexDirection: 'row',
        gap: 'm',
        alignItems: 'stretch',
        justifyContent: 'space-evenly'
    }
});
