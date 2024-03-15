package executer

type ProgramLanguage string

const (
	Golang     ProgramLanguage = "Golang"
	JavaScript ProgramLanguage = "JavaScript"
	Rust       ProgramLanguage = "Rust"
	C          ProgramLanguage = "C"
	CPlusPlus  ProgramLanguage = "C++"
	Java       ProgramLanguage = "Java"
)

func getFileExtension(language ProgramLanguage) string {
	switch language {
	case Golang:
		return ".go"
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

func getCmd(language ProgramLanguage) []string {
	switch language {
	case Golang:
		return []string{"go", "run"}
	default:
		return []string{}
	}
}
