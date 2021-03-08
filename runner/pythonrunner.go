package runner

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func (r CodeFile) pythonExecutor(path string) (string, error) {
	err := r.filemaker(path, "py")
	if err != nil {
		return "", err
	}

	result, err := r.pythonRun(path)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r CodeFile) pythonRun(path string) (string, error) {
	defer os.Remove("workspace/" + path)
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".py")
	defer func() {
		fmt.Println("garbage collector running")
	}()

	dir, err := os.Getwd()
	if err != nil {
		return "", nil
	}
	cmd := exec.Command("python3.8", r.Filename+".py")
	cmd.Dir = dir + "/workspace/" + path
	resultChan := make(chan string)
	go func() {
		output, _ := cmd.CombinedOutput()
		resultChan <- string(output)
	}()
	select {
	case sucess := <-resultChan:
		return string(sucess), nil
	case <-time.After(5 * time.Second):
		cmd.Process.Kill()
		return "", errors.New("Infinite")
	}
}
