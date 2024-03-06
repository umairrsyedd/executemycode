"use client";

import styles from "./editor.module.css";
import CodeMirror from "@uiw/react-codemirror";

import { andromeda } from "@uiw/codemirror-theme-andromeda";

export default function EditorComponent() {
  return (
    <div className={styles.editor}>
      <CodeMirror value="console.log('hello world!');" theme={andromeda} />
    </div>
  );
}
