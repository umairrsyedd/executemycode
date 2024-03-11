import { useContext } from "react";
import styles from "./navbar.module.css";
import { ThemeContext, Themes } from "@/context/theme";
import { MdDarkMode, MdLightMode } from "react-icons/md";
import { useLocalStorage } from "@uidotdev/usehooks";

export default function Toggler({ onToggle }) {
  const [theme] = useLocalStorage("theme");

  const handleToggle = () => {
    onToggle();
  };

  return (
    <div className={styles.theme_toggler} onClick={handleToggle}>
      {theme === Themes.Dark ? (
        <MdDarkMode size={20} color="white" />
      ) : (
        <MdLightMode size={20} color="black" />
      )}
    </div>
  );
}
