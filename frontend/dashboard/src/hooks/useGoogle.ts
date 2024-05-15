import { oauth2 } from "googleapis/build/src/apis/oauth2";
import { generateSHA256Hash } from "../utils/token";
import { useAppDispatch, useAppSelector } from "../store/store";
import { setGoogleCSRFToken } from "../store/userSlice";

const BASE_GOOGLE_OAUTH_URL = 'https://accounts.google.com/o/oauth2/v2/auth/oauthchooseaccount'

const googleConfig = {
  client_id: process.env.NEXT_PUBLIC_GOOGLE_OAUTH_CLIENT_ID || "",
  /** TODO: redirect_uri needs to be updated to pull from current environment config */
  redirect_uri: "http://localhost:3000",
  response_type: "code",
  scope: "https://www.googleapis.com/auth/calendar.events https://www.googleapis.com/auth/calendar.readonly",
  include_granted_scopes: 'true',
  access_type: "offline",
  prompt: "consent",
};

export default function useGoogle() {
  /** Redirect the user to Google OAuth */
  function googleRedirect() {
    const token = generateSHA256Hash(); 

    const config = {
      ...googleConfig,
      state: token,
    }

    const ouathUrl = BASE_GOOGLE_OAUTH_URL + "?" + new URLSearchParams(config).toString();
    window.location.href = ouathUrl;
  }

  return { 
    googleRedirect,
  };
}
