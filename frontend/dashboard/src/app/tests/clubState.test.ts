// import { setClub, updateClub, resetClub, clubReducer } from '../../store/slices/userSlice';
// import { Club, CreateClubRequestBody, UpdateClubRequestBody } from '../../types/club';

// describe('club slice', () => {
//   const initialState: Club = {
//     userId: '',
//     clubId: '',
//     clubName: '',
//     accessToken: '',
//     csrfToken: '',
//     preview: '',
//     description: '',
//     weeklyTimeCommitment: 0,
//     oneWordToDescribeUs: '',
//     logo: '',
//   }

//   it('set club', () => {
//     const club: CreateClubRequestBody = {
//       userId: '123',
//       name: 'Generate',
//       preview: 'preview',
//       description: 'description',
//       weeklyTimeCommitment: 10,
//       oneWordToDescribeUs: 'innovative',
//       logo: 'logo.png',
//     }

//     const action = setClub(club);
//     const nextState = clubReducer(initialState, action);
//     expect(nextState).toEqual({
//       ...initialState,
//       ...club
//     })
//   })

//   it('update club', () => {
//     const currentState: Club = {
//       userId: '123',
//       clubId: '456',
//       clubName: 'Generate Development Studio',
//       accessToken: 'token',
//       csrfToken: 'csrf',
//       preview: 'preview',
//       description: 'description',
//       weeklyTimeCommitment: 20,
//       oneWordToDescribeUs: 'daring',
//       logo: 'logo.png',
//     }

//     const updatedFields: UpdateClubRequestBody = {
//       preview: 'new preview',
//       weeklyTimeCommitment: 10,
//     }

//     const action = updateClub(updatedFields);
//     const nextState = clubReducer(currentState, action);
//     expect(nextState).toEqual({
//       ...currentState,
//       ...updatedFields,
//     })
//   })

//   it('reset club', () => {
//     const currentState: Club = {
//       userId: '123',
//       clubId: '456',
//       clubName: 'Generate Development Studio',
//       accessToken: 'token',
//       csrfToken: 'csrf',
//       preview: 'preview',
//       description: 'description',
//       weeklyTimeCommitment: 20,
//       oneWordToDescribeUs: 'daring',
//       logo: 'logo.png',
//     }

//     const action = resetClub();
//     const nextState = clubReducer(currentState, action);
//     expect(nextState).toEqual(initialState);
//   })
// })