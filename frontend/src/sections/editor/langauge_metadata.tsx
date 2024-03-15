import { LanguageName } from "@/types/languages";
import { Language } from "@codemirror/language";
import { langs } from "@uiw/codemirror-extensions-langs";

export let extensionMap = new Map<Language, any>([
  [LanguageName.JavaScript, langs.javascript],
  [LanguageName.Golang, langs.go],
  [LanguageName.Rust, langs.rust],
  [LanguageName.CPlusPlus, langs.cpp],
  [LanguageName.Java, langs.java],
  [LanguageName.C, langs.c],
]);
