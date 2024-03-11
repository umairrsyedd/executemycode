"use client";

import _ from "lodash";

import Navbar from "@/sections/navbar/navbar";
import Editor from "@/sections/editor/editor";
import Console from "@/sections/console/console";
import Notepad from "@/sections/notepad/notepad";

import styles from "./page.module.css";
import { LanguageName, DefaultLanguage } from "@/constants/languages";
import { useRef, useState } from "react";
import ResizableContainer, {
  Orientation,
} from "@/components/resizable/resizable_container";
import { ThemeContext, Themes } from "@/context/theme";
import { useLocalStorage } from "@uidotdev/usehooks";

export default function Page() {
  const [currentLanguage, setCurrentLanguage] = useState(DefaultLanguage);
  const [currentTheme, setTheme] = useLocalStorage("theme", Themes.Dark);

  const handleThemeToggle = () => {
    setTheme((prevTheme) =>
      prevTheme === Themes.Dark ? Themes.Light : Themes.Dark
    );
  };

  return (
    <ThemeContext.Provider value={currentTheme}>
      <div className={styles.page} data-theme={currentTheme}>
        <div className={styles.nav_container}>
          <Navbar
            currentLangauge={currentLanguage}
            setCurrentLanguage={setCurrentLanguage}
            setTheme={handleThemeToggle}
          />
        </div>
        <div className={styles.main_area}>
          <ResizableContainer
            orientation={Orientation.Horizontal}
            initialPercent={75}
            minSizePercent={30}
            maxSizePercent={75}
          >
            <Editor currentLanguage={currentLanguage} />
          </ResizableContainer>
          <div className={styles.main_area_right}>
            <ResizableContainer
              orientation={Orientation.Vertical}
              initialPercent={50}
              minSizePercent={20}
              maxSizePercent={90}
            >
              <Console />
            </ResizableContainer>
            <Notepad />
          </div>
        </div>
      </div>
    </ThemeContext.Provider>
  );
}
