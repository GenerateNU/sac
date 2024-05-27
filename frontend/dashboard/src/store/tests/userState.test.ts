import userSlice, { setUser, resetUser, setAccessToken, resetAccessToken, UserState } from '../../store/slices/userSlice';

describe('user slice', () => {
  const initialState: UserState = {
    userId: '',
    email: '',
    clubId: '',
    clubName: '',
    clubLogo: '',
  }

  it('set user', () => {
    const user: UserState = {
      userId: '123',
      email: 'Generate',
      clubId: 'preview',
      clubName: 'description',
      clubLogo: 'www.logo.com',
    }

    const action = setUser(user);
    const nextState = userSlice(initialState, action);
    expect(nextState).toEqual({
      ...initialState,
      ...user
    });
  });

  it('reset user', () => {
    const currentState: UserState = {
        userId: '123',
        email: 'Generate',
        clubId: 'preview',
        clubName: 'description',
        clubLogo: 'www.logo.com',
        accessToken: 'innovative',
      }

      const action = resetUser();
      const nextState = userSlice(currentState, action);
      expect(nextState).toEqual(initialState);
  });

  it('set access token user', () => {
    const currentState: UserState = {
        userId: '123',
        email: 'Generate',
        clubId: 'preview',
        clubName: 'description',
        clubLogo: 'www.logo.com',
    }

    const accessToken = 'mew'

    const action = setAccessToken(accessToken);
    const nextState = userSlice(currentState, action);
    expect(nextState).toEqual({
        ...currentState,
        accessToken,
    });
  });

  it('reset access token user', () => {
    const currentState: UserState = {
        userId: '123',
        email: 'Generate',
        clubId: 'preview',
        clubName: 'description',
        clubLogo: 'www.logo.com',
        accessToken: 'mew',
    }

    const action = resetAccessToken();
    const nextState = userSlice(currentState, action);
    expect(nextState).toEqual({
        ...currentState,
        accessToken: undefined,
    });
  });
});
