'use client'

import { useEffect, useState } from "react";
import Editor from "../components/RichTextEditor/BlockNote";
import { FirebaseController } from "../firebase/controller";
import { PartialBlock } from "@blocknote/core";
import { EditorFonts, editorFonts } from "../components/RichTextEditor/config/config";
 
export default function Home() {
  const [initialContent, setInitialContent] = useState<PartialBlock[]>([]);
  const [initialContentLoading, setInitialContentLoading] = useState<boolean>(true);
  const [description, setDescription] = useState<Record<string, string>>({});
  const [fontFamily, setFontFamily] = useState<EditorFonts>('monospace');
  const mockFonts = ['monospace', 'timesNewRoman', 'comicSans'] as EditorFonts[];

  useEffect(() => {
    setInitialContentLoading(true);
    FirebaseController.getDescription('1')
      .then((data) => {
        setInitialContent(Object.values(data) as PartialBlock[]);
      })
      .catch((e) => {
        console.log(e);
      })
      .finally(() => {
        setInitialContentLoading(false);
      })
  }, []);

  function handleSave() {
    FirebaseController.addDescription(description, '1')
      .then(() => {
        console.log("Document successfully written!");
      }).catch((e) => {
        console.log(e);
      });
  };

  return (
    <main>
      {!initialContentLoading &&
        <>
          <button 
            style={{
              backgroundColor: 'black',
              borderRadius: '10px',
              color: 'white',
              padding: '10px 20px 10px 20px',
            }}
            onClick={() => setFontFamily(mockFonts[Math.round(Math.random() * 2)])}
          >
            Change Font
          </button>
          <div style={{ padding: '20px', display: 'flex', flexDirection: 'column', gap: '10px' }}>
            <Editor 
              editable={true}
              description={initialContent}
              setDescription={setDescription}
              fontFamily={fontFamily}
            />
            <div>
              <button 
                style={{
                  backgroundColor: 'black',
                  borderRadius: '10px',
                  color: 'white',
                  padding: '10px 20px 10px 20px',
                }}
                onClick={handleSave}
              >
                Save
              </button>
            </div>
          </div>
        </>
      }
    </main>
  );
}
