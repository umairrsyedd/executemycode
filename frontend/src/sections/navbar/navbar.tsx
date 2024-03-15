"use client";

import FileName from "./filename";
import InviteButton from "./invite_button";
import LanguageSelect from "./language_select";
import Executer from "./executer/executer";
import Toggler from "./theme_toggle";

import styles from "./navbar.module.css";

export default function Navbar({
  currentLangauge,
  handleLanguageChange,
  setTheme,
  programState,
  setProgramState,
  handleExecute,
  handleStop,
  socketStatus,
}) {
  return (
    <div className={styles.navbar}>
      <div className={styles.navbar_section_left}>
        <FileName />
        <LanguageSelect
          currentLanguage={currentLangauge}
          handleLanguageChange={handleLanguageChange}
        />
      </div>
      <div className={styles.navbar_section_right}>
        <Executer
          programState={programState}
          setProgramState={setProgramState}
          handleExecute={handleExecute}
          handleStop={handleStop}
          socketStatus={socketStatus}
        />
        <div className={styles.navbar_section_right_right}>
          <Toggler onToggle={setTheme} />
          <InviteButton />
        </div>
      </div>
    </div>
  );
}
