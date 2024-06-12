import { Controller, useForm } from 'react-hook-form';
import { Alert } from 'react-native';

import { zodResolver } from '@hookform/resolvers/zod';
import { ZodError, z } from 'zod';

import { Item } from '@/src/types/item';

import { Box, Text } from '../../(design-system)';
import { MultiSelect, SelectOne } from '../../(design-system)';
import { Button } from '../../(design-system)/components/Button/Button';
import { TextBox } from '../../(design-system)/components/Textbox/Textbox';
import { COLLEGE, graduationYear, MAJORS } from '../../../consts/user-info';

type UserInfoData = {
    major: Item[];
    college: Item;
    nuid: string;
    startYear: Item;
    graduationYear: Item;
};

const itemSchema = z.object({
    key: z.string(),
    value: z.string()
});

const userInfoSchema = z.object({
    college: itemSchema,
    major: z.array(itemSchema),
    nuid: z.string().refine(
        (value) => {
            return /^\d{9}$/.test(value) && value.startsWith('00');
        },
        {
            message: 'Invalid NUID'
        }
    ),
    graduationYear: itemSchema,
    startYear: itemSchema
});

interface UserInfoProps {
    onPress: () => void;
}

export const UserInfo: React.FC<UserInfoProps> = ({ onPress }) => {
    const {
        control,
        handleSubmit,
        trigger,
        formState: { isValid, errors }
    } = useForm<UserInfoData>({
        resolver: zodResolver(userInfoSchema),
        mode: 'onTouched'
    });

    const onSubmit = (data: UserInfoData) => {
        onPress();
        try {
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
        <Box
            paddingHorizontal="xl"
            justifyContent="center"
            alignItems="flex-start"
            flexDirection="column"
            gap="xs"
        >
            <Text paddingBottom="m" fontSize={40} variant="header-1">
                Let's learn more about you
            </Text>
            <Controller
                control={control}
                render={({ field: { onChange, value } }) => (
                    <MultiSelect
                        onChange={onChange}
                        title="Major"
                        data={MAJORS}
                        limit={3}
                        value={value || []}
                        placeholder="Select up to 3 majors"
                    />
                )}
                name="major"
            />
            <Controller
                control={control}
                render={({ field: { onChange, value } }) => (
                    <SelectOne
                        value={value}
                        onChange={onChange}
                        placeholder="Select home college"
                        title="College"
                        data={COLLEGE}
                    />
                )}
                name="college"
            />
            <Controller
                control={control}
                render={({ field: { onChange, value } }) => (
                    <TextBox
                        onChangeText={(text: string) => {
                            onChange(text);
                            trigger('nuid');
                        }}
                        value={value}
                        error={errors.nuid ? errors.nuid.message : ''}
                        placeholder="Enter your NUID"
                        title="NUID"
                    />
                )}
                name="nuid"
            />
            <Controller
                control={control}
                render={({ field: { onChange, value } }) => (
                    <SelectOne
                        value={value}
                        onChange={onChange}
                        placeholder="Select start year"
                        title="Start Year"
                        data={graduationYear()}
                    />
                )}
                name="startYear"
            />
            <Controller
                control={control}
                render={({ field: { onChange, value } }) => (
                    <SelectOne
                        value={value}
                        onChange={onChange}
                        placeholder="Select intended graduation year"
                        title="Intended Graduation Year"
                        data={graduationYear()}
                    />
                )}
                name="graduationYear"
            />
            <Box paddingTop="s">
                <Button
                    disabled={!isValid}
                    onPress={handleSubmit(onSubmit)}
                    variant="standardButton"
                >
                    Continue
                </Button>
            </Box>
        </Box>
    );
};
