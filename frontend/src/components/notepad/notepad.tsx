"use client";

import { useEffect, useState } from "react";
import styles from "./notepad.module.css";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";

export default function Notepad() {
  const [value, setValue] = useState("Jot down some notes...");
  const toolbarOptions = [
    { size: ["small", "normal", "large"] },
    "bold",
    "italic",
    "underline",
    { align: [] },
    { list: "ordered" },
    { list: "bullet" },
    { list: "check" },
  ];
  const toolBarConfig = {
    modules: {
      toolbar: toolbarOptions,
    },
  };
  return (
    <>
      <div className={styles.notepad}>
        <ReactQuill
          theme="snow"
          value={value}
          onChange={setValue}
          modules={{ toolbar: toolbarOptions }}
        />
      </div>
    </>
  );
}
