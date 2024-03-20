package execlangauge

import "executemycode/pkg/language"

type LanguageExecuter interface {
	GetFileExt() string
	IsRunCompiled() bool
	GetCompileCmd(fileName string) []string
	GetRunCmd(fileName string) []string
}

func New(lang language.ProgramLanguage) LanguageExecuter {
	switch lang {
	case language.Golang:
		return &Golang{}
	case language.JavaScript:
		return &JavaScript{}
	case language.Rust:
		return &Rust{}
	case language.Java:
		return &Java{}
	case language.CPlusPlus:
		return &CPlusPlus{}
	case language.C:
		return &C{}
	}

	return nil
}
