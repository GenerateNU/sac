# React Style Guide

## Table of Contents

1. [General](#general)
2. [Naming](#naming)
3. [Imports](#imports)
4. [Components](#components-and-screens--pages)

## General

- Use functional components over class components
- Use Arrow functions over functions
    - Unless defining a class or an object only use functions
- Use `camelCase` for variables, functions, and file names
- Use `PascalCase` for components and their files
- Use `UPPER_SNAKE_CASE` for constants
- Name hooks with the `use` prefix

## Naming

- Use descriptive names for variables, functions, and components
- Use `is` or `has` for boolean variables
- Use `get` for functions that return a value
- Use `set` for functions that set a value
- Use `handle` for functions that handle events
- Use `on` for functions that are event handlers

## Imports

- Import from `react` and `react-dom` first

```js
import React from 'react';
```

- Sort imports 
1. Packages from `react` or `react-native`
2. Packages from `node_modules` 
3. Packages from `@/`
4. Components from `src`

```js   
import React, { useState } from 'react';

import { CopyrightIcon } from "lucide-react"

import { useQuery } from '@/hooks/query';

import { Button } from '@/components/Button';
```

- Use absolute imports

```js
import { Button } from '@/components/Button';
```

## Components and Screens / Pages

- Use `PascalCase` for component names and `camelCase` for props
- Use interfaces to define the types of props
- Inline export components
- Only reusable components should be in the `components` folder
    - Components that are only used in one part of the app should be in a `_components` folder in that part of the app
    - Avoid hardcoding properties
    - Never use `any` as a type
- Avoid using `Component` or `Screen` as an modifier for a file name as it is implied

```js
// Button.tsx

interface ButtonProps {
  text: string;
  className?: string;
  onClick: () => void;
}

export const Button = ({ text, className, onClick }: ButtonProps) => {
  return (
    <button onClick={onClick} className={className}>
      {text}
    </button>
  );
};
```

- For screens or pages, use default exports

```js
// Home.tsx

interface HomeProps {
  name: string;
}

const Home = ({ name }: HomeProps) => {
  return <Text>Hello, {name}!</Text>;
}

export default Home;
```
