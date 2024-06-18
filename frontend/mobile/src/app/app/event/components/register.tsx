import React, { forwardRef, useCallback, useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { Alert, Linking } from 'react-native';

import { EventType } from '@generatesac/lib';
import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';
import { zodResolver } from '@hookform/resolvers/zod';
import { ZodError, z } from 'zod';

import { Box, Text } from '@/src/app/design-system';
import { SelectOne } from '@/src/app/design-system';
import { Button } from '@/src/app/design-system/components/Button/Button';
import { CalendarLink } from '@/src/types/calendarLink';
import { Item } from '@/src/types/item';
import { createGoogleCalendarLink } from '@/src/utils/string';

import { Divider } from './divider';

type Ref = BottomSheet;

const SaveEventText = `By saving this event, you are automatically signed up for notifications and this event will be added to your calendar.`;

interface RegisterSheetProps {
    eventType: EventType;
    startTime: string;
    endTime: string;
    location: string;
    eventName: string;
    eventDetail: string;
}

const RegisterBottomSheet = forwardRef<Ref, RegisterSheetProps>(
    (
        { eventType, eventName, location, eventDetail, startTime, endTime },
        ref
    ) => {
        const [sheet, setSheet] = useState('register');

        const renderBackdrop = useCallback(
            (props: any) => (
                <BottomSheetBackdrop
                    appearsOnIndex={0}
                    disappearsOnIndex={-1}
                    {...props}
                />
            ),
            []
        );

        const CALENDAR_INFO: CalendarLink = {
            eventDetail: eventDetail,
            eventName: eventName,
            location: location,
            startTime: startTime,
            endTime: endTime
        };

        return (
            <BottomSheet
                ref={ref}
                index={-1}
                snapPoints={sheet === 'save' ? ['40%'] : ['50%']}
                enablePanDownToClose
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
            >
                {sheet === 'register' && (
                    <Register setSheet={setSheet} eventType={eventType} />
                )}
                {sheet === 'attend' && (
                    <AttendanceType onPress={() => setSheet('save')} />
                )}
                {sheet === 'save' && (
                    <SaveEvent
                        googleCalendar={createGoogleCalendarLink(CALENDAR_INFO)}
                    />
                )}
            </BottomSheet>
        );
    }
);

interface RegisterProps {
    setSheet: React.Dispatch<React.SetStateAction<string>>;
    eventType: EventType;
}

type UserData = {
    select: Item;
};

const userDataSchema = z.object({
    key: z.string(),
    value: z.string()
});

const userInfoSchema = z.object({
    select: userDataSchema
});

const Register: React.FC<RegisterProps> = ({ setSheet, eventType }) => {
    const {
        control,
        handleSubmit,
        formState: { isValid }
    } = useForm<UserData>({
        resolver: zodResolver(userInfoSchema),
        mode: 'onTouched'
    });

    const onSubmit = (data: UserData) => {
        try {
            if (data.select.key === 'maybe') {
                setSheet('save');
            } else if (data.select.key === 'yes' && eventType === 'hybrid') {
                setSheet('attend');
            } else {
                setSheet('save');
            }
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <Box marginHorizontal="xl" flexDirection="column" gap="s">
            <Text
                paddingTop="s"
                textAlign="center"
                variant="subheader-2"
                fontWeight="500"
            >
                Register
            </Text>
            <Divider />
            <Box paddingTop="m" height={250}>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <SelectOne
                            data={[
                                { key: 'yes', value: 'Yes' },
                                { key: 'maybe', value: 'Maybe' }
                            ]}
                            onChange={onChange}
                            value={value}
                            placeholder="Please select an option"
                            title="Will you be joining us?"
                        />
                    )}
                    name="select"
                />
            </Box>
            <Button
                onPress={handleSubmit(onSubmit)}
                disabled={!isValid}
                variant="standardButton"
            >
                Continue
            </Button>
        </Box>
    );
};

interface SaveEventProps {
    googleCalendar: string;
    appleCalendar?: string;
}
const SaveEvent: React.FC<SaveEventProps> = ({ googleCalendar }) => (
    <Box marginHorizontal="xl" flexDirection="column" gap="s">
        <Text
            paddingTop="s"
            textAlign="center"
            variant="subheader-2"
            fontWeight="500"
        >
            Save Event
        </Text>
        <Divider />
        <Text paddingTop="s">{SaveEventText}</Text>
        <Box marginTop="s" flexDirection="column" gap="s">
            <Button
                onPress={() => Linking.openURL(googleCalendar)}
                variant="standardButton"
            >
                <Text>Add to Google Calendar</Text>
            </Button>
            <Button variant="standardButton">Add to Apple Calendar</Button>
        </Box>
    </Box>
);

interface AttendanceTypeProps {
    onPress: () => void;
}

const AttendanceType: React.FC<AttendanceTypeProps> = ({ onPress }) => {
    const {
        control,
        handleSubmit,
        formState: { isValid }
    } = useForm<UserData>({
        resolver: zodResolver(userInfoSchema),
        mode: 'onTouched'
    });

    const onSubmit = (data: UserData) => {
        try {
            onPress();
            console.log(data);
        } catch (error) {
            if (error instanceof ZodError) {
                Alert.alert('Validation Error', error.errors[0].message);
            } else {
                console.error('An unexpected error occurred:', error);
            }
        }
    };

    return (
        <Box marginHorizontal="xl" flexDirection="column" gap="s">
            <Text
                paddingTop="s"
                textAlign="center"
                variant="subheader-2"
                fontWeight="500"
            >
                Register
            </Text>
            <Divider />
            <Text paddingTop="s">
                This event is available for both in-person and virtual
                attendance. Please choose your preferred option.
            </Text>
            <Box height={170}>
                <Controller
                    control={control}
                    render={({ field: { onChange, value } }) => (
                        <SelectOne
                            data={[
                                { key: 'in-person', value: 'In-person' },
                                { key: 'virtual', value: 'Virtual' }
                            ]}
                            onChange={onChange}
                            value={value}
                            placeholder="Will you join in-person or virtually?"
                        />
                    )}
                    name="select"
                />
            </Box>
            <Button
                onPress={handleSubmit(onSubmit)}
                disabled={!isValid}
                variant="standardButton"
            >
                Continue
            </Button>
        </Box>
    );
};

export { RegisterBottomSheet };
