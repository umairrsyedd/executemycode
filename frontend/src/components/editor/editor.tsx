"use client";

import styles from "./editor.module.css";
import CodeMirror from "@uiw/react-codemirror";
import { andromeda } from "@uiw/codemirror-theme-andromeda";
import { langaugeMetadata } from "./langauge_metadata";
import { LanguageName } from "@/common/languages";

export default function EditorComponent({ currentLanguage }) {
  var languageMetadata = langaugeMetadata.get(currentLanguage);
  const editorStyle = {
    fontSize: "16px",
  };

  return (
    <div className={styles.editor}>
      <CodeMirror
        value={languageMetadata?.sampleCode}
        extensions={[languageMetadata?.streamLanguage()]}
        theme={andromeda}
        style={editorStyle}
      />
    </div>
  );
}
