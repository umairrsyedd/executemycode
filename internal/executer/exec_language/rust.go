package execlangauge

import "fmt"

type Rust struct{}

func (rs *Rust) GetName() string {
	return "rust"
}

func (rs *Rust) GetFileExt() string {
	return ".rc"
}

func (rs *Rust) IsRunCompiled() bool {
	return true
}

func (r *Rust) GetCompileCmd(fileName string) []string {
	return []string{"rustc", fmt.Sprintf("%s%s", fileName, r.GetFileExt())}
}

func (r *Rust) GetRunCmd(fileName string) []string {
	return []string{fmt.Sprintf("./%s", fileName)}
}
