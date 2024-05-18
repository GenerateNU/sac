// export default function (api) {
//     api.cache(true);
//     return {
//         presets: ['babel-preset-expo']
//     };
// };

module.exports = function(api) {
    api.cache(true);
    return {
      presets: ['babel-preset-expo'],
    };
};
