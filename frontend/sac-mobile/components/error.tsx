import React from 'react'; 

import {Text, DimensionValue} from 'react-native'; 

type ErrorMessage = {
    message: string | undefined; 
}

const Error = (props: ErrorMessage) => {
    return (
        <Text className="text-red-600 pt-[2%]">{props.message}</Text>
    )
}

export default Error;