import React, { forwardRef, useCallback, useState } from 'react';

import BottomSheet, { BottomSheetBackdrop } from '@gorhom/bottom-sheet';

import { Box, Text } from '@/app/(design-system)';
import { SelectOne } from '@/app/(design-system)';
import { Button } from '@/app/(design-system)/components/Button/Button';
import { Item } from '@/types/item';

import { Divider } from './divider';

type Ref = BottomSheet;
const SaveEventText = `By saving this event, you are automatically signed up for notifications and this event will be added to your calendar.`;

interface RegisterSheetProps {
    eventType: 'hybrid' | 'in-person' | 'virtual';
    // eventId: string,
}

const RegisterBottomSheet = forwardRef<Ref, RegisterSheetProps>(
    ({ eventType }, ref) => {
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

        const [format, setFormat] = useState<Item>();
        const [sheet, setSheet] = useState('register');
        const [attending, setAttending] = useState<Item>();

        const handleRegisterPress = () => {
            if (attending?.key === 'maybe') {
                setSheet('save');
            } else {
                if (eventType === 'hybrid') {
                    setSheet('attend');
                } else {
                    setSheet('save');
                }
            }
        };

        return (
            <BottomSheet
                ref={ref}
                index={-1}
                snapPoints={sheet === 'save' ? ['40%'] : ['50%']}
                enablePanDownToClose={true}
                backgroundStyle={{ backgroundColor: 'white' }}
                backdropComponent={renderBackdrop}
            >
                {sheet === 'register' && (
                    <Register
                        onChange={setAttending}
                        onPress={handleRegisterPress}
                        value={attending}
                    />
                )}
                {sheet === 'attend' && (
                    <AttendanceType
                        onPress={() => setSheet('save')}
                        onChange={setFormat}
                        value={format}
                    />
                )}
                {sheet === 'save' && <SaveEvent />}
            </BottomSheet>
        );
    }
);

interface RegisterProps {
    onPress: () => void;
    onChange: (value: Item) => void;
    value: Item | undefined;
}

const Register: React.FC<RegisterProps> = ({ onPress, onChange, value }) => {
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
                <SelectOne
                    data={[
                        { key: 'yes', value: 'Yes' },
                        { key: 'maybe', value: 'Maybe' }
                    ]}
                    onChange={onChange}
                    placeholder="Please select an option"
                    title="Will you be joining us?"
                />
            </Box>
            <Button
                onPress={onPress}
                disabled={!value}
                variant="standardButton"
            >
                Continue
            </Button>
        </Box>
    );
};

const SaveEvent = () => {
    return (
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
                <Button variant="standardButton">Add to Google Calendar</Button>
                <Button variant="standardButton">Add to Apple Calendar</Button>
            </Box>
        </Box>
    );
};

interface AttendanceTypeProps {
    onPress: () => void;
    onChange: (value: Item) => void;
    value: Item | undefined;
}

const AttendanceType: React.FC<AttendanceTypeProps> = ({
    onPress,
    onChange,
    value
}) => {
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
                <SelectOne
                    data={[
                        { key: 'in-person', value: 'In-person' },
                        { key: 'virtual', value: 'Virtual' }
                    ]}
                    onChange={onChange}
                    placeholder="Will you join in-person or virtually?"
                />
            </Box>
            <Button
                onPress={onPress}
                disabled={!value}
                variant="standardButton"
            >
                Continue
            </Button>
        </Box>
    );
};

export { RegisterBottomSheet };
