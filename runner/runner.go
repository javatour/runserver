package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	NoCompilerError = errors.New("Server does not have the compiler")
	NoVersionError  = errors.New("Server does not support the version")
	NoLanguageError = errors.New("Server does not support the language")
)

type Runner interface {
	Run(path string) (string, error)
}

type CodeFile struct {
	Text     string
	Filename string
	Version  string
	debugger string
	Mode     uint8
	Language string
}

func (target CodeFile) filemaker(path string, lang string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.Command("mkdir", path)
	cmd.Dir = dir + "/workspace"
	cmd.CombinedOutput()
	ioutil.WriteFile("./workspace/"+path+"/"+target.Filename+"."+lang, []byte(target.Text), 0644)
	return nil
}

// Run is the method for running programming source
func (target CodeFile) Run(path string) (string, error) {
	switch target.Language {
	case "java":
		result, err := target.javaExecutor(path)
		if err != nil {
			return result, err
		}
		return result, nil
	case "python":
		result, err := target.pythonExecutor(path)
		if err != nil {
			return result, err
		}
		return result, nil
	case "cplusplus":
		result, err := target.cplusExecutor(path)
		fmt.Println(result, err)
		if err != nil {
			return result, err
		}
		return result, nil
	default:
		return "", NoLanguageError
	}
}
