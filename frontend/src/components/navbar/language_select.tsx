import { LanguageName } from "@/common/languages";
import styles from "./navbar.module.css";

export default function LanguageSelect({
  currentLanguage,
  setCurrentLanguage,
}) {
  const languages = Object.keys(LanguageName);

  const handleLanguageChange = (event) => {
    let selectedLanguage = event.target.value;
    setCurrentLanguage(selectedLanguage);
  };

  return (
    <select
      className={styles.language_select}
      value={currentLanguage}
      onChange={handleLanguageChange}
    >
      {languages.map((key, index) => (
        <>
          <option key={index} value={key}>
            {key}
          </option>
        </>
      ))}
    </select>
  );
}
