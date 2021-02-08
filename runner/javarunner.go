package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func (r CodeFile) javaRun(path string) string {
	//@dev 변경 예정 -> context사용해 무한루프는 강제 종료하도록 변경하고, 해당 부분에서 처리예정
	defer os.Remove("workspace/" + path)
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".java")
	defer os.Remove("workspace/" + path + "/" + r.Filename + ".class")

	dir, _ := os.Getwd()
	cmd := exec.Command("java", "-cp", ".", r.Filename)
	cmd.Dir = dir + "/workspace/" + path
	output, _ := cmd.CombinedOutput()
	return string(output)
}

func (r CodeFile) javaFilemaker(path string) {
	dir, _ := os.Getwd()
	cmd := exec.Command("mkdir", path)
	cmd.Dir = dir + "/workspace"
	cmd.CombinedOutput()
	ioutil.WriteFile("./workspace/"+path+"/"+r.Filename+".java", []byte(r.Text), 0644)
}

func (r CodeFile) javaBuild(path string) string {
	dir, _ := os.Getwd()
	cmd := exec.Command("javac", r.Filename+".java")
	cmd.Dir = dir + "/workspace/" + path
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// Run : return string about result of your java program
// 에러 추가해야함
func (r CodeFile) javaExecutor(path string) (string, error) {
	fmt.Println("Receive")
	r.javaFilemaker(path)
	fmt.Println("File saved by", r.Filename)
	r.javaBuild(path)
	fmt.Println("Build finished")
	result := r.javaRun(path)
	fmt.Println("Finished")
	return result, nil
}
