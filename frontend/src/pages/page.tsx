"use client";

import Navbar from "@/components/navbar/navbar";
import Editor from "@/components/editor/editor";
import Console from "@/components/console/console";
import Notepad from "@/components/notepad/notepad";

import styles from "./page.module.css";
import { LanguageName } from "@/common/languages";
import { useState } from "react";

export default function Page() {
  let [currentLanguage, setCurrentLanguage] = useState(LanguageName.JavaScript);
  return (
    <div className={styles.page}>
      <div className={styles.nav_container}>
        <Navbar
          currentLangauge={currentLanguage}
          setCurrentLanguage={setCurrentLanguage}
        />
      </div>
      <div className={styles.main_area}>
        <Editor currentLanguage={currentLanguage} />
        <div className={styles.main_area_right}>
          <Console />
          <Notepad />
        </div>
      </div>
    </div>
  );
}
