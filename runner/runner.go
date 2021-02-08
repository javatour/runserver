package runner

import "errors"

var (
	NoCompilerError = errors.New("Server does not have the compiler")
	NoVersionError  = errors.New("Server does not support the version")
	NoLanguageError = errors.New("Server does not support the language")
)

type Runner interface {
	Run() (string, error)
}

type CodeFile struct {
	Text     string
	Filename string
	Version  string
	debugger string
	Mode     uint8
	Language string
}

func (target CodeFile) Run(path string) (string, error) {
	switch target.Language {
	case "java":
		result, err := target.javaExecutor(path)
		if err != nil {
			return result, err
		}
		return result, nil
	case "python":
		return "", NoLanguageError
	case "cplusplus":
		return "", NoLanguageError
	default:
		return "", NoLanguageError
	}
}
