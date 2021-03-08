package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func (r CodeFile) javaRun(path string) (string, error) {
	defer os.Remove("workspace/" + path)
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".java")
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".class")

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("java", "-cp", ".", r.Filename)
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
