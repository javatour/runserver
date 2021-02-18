package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func (r CodeFile) javaRun(path string) (string, error) {
	//@dev 변경 예정 -> context사용해 무한루프는 강제 종료하도록 변경하고, 해당 부분에서 처리예정
	defer os.Remove("workspace/" + path)
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".java")
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".class")

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("java", "-cp", ".", r.Filename)
	cmd.Dir = dir + "/workspace/" + path
	output, _ := cmd.CombinedOutput()
	return string(output), nil
}

func (r CodeFile) javaFilemaker(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.Command("mkdir", path)
	cmd.Dir = dir + "/workspace"
	cmd.CombinedOutput()
	ioutil.WriteFile("./workspace/"+path+"/"+r.Filename+".java", []byte(r.Text), 0644)
	return nil
}

func (r CodeFile) javaBuild(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.Command("javac", r.Filename+".java")
	cmd.Dir = dir + "/workspace/" + path
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

// @description Make result by executing a java program
// @author suh wang gyu
// @return value : result and error(if exist)
func (r CodeFile) javaExecutor(path string) (string, error) {
	fmt.Println("Receive")
	err := r.javaFilemaker(path)
	if err != nil {
		return "", err
	}
	fmt.Println("File saved by", r.Filename)
	err = r.javaBuild(path)
	if err != nil {
		return "", err
	}
	fmt.Println("Build finished")
	result, err := r.javaRun(path)
	if err != nil {
		return "", err
	}
	fmt.Println("Finished")
	return result, nil
}
