package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("编译参数错误")
	}
	file := os.Args[1]
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		log.Fatalf("代码文件(%s)不存在", file)
	} else if err != nil {
		log.Fatal(err)
	}
	if matched, err := regexp.Match(".he$", []byte(file)); !matched || err != nil {
		log.Fatalf("代码文件(%s)格式错误", file)
	}
	absPath, err := filepath.Abs(file)
	if err != nil {
		log.Fatal(err)
	}
	fileName := strings.Replace(file, ".he", "", -1)
	workDir := path.Dir(absPath)

	vmFile, err := ioutil.ReadFile("./vm.go")
	if err != nil {
		log.Fatalf("读取vm代码错误: %s", err)
	}
	codeName := fmt.Sprintf("compile_%s.go", fileName)
	targetCodePath := path.Join(workDir, codeName)
	err = ioutil.WriteFile(targetCodePath, vmFile, 0755)
	if err != nil {
		log.Fatalf("写入目标代码错误: %s", err)
	}
	fmt.Printf("执行编译: %s...\n", file)

	outputName := fileName
	if runtime.GOOS == "windows" {
		outputName = fileName + ".exe"
	}

	timeoutCtx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cmd := exec.CommandContext(timeoutCtx, "go", "build", "-o", outputName, codeName)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workDir
	err = cmd.Start()
	if err != nil { panic(err) }
	err = cmd.Wait()
	if err != nil { panic(err) }

	_ = os.Remove(targetCodePath)

	fmt.Printf("编译结束，可执行程序为 ./%s\n", outputName)
}