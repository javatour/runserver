package runner

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// cplusrunner.go 는 c++ 파일을 생성하고, 빌드하고, 실행하는 함수들의 집합입니다.
// 패키지 내부의 함수로 External interface는 runner.go 에만 있습니다.


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
	if err != nil {
		return "", err
	}
	err = r.cplusBuild(path)
	if err != nil {
		return "", err
	}
	result, err := r.cplusRun(path)
	if err != nil {
		return "", err
	}
	return result, nil
}
