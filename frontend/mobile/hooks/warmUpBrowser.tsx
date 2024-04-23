import React from 'react';

import * as WebBrowser from 'expo-web-browser';

export const useWarmUpBrowser = () => {
    React.useEffect(() => {
        WebBrowser.warmUpAsync();
        return () => {
            WebBrowser.coolDownAsync();
        };
    }, []);
};
