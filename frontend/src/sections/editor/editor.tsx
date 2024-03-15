"use client";

import styles from "./editor.module.css";
import _ from "lodash";
import CodeMirror, { EditorView } from "@uiw/react-codemirror";
import { andromeda } from "@uiw/codemirror-theme-andromeda";
import { quietlight } from "@uiw/codemirror-theme-quietlight";
import { LanguageName, sampleCodeMap } from "@/types/languages";
import { useContext, useEffect } from "react";
import { ThemeContext, Themes } from "@/context/theme";

import { extensionMap } from "./langauge_metadata";

export default function EditorComponent({ currentLanguage, code, setCode }) {
  const theme = useContext(ThemeContext);
  const editorStyle = {
    fontSize: "16px",
  };

  const throttledCodeSetter = _.debounce((event) => {
    setCode(event);
  }, 500);

  return (
    <div className={styles.editor}>
      <CodeMirror
        onChange={(event) => {
          throttledCodeSetter(event);
        }}
        value={code}
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
