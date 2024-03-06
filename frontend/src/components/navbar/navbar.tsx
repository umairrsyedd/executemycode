"use client";

import InviteButton from "./invite_button";
import LanguageSelect from "./language_select";
import styles from "./navbar.module.css";
import RunButton from "./run_button";

export default function Navbar() {
  return (
    <div className={styles.navbar}>
      <div className={styles.navbar_section_left}>
        <div>NastyProgrammer </div>
        <LanguageSelect />
      </div>
      <div className={styles.navbar_section_right}>
        <RunButton />
        <InviteButton />
      </div>
    </div>
  );
}
