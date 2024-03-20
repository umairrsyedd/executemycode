package execlangauge

import "fmt"

type CPlusPlus struct{}

func (cpp *CPlusPlus) GetName() string {
	return "c++"
}

func (cpp *CPlusPlus) GetFileExt() string {
	return ".cpp"
}

func (cpp *CPlusPlus) IsRunCompiled() bool {
	return true
}

func (cpp *CPlusPlus) GetCompileCmd(fileName string) []string {
	return []string{"g++", fmt.Sprintf("%s%s", fileName, cpp.GetFileExt())}
}

func (cpp *CPlusPlus) GetRunCmd(fileName string) []string {
	return []string{"./a.out"}
}
