module.exports = function async(api) {
    api.cache(true);
    return {
        presets: ['babel-preset-expo'],
        plugins: [['module:react-native-dotenv']]
    };
};
