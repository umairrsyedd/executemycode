import { LanguageName } from "@/types/languages";
import styles from "./navbar.module.css";
import { useEffect } from "react";

export default function LanguageSelect({
  currentLanguage,
  handleLanguageChange,
}) {
  const languages = Object.keys(LanguageName).map(
    (key) => LanguageName[key as keyof typeof LanguageName]
  );

  const handleChange = (event) => {
    let selectedLanguage = event.target.value;
    handleLanguageChange(selectedLanguage);
  };

  return (
    <select
      className={styles.language_select}
      value={currentLanguage}
      onChange={handleChange}
    >
      {languages.map((key, index) => (
        <option key={index} value={key}>
          {key}
        </option>
      ))}
    </select>
  );
}
