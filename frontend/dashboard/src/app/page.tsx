"use client";
import Editor from "../components/editor/advanced-editor";
import { JSONContent } from "novel";

import { useState, useEffect } from "react";
import { defaultValue } from "./default-value";
import { Button } from "@/components/ui/button";

export default function Home() {
  const [json, setJSON] = useState<JSONContent>(defaultValue);
  const [isEditing, setIsEditing] = useState(true);

  useEffect(() => {
    const savedContent = localStorage.getItem("editorContent");
    if (savedContent) {
      setJSON(JSON.parse(savedContent));
    }
  }, []);

  const handleSave = () => {
    setIsEditing(false);  // Switch to view mode
  };

  const handleEdit = () => {
    setIsEditing(true);  // Switch back to edit mode
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="flex flex-col p-6 border w-full gap-6 rounded-md bg-card">
        <div className="flex justify-between items-center">
          <h1 className="text-4xl font-semibold">
            {isEditing ? "Edit your text" : "View your JSON"}
          </h1>
          {isEditing ? (
            <Button onClick={handleSave}>Save</Button>
          ) : (
            <Button onClick={handleEdit}>Edit</Button>
          )}
        </div>
        {isEditing ? (
          <Editor initialValue={json} onChange={setJSON} />
        ) : (
          <pre className="whitespace-pre-wrap">{JSON.stringify(json, null, 2)}</pre>
        )}
      </div>
    </main>
  );
}
