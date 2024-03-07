import { LanguageName } from "@/common/languages";
import { StreamLanguage } from "@codemirror/language";
import { langs } from "@uiw/codemirror-extensions-langs";

type LangaugeMetadata = {
  streamLanguage: StreamLanguage;
  sampleCode: string;
};

export const langaugeMetadata = new Map<LanguageName, LangaugeMetadata>([
  [
    LanguageName.JavaScript,
    {
      streamLanguage: langs.javascript,
      sampleCode: 'console.log("Hello, World!");',
    },
  ],
  [
    LanguageName.Golang,
    {
      streamLanguage: langs.go,
      sampleCode:
        'package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("Hello, World!")\n}',
    },
  ],
  [
    LanguageName.Rust,
    {
      streamLanguage: langs.rust,
      sampleCode: 'fn main() {\n    println!("Hello, World!");\n}',
    },
  ],
  [
    LanguageName.CPlusPlus,
    {
      streamLangauge: langs.cpp,
      sampleCode:
        '#include <iostream>\n\nint main() {\n    std::cout << "Hello, World!" << std::endl;\n    return 0;\n}',
    },
  ],
  [
    LanguageName.Java,
    {
      streamLangauge: langs.java,
      sampleCode:
        'public class HelloWorld {\n    public static void main(String[] args) {\n        System.out.println("Hello, World!");\n    }\n}',
    },
  ],
  [
    LanguageName.C,
    {
      streamLangauge: langs.c,
      sampleCode:
        '#include <stdio.h>\n\nint main() {\n    printf("Hello, World!\\n");\n    return 0;\n}',
    },
  ],
]);
