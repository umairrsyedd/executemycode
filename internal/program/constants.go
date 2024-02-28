package program

type Language string

const (
	Golang     Language = "golang"
	Python3    Language = "python3"
	JavaScript Language = "javaScript"
	Rust       Language = "rust"
	C          Language = "c"
	CPlusPlus  Language = "c++"
	Java       Language = "java"
)

type ExecutionInfo struct {
	ExecName      string
	ExecArgs      []string
	FileExtension string
}

func GetExecInfo(language Language) ExecutionInfo {
	switch language {
	case Golang:
		return ExecutionInfo{"go", []string{"run"}, "go"}
	case Python3:
		return ExecutionInfo{"python3", nil, "py"}
	default:
		return ExecutionInfo{"unknown", []string{"run"}, "txt"}
	}
}
