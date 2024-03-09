import { LanguageName } from "@/constants/languages";
import { Language, StreamLanguage } from "@codemirror/language";
import { langs } from "@uiw/codemirror-extensions-langs";

type LangaugeMetadata = {
  streamLanguage: any;
  sampleCode: string;
};

export let extensionMap = new Map<Language, any>([
  [LanguageName.JavaScript, langs.javascript],
  [LanguageName.Golang, langs.go],
  [LanguageName.Rust, langs.rust],
  [LanguageName.CPlusPlus, langs.cpp],
  [LanguageName.Java, langs.java],
  [LanguageName.C, langs.c],
]);

export let sampleCodeMap = new Map<Language, string>([
  [LanguageName.JavaScript, `console.log("Hello, World!")`],
  [
    LanguageName.Golang,
    `package main\n\nimport "fmt"\n\nfunc main() {\n  fmt.Println("Hello, World!")\n}`,
  ],
  [LanguageName.Rust, `fn main() {\n  println!("Hello, World!");\n}`],
  [
    LanguageName.CPlusPlus,
    `#include <iostream>\n\nint main() {\n  std::cout << "Hello, World!" << std::endl;\n  return 0;\n}`,
  ],
  [
    LanguageName.Java,
    `public class HelloWorld {\n  public static void main(String[] args) {\n    System.out.println("Hello, World!");\n    }\n}`,
  ],
  [
    LanguageName.C,
    `#include <stdio.h>\n\nint main() {\n  printf("Hello, World!\\n");\n  return 0;\n}`,
  ],
]);
