package execlangauge

import "fmt"

type C struct{}

func (c *C) GetName() string {
	return "c"
}

func (c *C) GetFileExt() string {
	return ".c"
}

func (c *C) IsRunCompiled() bool {
	return true
}

func (c *C) GetCompileCmd(fileName string) []string {
	return []string{"gcc", fmt.Sprintf("%s%s", fileName, c.GetFileExt())}
}

func (c *C) GetRunCmd(fileName string) []string {
	return []string{"./a.out"}
}
