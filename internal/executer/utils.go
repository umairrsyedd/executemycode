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
	case Python3:
		return ExecutionInfo{"python3", nil, "py"}
	default:
		return ExecutionInfo{"unknown", []string{"run"}, "txt"}
	}
}
