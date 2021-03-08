package runner

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func (r CodeFile) cplusRun(path string) (string, error) {
	defer os.Remove("workspace/" + path)
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".cpp")
	defer os.Remove("workspace/" + path + "/a.out")
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("./a.out")
	cmd.Dir = dir + "/workspace/" + path
	resultChan := make(chan string)
	go func() {
		output, _ := cmd.CombinedOutput()
		resultChan <- string(output)
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case <-time.After(5 * time.Second):
		cmd.Process.Kill()
		return "", errors.New("Infinite")
	}
}

func (r CodeFile) cplusBuild(path string) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.Command("c++", r.Filename+".cpp")
	cmd.Dir = dir + "/workspace/" + path
	_, err = cmd.CombinedOutput()
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (r CodeFile) cplusExecutor(path string) (string, error) {
	err := r.filemaker(path, "cpp")
	fmt.Println("here")
	if err != nil {
		return "", err
	}
	fmt.Println("here2")
	err = r.cplusBuild(path)
	if err != nil {
		return "", err
	}
	fmt.Println("here3")
	result, err := r.cplusRun(path)
	if err != nil {
		return "", err
	}
	fmt.Println("here4")
	return result, nil
}
