import { Box, Text, ColorName, Colors} from "@/app/(design-system)"
import RecruitmentCycleIcon from '@/assets/svg/recruitment-cycle.svg';
import RecruitmentStatusIcon from '@/assets/svg/recruiting-status.svg';
import ApplicationTypeIcon from '@/assets/svg/application-type.svg';
import { RecruitmentCycle, RecruitmentType } from "@generatesac/lib";
import { SvgProps } from 'react-native-svg';
import { ReactElement } from 'react';

interface RecruitmentInfoProps {
    color: ColorName,
    recruitmentCycle: RecruitmentCycle,
    recruitingType: RecruitmentType,
    isRecruiting?: boolean,
    recruitmentText: string,
}

interface RecruitmentItemProps {
    icon: ReactElement<SvgProps>,
    title: string,
    text?: string | undefined,
}

const capitalize = (str: string) => {
    if (!str) return '';
    return str.charAt(0).toUpperCase() + str.slice(1);
}

export const RecruitmentInfo = ({recruitmentText, color, recruitmentCycle, recruitingType, isRecruiting = false}: RecruitmentInfoProps) => {
    const RecruitmentItem = ({icon, title, text}: RecruitmentItemProps) => {
        return (
            <Box alignItems="center" borderWidth={1} borderColor={color} borderRadius={8} width={'30%'}>
                <Box paddingTop="l" paddingBottom="l" gap="xs" flexDirection="column" alignItems="center">
                    {icon}
                    <Text style={{textAlign: 'center', fontWeight: '600', fontSize: 14 }}>{capitalize(title)}</Text>
                    {text && <Text variant="caption-1" style={{textAlign: 'center'}}>{capitalize(text)}</Text>}
                </Box>
            </Box>
        )
    }

    return (
        <Box padding="l" flexDirection="row" gap="m" alignItems="stretch" justifyContent="space-evenly">
            <RecruitmentItem icon={<RecruitmentCycleIcon fill={Colors[color]}/>} title="Recruitment Cycle" text={recruitmentCycle}/>
            <RecruitmentItem icon={<RecruitmentStatusIcon fill={Colors[color]}/>} title={isRecruiting ? "Currently Recruiting" : "Currently Not Recruiting"} text={recruitmentText}/>
            <RecruitmentItem icon={<ApplicationTypeIcon fill={Colors[color]}/>} title="Application Type" text={recruitingType}/>
        </Box>
    )
}