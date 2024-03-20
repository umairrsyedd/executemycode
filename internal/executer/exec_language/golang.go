package execlangauge

import "fmt"

type Golang struct{}

func (g *Golang) GetFileExt() string {
	return ".go"
}

func (g *Golang) IsRunCompiled() bool {
	return false
}

func (g *Golang) GetCompileCmd(fileName string) []string {
	return nil
}

func (g *Golang) GetRunCmd(fileName string) []string {
	return []string{"go", "run", fmt.Sprintf("%s%s", fileName, g.GetFileExt())}
}
