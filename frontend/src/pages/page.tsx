import Navbar from "@/components/navbar/navbar";
import Editor from "@/components/editor/editor";
import Console from "@/components/console/console";

import styles from "./page.module.css";

export default function Page() {
  return (
    <div className={styles.page}>
      <div className={styles.nav_container}>
        <Navbar />
      </div>
      <div className={styles.main_area}>
        <Editor />
        <Console />
      </div>
    </div>
  );
}
