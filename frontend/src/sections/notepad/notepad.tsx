"use client";

import { useEffect, useState, useMemo } from "react";
import dynamic from "next/dynamic";
import styles from "./notepad.module.css";

import "react-quill/dist/quill.snow.css";

export default function Notepad() {
  const [value, setValue] = useState("");
  const ReactQuill = useMemo(
    () => dynamic(() => import("react-quill"), { ssr: false }),
    []
  );
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
    <div className={styles.notepad}>
      <ReactQuill
        theme="snow"
        value={value}
        onChange={setValue}
        modules={{ toolbar: toolbarOptions }}
        placeholder="Jot down some notes..."
      />
    </div>
  );
}
