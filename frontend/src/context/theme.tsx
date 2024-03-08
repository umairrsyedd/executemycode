import { createContext } from "react";

export const enum Themes {
  Dark = "dark",
  Light = "light",
}

export const ThemeContext = createContext(Themes.Dark);
