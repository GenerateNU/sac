"use client";
import { Button } from "@/components/ui/button";
import {InputWithLabel} from "@/components/ui/inputWithLabel";
import { SelectScrollable } from "@/components/ui/selectScrollable";
import { College, Interest } from "@/const/dummyConst";
import MultipleSelector from '@/components/ui/multiSelect';
import { useForm, Controller } from "react-hook-form";
import { z } from "zod";
import {Option} from "@/types/option";
import { useState } from 'react';
import { zodResolver } from "@hookform/resolvers/zod";
import { collegeEnum } from "../types/club";

type FormData = {
  firstName: string;
  lastName: string;
  email: string;
  college: string;
  interests: Option[];
};

const formSchema = z.object({
  firstName: z.string()
    .min(2, "First name must be at least 2 characters")
    .refine(value => /^[a-zA-Z\s]+$/.test(value), {
      message: "First name can only contain alphabetic characters and spaces",
    }),
  lastName: z.string()
    .min(2, "Last name must be at least 2 characters")
    .refine(value => /^[a-zA-Z]+$/.test(value), {
      message: "Last name can only contain alphabetic characters",
    }),
  email: z.string()
    .email("Please enter a valid email")
    .refine(value => /^[a-zA-Z0-9.]+@[a-zA-Z0-9]+\.[a-zA-Z]+$/.test(value), {
        message: "Please enter your Northeastern email",
    }),
  college: z.nativeEnum(collegeEnum).refine(value => value !== undefined, {
    message: "Please select a college",
  }),
  interests: z.array(z.object({
    label: z.string(),
    value: z.string(),
  })).nonempty("At least one interest is required"),
});

export default function Form() {
  const [formDataJson, setFormDataJson] = useState<string>('');
  const {
    control,
    handleSubmit,
    trigger,
    formState: { errors }
  } = useForm<FormData>({
    resolver: zodResolver(formSchema),
    mode: 'onTouched',
  });

  const onSubmit = (data: FormData) => {
    setFormDataJson(JSON.stringify(data));
  };

  return (
    <main className="flex flex-col items-center justify-between min-h-screen p-24">
      <form onSubmit={handleSubmit(onSubmit)} className="max-w-md w-full space-y-10 flex flex-col">
        <div>
          <Controller
            name="firstName"
            control={control}
            render={({ field: { onChange, onBlur, value } }) => (
              <>
                <InputWithLabel
                  value={value}
                  onChange={(event) => {
                    onChange(event.target.value);
                    trigger("firstName");
                  }}
                  onBlur={(e) => {
                    onBlur();
                    trigger("firstName");
                  }}
                  type="text"
                  id="first-name"
                  placeholder="Please enter your first name"
                  label="First Name"
                />
                {errors.firstName && <p className="mt-1 text-sm text-red-500">{errors.firstName.message}</p>}
              </>
            )}
          />
        </div>
        <div>
          <Controller
            name="lastName"
            control={control}
            render={({ field: { onChange, onBlur, value } }) => (
              <>
                <InputWithLabel
                  value={value}
                  onChange={(event) => {
                    onChange(event.target.value);
                    trigger("lastName");
                  }}
                  onBlur={(e) => {
                    onBlur();
                    trigger("lastName");
                  }}
                  type="text"
                  id="last-name"
                  placeholder="Please enter your last name"
                  label="Last Name"
                />
                {errors.lastName && <p className="mt-1 text-sm text-red-500">{errors.lastName.message}</p>}
              </>
            )}
          />
        </div>
        <div>
          <Controller
            name="email"
            control={control}
            render={({ field: { onChange, onBlur, value } }) => (
              <>
                <InputWithLabel
                  value={value}
                  onChange={(event) => {
                    onChange(event.target.value);
                    trigger("email");
                  }}
                  onBlur={(e) => {
                    onBlur();
                    trigger("email");
                  }}
                  type="email"
                  id="email"
                  placeholder="Please enter your email"
                  label="Email"
                />
                {errors.email && <p className="mt-1 text-sm text-red-500">{errors.email.message}</p>}
              </>
            )}
          />
        </div>
        <div>
          <Controller
            name="college"
            control={control}
            render={({ field: { onChange, onBlur, value } }) => (
              <>
                <SelectScrollable
                  value={value}
                  data={College}
                  onChange={onChange}
                  onBlur={(e) => {
                    onBlur();
                    trigger("college");
                  }}
                  label="College"
                  placeholder="Select your home college"
                />
                {errors.college && <p className="mt-1 text-sm text-red-500">{errors.college.message}</p>}
              </>
            )}
          />
        </div>
        <div>
          <Controller
            name="interests"
            control={control}
            render={({ field: { onChange, onBlur, value } }) => (
              <>
                <MultipleSelector
                  value={value}
                  onChange={onChange}
                  onBlur={(e) => {
                    onBlur();
                    trigger("interests");
                  }}
                  label="Tags"
                  placeholder="Select at least one interest"
                  options={Interest}
                />
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