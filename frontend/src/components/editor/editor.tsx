"use client";

import styles from "./editor.module.css";
import MonacoReact from "@monaco-editor/react";
import * as monaco from "monaco-editor";

import EditorTopBar from "./editor-topbar";

export default function EditorComponent() {
  var editorOptions = {
    fontSize: 16,
  };
  return (
    <div className={styles.editor}>
      <EditorTopBar />
      <MonacoReact
        defaultLanguage="javascript"
        defaultValue=""
        options={editorOptions}
        theme="vs-dark"
      />
    </div>
  );
}
