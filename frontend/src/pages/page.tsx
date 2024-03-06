import Navbar from "@/components/navbar/navbar";
import Editor from "@/components/editor/editor";
import Console from "@/components/console/console";
import Notepad from "@/components/notepad/notepad";

import styles from "./page.module.css";

export default function Page() {
  return (
    <div className={styles.page}>
      <div className={styles.nav_container}>
        <Navbar />
      </div>
      <div className={styles.main_area}>
        <Editor />
        <div className={styles.main_area_right}>
          <Console />
          <Notepad />
        </div>
      </div>
    </div>
  );
}
