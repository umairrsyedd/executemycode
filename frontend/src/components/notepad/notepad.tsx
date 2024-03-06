"use client";

import styles from "./notepad.module.css";

export default function Notepad() {
  return <div contentEditable={true} className={styles.notepad}></div>;
}
