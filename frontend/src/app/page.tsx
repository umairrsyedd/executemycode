import Image from "next/image";
import { Editor } from "../components/editor/editor";
import styles from "./style.module.css";

export default function Home() {
  return (
    <div className={styles.page}>
      <h2 className={styles.nav}>Nav Bar</h2>
      <div className={styles.main_area}>
        <Editor />
        <h2 className={styles.console}>Console</h2>
      </div>
    </div>
  );
}
