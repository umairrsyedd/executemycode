"use client";

import styles from "./console.module.css";

import Terminal from "react-console-emulator";

export default function Console() {
  return (
    <div className={styles.console}>
      <Terminal disabled={false} />
    </div>
  );
}
