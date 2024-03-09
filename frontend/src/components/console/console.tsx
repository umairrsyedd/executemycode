"use client";

import styles from "./console.module.css";

import Terminal from "react-console-emulator";

const commands = {
  echo: {
    description: "Echo a passed string.",
    usage: "echo <string>",
    fn: (...args) => args.join(" "),
  },
};

export default function Console() {
  return (
    <div className={styles.console}>
      <Terminal commands={commands} disabled={false} />
    </div>
  );
}
