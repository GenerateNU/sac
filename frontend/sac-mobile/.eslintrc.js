module.exports = {
  root: true,
  extends: ['@react-native-community', 'plugin:prettier/recommended'],
  plugins: ['prettier', 'jest'],
  rules: {
    'prettier/prettier': 'error',
    'react/react-in-jsx-scope': 'off'
  },
  env: {
    'jest/globals': true
  }
};
