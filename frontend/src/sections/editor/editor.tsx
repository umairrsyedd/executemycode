"use client";

import styles from "./editor.module.css";
import CodeMirror, { EditorView } from "@uiw/react-codemirror";
import { andromeda } from "@uiw/codemirror-theme-andromeda";
import { quietlight } from "@uiw/codemirror-theme-quietlight";
import { LanguageName } from "@/constants/languages";
import { useContext } from "react";
import { ThemeContext, Themes } from "@/context/theme";

import { extensionMap, sampleCodeMap } from "./langauge_metadata";

export default function EditorComponent({ currentLanguage }) {
  const theme = useContext(ThemeContext);
  const editorStyle = {
    fontSize: "16px",
  };

  return (
    <div className={styles.editor}>
      <CodeMirror
        value={sampleCodeMap.get(currentLanguage)}
        extensions={[
          extensionMap.get(currentLanguage)(),
          EditorView.lineWrapping,
        ]}
        theme={theme === Themes.Dark ? andromeda : quietlight}
        style={editorStyle}
        height="90vh"
      />
    </div>
  );
}
