"use client";
import {Button} from "@/components/ui/button";
import {InputWithLabel} from "@/components/ui/inputWithLabel";
import { SelectScrollable } from "@/components/ui/selectScrollable";
import { College, Interest } from "@/const/dummyConst";
import MultipleSelector from '@/components/ui/multiSelect';
import { useForm, Controller } from "react-hook-form";
import { z } from "zod";
import {Option} from "@/types/option";
import { useState } from 'react';

type FormData = {
  firstName: string;
  lastName: string;
  email: string;
  college: string;
  interests: Option[];
};

const formSchema = z.object({
  firstName: z.string().min(2),
  lastName: z.string().min(2),
  email: z.string().email({
    message: "Please enter your email"
  }),
  college: z.string().min(1),
  interests: z.array(z.object({
    label: z.string(),
    value: z.string()
  })).nonempty(),
})

export default function Home() {
  const [formDataJson, setFormDataJson] = useState<string>('');
  const {
    control,
    handleSubmit,
    formState: { errors }
} = useForm<FormData>();

const onSubmit = (data: FormData) => {
  formSchema.parse(data);
  const updatedData = {
    ...data,
    interests: data.interests.map(option => option.value)
  };
  setFormDataJson(JSON.stringify(updatedData));
};

  return (
    <main className="flex flex-col items-center justify-between min-h-screen p-24">
      <form onSubmit={handleSubmit(onSubmit)} className="max-w-md w-full space-y-10 flex flex-col">
        <div>
        <Controller
          name="firstName"
          control={control}
          rules={{
            required: 'First name is required',
            validate: (value) => {
              const isValid = /^[a-zA-Z\s]+$/.test(value);
              if (!isValid) {
                  return 'Please enter proper first name';
              } else if (value.length < 2) {
                  return 'First name must be at least 2 characters long';
              }
              return true;
            }
          }}
          render={({ field: { onChange, value } }) => (
            <>
              <InputWithLabel value={value} onChange={onChange} type="text" id="first-name" placeholder="Please enter your first name" label="First Name" />
              {errors.firstName && <p className="mt-1 text-sm text-red-500">{errors.firstName.message}</p>}
            </>
          )}
        />
        </div>
        <div>
        <Controller
          name="lastName"
          control={control}
          rules={{
            required: 'Last name is required',
            validate: (value) => {
              const isValid = /^[a-zA-Z]+$/.test(value);
              if (!isValid) {
                  return 'Please enter proper first name';
              } else if (value.length < 2) {
                  return 'Last name must be at least 2 characters long';
              }
              return true;
            }
          }}
          render={({ field: { onChange, value } }) => (
            <>
              <InputWithLabel value={value} onChange={onChange} type="text" id="last-name" placeholder="Please enter your last name" label="Last Name" />
              {errors.lastName && <p className="mt-1 text-sm text-red-500">{errors.lastName.message}</p>}
            </>
          )}
        />
        </div>
        <div>
        <Controller
          name="email"
          control={control}
          rules={{
            required: 'Email is required',
            validate: (value) => {
              if (!value.endsWith('@northeastern.edu')) {
                  return 'Please enter your Northeastern email';
              }
              return true;
            }
          }}
          render={({ field }) => (
            <>
              <InputWithLabel {...field} type="email" id="email" placeholder="Please enter your email" label="Email" />
              {errors.email && <p className="mt-1 text-sm text-red-500">{errors.email.message}</p>}
            </>
          )}
        />
        </div>
        <div>
        <Controller
          name="college"
          control={control}
          rules={{ required: 'College is required' }}
          render={({ field: { onChange, value } }) => (
            <>
              <SelectScrollable value={value} data={College} onChange={onChange} label="College" placeholder="Select your home college" />
              {errors.college && <p className="mt-1 text-sm text-red-500">{errors.college.message}</p>}
            </>
          )}
        />
        </div>
        <div>
        <Controller
          name="interests"
          control={control}
          rules={{ required: 'Select at least one interest' }}
          render={({ field }) => (
            <>
              <MultipleSelector {...field} label="Tags" placeholder="Select your tags" options={Interest} />
              {errors.interests && <p className="text-sm text-red-500">{errors.interests.message}</p>}
            </>
          )}
        />
        </div>
        <Button type="submit">Submit Form</Button>
      </form>
      {formDataJson !== '' && <p>{formDataJson}</p>}
    </main>
  );
}
