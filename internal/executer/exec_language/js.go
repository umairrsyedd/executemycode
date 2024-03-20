package execlangauge

import "fmt"

type JavaScript struct{}

func (js *JavaScript) GetFileExt() string {
	return ".js"
}

func (js *JavaScript) IsRunCompiled() bool {
	return false
}

func (js *JavaScript) GetCompileCmd(fileName string) []string {
	return nil
}

func (js *JavaScript) GetRunCmd(fileName string) []string {
	return []string{"node", fmt.Sprintf("%s%s", fileName, js.GetFileExt())}
}
