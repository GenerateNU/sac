import Image from "next/image";
import { FirebaseController } from "../firebase/controller";
import Tiptap from "../components/RichTextEditor/Tiptap";

export default function Home() {
  return (
    <main>
      <Tiptap />
    </main>
  );
}
