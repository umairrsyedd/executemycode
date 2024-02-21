package executer

type ExecutionInfo struct {
	ExecName      string
	ExecArgs      []string
	FileExtension string
}

func getExecInfo(language Language) ExecutionInfo {
	switch language {
	case Golang:
		return ExecutionInfo{"go", []string{"run"}, "go"}
	case Python:
		return ExecutionInfo{"python", []string{"run"}, "py"}
	default:
		return ExecutionInfo{"unknown", []string{"run"}, "txt"}
	}
}
