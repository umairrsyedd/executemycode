"use client";

import styles from "./editor.module.css";
import CodeMirror from "@uiw/react-codemirror";
import { andromeda } from "@uiw/codemirror-theme-andromeda";
import { quietlight } from "@uiw/codemirror-theme-quietlight";
import { langaugeMetadata } from "./langauge_metadata";
import { LanguageName } from "@/constants/languages";
import { useContext } from "react";
import { ThemeContext, Themes } from "@/context/theme";

export default function EditorComponent({ currentLanguage }) {
  const theme = useContext(ThemeContext);
  var languageMetadata = langaugeMetadata.get(currentLanguage);
  const editorStyle = {
    fontSize: "16px",
  };

  return (
    <div className={styles.editor}>
      <CodeMirror
        value={languageMetadata?.sampleCode}
        extensions={[languageMetadata?.streamLanguage()]}
        theme={theme === Themes.Dark ? andromeda : quietlight}
        style={editorStyle}
        height="100vh"
      />
    </div>
  );
}
