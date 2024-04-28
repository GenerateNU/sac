import { Block } from "@blocknote/core";
import { FirebaseApp } from "firebase/app";
import { DocumentData, Firestore, doc, getDoc, setDoc } from "firebase/firestore";
import { db } from "./config";

const FirebaseController = {
    // Add a new document to firebase:
    async addDescription(description: Record<string, string>, id: string): Promise<void> {
        await setDoc(doc(db, "descriptions", id), { ...description })
            .then(() => {
                console.log("Document successfully written!");
            })
            .catch(() => {
                console.error("Error writing document");
            });
    },

    // Reads a document from firebase given an id:
    async getDescription(id: string): Promise<DocumentData> {
        const docRef = doc(db, "descriptions", id);
        const docSnap = await getDoc(docRef);

        if (docSnap.exists()) {
            console.log("Document data:", docSnap.data());
            return docSnap.data();
        } else {
            console.error("No such document!");
            return {};
        }
    },
};

export { FirebaseController };
