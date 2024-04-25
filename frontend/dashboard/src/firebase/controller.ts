import { FirebaseApp } from "firebase/app";
import { Firestore, doc, getDoc, setDoc } from "firebase/firestore";

const FirebaseController = {
    db: null as Firestore | null,
    app: null as FirebaseApp | null,

    initialize(app: FirebaseApp, db: Firestore) {
        this.app = app;
        this.db = db;
    },

    // Add a new document to firebase:
    async addDescription(description: Record<string, string>, id: string): Promise<void> {
        if (!this.db) {
            throw new Error("Firebase Firestore not initialized");
        }

        await setDoc(doc(this.db, "descriptions", id), { ...description })
            .then(() => {
                console.log("Document successfully written!");
            })
            .catch(() => {
                console.error("Error writing document");
            });
    },

    // Reads a document from firebase given an id:
    async getDescription(id: string): Promise<Record<string, string>> {
        if (!this.db) {
            throw new Error("Firebase Firestore not initialized");
        }

        const docRef = doc(this.db, "descriptions", id);
        const docSnap = await getDoc(docRef);

        if (docSnap.exists()) {
            console.log("Document data:", docSnap.data());
            return docSnap.data() as Record<string, string>;
        } else {
            console.error("No such document!");
            return {};
        }
    },
};

export { FirebaseController };
