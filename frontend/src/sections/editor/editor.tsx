"use client";

import styles from "./editor.module.css";
import _ from "lodash";
import CodeMirror, { EditorView } from "@uiw/react-codemirror";
import { andromeda } from "@uiw/codemirror-theme-andromeda";
import { quietlight } from "@uiw/codemirror-theme-quietlight";
import { LanguageName, sampleCodeMap } from "@/types/languages";
import { useContext, useEffect } from "react";
import { ThemeContext, Themes } from "@/context/theme";
import { BiReset } from "react-icons/bi";

import { extensionMap } from "./langauge_metadata";

export default function EditorComponent({
  currentLanguage,
  code,
  setCode,
  handleEditorCodeReset,
}) {
  const theme = useContext(ThemeContext);
  const editorStyle = {
    fontSize: "16px",
  };

  const debouncedCodeSetter = _.debounce((event) => {
    setCode(event);
  }, 500);

  return (
    <div className={styles.editor}>
      <div className={styles.editor__topbar}>
        <BiReset
          color={theme === Themes.Dark ? "white" : "black"}
          size={20}
          className={styles.editor__topbar_reset}
          onClick={handleEditorCodeReset}
        />
      </div>
      <CodeMirror
        onChange={(event) => {
          debouncedCodeSetter(event);
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
