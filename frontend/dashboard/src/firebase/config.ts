import { initializeApp } from "firebase/app";
import { getFirestore } from "firebase/firestore";

const firebaseConfig = {
  apiKey: "AIzaSyABe7W2A_6gsbz1DONaE0T8wMkH1-NUffc",
  authDomain: "generatesac-testing.firebaseapp.com",
  projectId: "generatesac-testing",
  storageBucket: "generatesac-testing.appspot.com",
  messagingSenderId: "430603036264",
  appId: "1:430603036264:web:80c157c86a6120a1397ddd"
};

const app = initializeApp(firebaseConfig);
const db = getFirestore(app);

export { app, db };