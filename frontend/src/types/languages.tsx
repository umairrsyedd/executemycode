export enum LanguageName {
  JavaScript = "JavaScript",
  Golang = "Golang",
  Rust = "Rust",
  CPlusPlus = "C++",
  Java = "Java",
  C = "C",
}

export const DefaultLanguage = LanguageName.Golang;

export const LocalStoragePrefix = "code_";

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
