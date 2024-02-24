package executer

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

type ProgramState string

const (
	Idle       ProgramState = "idle"
	Running    ProgramState = "running"
	Terminated ProgramState = "terminated"
)
