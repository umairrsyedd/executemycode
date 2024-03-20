package execlangauge

import "fmt"

type Java struct{}

func (java *Java) GetName() string {
	return "java"
}

func (java *Java) GetFileExt() string {
	return ".java"
}

func (java *Java) IsRunCompiled() bool {
	return true
}

func (java *Java) GetCompileCmd(fileName string) []string {
	return []string{"javac", fmt.Sprintf("%s%s", fileName, java.GetFileExt())}
}

func (java *Java) GetRunCmd(fileName string) []string {
	return []string{fmt.Sprintf("java %s", fileName)}
}
