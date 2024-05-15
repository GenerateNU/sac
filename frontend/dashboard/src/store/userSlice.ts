import { createSlice } from "@reduxjs/toolkit";
import { UserState } from "@/src/store/store.types";

const initialState: UserState = {
  email: "",
  clubId: "",
  clubName: "",
  clubLogo: "",
  googleCSRFToken: "",
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser: (state, action) => {
      state.email = action.payload.email;
      state.clubId = action.payload.clubId;
      state.clubName = action.payload.clubName;
      state.clubLogo = action.payload.clubLogo;
      state.googleCSRFToken = action.payload.googleCSRFToken;
    },
    resetUser: (state) => {
      state.email = initialState.email;
      state.clubId = initialState.clubId;
      state.clubName = initialState.clubName;
      state.clubLogo = initialState.clubLogo;
      state.googleCSRFToken = initialState.googleCSRFToken;
    },
    setGoogleCSRFToken: (state, action) => {
      state.googleCSRFToken = action.payload;
    },
    resetGoogleCSRFToken: (state) => {
      state.googleCSRFToken = "";
    },
  },
});

export const {
  setUser,
  resetUser,
  setGoogleCSRFToken,
  resetGoogleCSRFToken,
} = userSlice.actions;

export default userSlice.reducer;
