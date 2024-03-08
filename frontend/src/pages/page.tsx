"use client";

import _ from "lodash";

import Navbar from "@/components/navbar/navbar";
import Editor from "@/components/editor/editor";
import Console from "@/components/console/console";
import Notepad from "@/components/notepad/notepad";

import styles from "./page.module.css";
import { LanguageName, DefaultLanguage } from "@/constants/languages";
import { useRef, useState } from "react";
import ResizableContainer, {
  Orientation,
} from "@/common/resizable/resizable_container";

export default function Page() {
  const [currentLanguage, setCurrentLanguage] = useState(DefaultLanguage);

  return (
    <div className={styles.page}>
      <div className={styles.nav_container}>
        <Navbar
          currentLangauge={currentLanguage}
          setCurrentLanguage={setCurrentLanguage}
        />
      </div>
      <div className={styles.main_area}>
        <ResizableContainer
          orientation={Orientation.Horizontal}
          initialPercent={75}
          minSizePercent={30}
          maxSizePercent={75}
          throttleResize={16}
        >
          <Editor currentLanguage={currentLanguage} />
        </ResizableContainer>
        <div className={styles.main_area_right}>
          <ResizableContainer
            orientation={Orientation.Vertical}
            initialPercent={75}
            minSizePercent={20}
            maxSizePercent={75}
            throttleResize={16}
          >
            <Console />
          </ResizableContainer>
          <Notepad />
        </div>
      </div>
    </div>
  );
}
