package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// JavaExecutor is the struct of java code
type JavaExecutor struct {
	Text     string
	Filename string
}

func (r *JavaExecutor) javaRun() string {
	defer os.Remove("workspace/" + r.Filename + ".java")
	defer os.Remove("workspace/" + r.Filename + ".class")
	dir, _ := os.Getwd()
	cmd := exec.Command("java", r.Filename)
	cmd.Dir = dir + "/workspace"
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
	return string(output)
}

func (r *JavaExecutor) javaFilemaker() {
	ioutil.WriteFile("./workspace/"+r.Filename+".java", []byte(r.Text), 0644)
}

func (r *JavaExecutor) javaBuild() string {
	dir, _ := os.Getwd()
	cmd := exec.Command("javac", r.Filename+".java")
	cmd.Dir = dir + "/workspace"
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// JavaRunner : return string about result of your java program
func (r *JavaExecutor) JavaRunner() (result string) {
	fmt.Println("Receive")
	r.javaFilemaker()
	fmt.Println("File saved by", r.Filename)
	r.javaBuild()
	fmt.Println("Build finished")
	result = r.javaRun()
	fmt.Println("Finished")
	return
}
