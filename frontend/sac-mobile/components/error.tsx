import React from 'react'; 

import {Text, DimensionValue} from 'react-native'; 

interface ErrorMessageProps  {
    message: string | undefined; 
}

const Error = ({message}: ErrorMessageProps) => {
    return (
        <Text className="text-red-600 pt-[2%]">{message}</Text>
    )
}

export default Error;