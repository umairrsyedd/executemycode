package executer

type ProgramLanguage string

const (
	Golang     ProgramLanguage = "golang"
	Python3    ProgramLanguage = "python3"
	JavaScript ProgramLanguage = "javaScript"
	Rust       ProgramLanguage = "rust"
	C          ProgramLanguage = "c"
	CPlusPlus  ProgramLanguage = "c++"
	Java       ProgramLanguage = "java"
)

func GetFileExtension(language ProgramLanguage) string {
	switch language {
	case Golang:
		return ".go"
	case Python3:
		return ".py"
	case JavaScript:
		return ".js"
	case Rust:
		return ".rs"
	case C:
		return ".c"
	case CPlusPlus:
		return ".cpp"
	case Java:
		return ".java"
	default:
		return ""
	}
}

func GetCmd(language ProgramLanguage) []string {
	switch language {
	case Golang:
		return []string{"go", "run"}
	default:
		return []string{}
	}
}
